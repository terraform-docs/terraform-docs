/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package terraform

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/hcl/v2/ext/typeexpr"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/opentofu/opentofu-schema/earlydecoder"
	"github.com/opentofu/opentofu-schema/module"
	tfaddr "github.com/opentofu/registry-address"
	"github.com/zclconf/go-cty/cty"
	ctyjson "github.com/zclconf/go-cty/cty/json"

	"github.com/terraform-docs/terraform-docs/internal/reader"
	"github.com/terraform-docs/terraform-docs/internal/types"
	"github.com/terraform-docs/terraform-docs/print"
)

// LoadWithOptions returns new instance of Module with all the inputs and
// outputs discovered from provided 'path' containing Terraform config
func LoadWithOptions(config *print.Config) (*Module, error) {
	tfmodule, err := loadModule(config.ModuleRoot)
	if err != nil {
		return nil, err
	}

	module, err := loadModuleItems(tfmodule, config)
	if err != nil {
		return nil, err
	}
	sortItems(module, config)
	return module, nil
}

func loadModule(path string) (*module.Meta, error) {
	files, diags := parseModuleFiles(path)
	if diags.HasErrors() {
		return nil, diags
	}
	meta, diags := earlydecoder.LoadModule(path, files)
	if diags.HasErrors() {
		return nil, diags
	}
	return meta, nil
}

func ctyTypetoString(t cty.Type) string {
	if t == cty.NilType {
		return ""
	}
	return typeexpr.TypeString(t)
}

func ctyValueToString(v cty.Value) string {
	if v == cty.NilVal {
		return ""
	}
	b, err := ctyjson.Marshal(v, v.Type())
	if err != nil {
		return ""
	}
	return string(b)
}

func loadModuleItems(tfmodule *tfconfig.Module, config *print.Config) (*Module, error) {
	header, err := loadHeader(config)
	if err != nil {
		return nil, err
	}

	footer, err := loadFooter(config)
	if err != nil {
		return nil, err
	}

	inputs, required, optional := loadInputs(tfmodule, config)
	modulecalls := loadModulecalls(tfmodule, config)
	outputs, err := loadOutputs(tfmodule, config)
	if err != nil {
		return nil, err
	}
	providers := loadProviders(tfmodule, config)
	requirements := loadRequirements(tfmodule)
	resources := loadResources(tfmodule, config)

	return &Module{
		Header:       header,
		Footer:       footer,
		Inputs:       inputs,
		ModuleCalls:  modulecalls,
		Outputs:      outputs,
		Providers:    providers,
		Requirements: requirements,
		Resources:    resources,

		RequiredInputs: required,
		OptionalInputs: optional,
	}, nil
}

func getFileFormat(filename string) string {
	if filename == "" {
		return ""
	}
	last := strings.LastIndex(filename, ".")
	if last == -1 {
		return ""
	}
	return filename[last:]
}

func isFileFormatSupported(filename string, section string) (bool, error) {
	if section == "" {
		return false, errors.New("section is missing")
	}
	if filename == "" {
		return false, fmt.Errorf("--%s-from value is missing", section)
	}
	switch getFileFormat(filename) {
	case ".adoc", ".md", ".tf", ".tofu", ".txt":
		return true, nil
	}
	return false, fmt.Errorf("only .adoc, .md, .tf, .tofu and .txt formats are supported to read %s from", section)
}

func loadHeader(config *print.Config) (string, error) {
	if !config.Sections.Header {
		return "", nil
	}
	return loadSection(config, config.HeaderFrom, "header")
}

func loadFooter(config *print.Config) (string, error) {
	if !config.Sections.Footer {
		return "", nil
	}
	return loadSection(config, config.FooterFrom, "footer")
}

func loadSection(config *print.Config, file string, section string) (string, error) { //nolint:gocyclo
	// NOTE(khos2ow): this function is over our cyclomatic complexity goal.
	// Be wary when adding branches, and look for functionality that could
	// be reasonably moved into an injected dependency.

	if section == "" {
		return "", errors.New("section is missing")
	}
	filename := filepath.Join(config.ModuleRoot, file)
	if ok, err := isFileFormatSupported(file, section); !ok {
		return "", err
	}
	if info, err := os.Stat(filename); os.IsNotExist(err) || info.IsDir() {
		if section == "header" && file == "main.tf" {
			return "", nil // absorb the error to not break workflow for default value of header and missing 'main.tf'
		}
		return "", err // user explicitly asked for a file which doesn't exist
	}
	format := getFileFormat(file)
	if format != ".tf" && format != ".tofu" {
		content, err := os.ReadFile(filepath.Clean(filename))
		if err != nil {
			return "", err
		}
		return string(content), nil
	}
	lines := reader.Lines{
		FileName: filename,
		LineNum:  -1,
		Condition: func(line string) bool {
			line = strings.TrimSpace(line)
			return strings.HasPrefix(line, "/*") || strings.HasPrefix(line, "*") || strings.HasPrefix(line, "*/")
		},
		Parser: func(line string) (string, bool) {
			tmp := strings.TrimSpace(line)
			if strings.HasPrefix(tmp, "/*") || strings.HasPrefix(tmp, "*/") {
				return "", false
			}
			if tmp == "*" {
				return "", true
			}
			line = strings.TrimLeft(line, " ")
			line = strings.TrimRight(line, "\r\n")
			line = strings.TrimPrefix(line, "* ")
			return line, true
		},
	}
	sectionText, err := lines.Extract()
	if err != nil {
		return "", err
	}
	return strings.Join(sectionText, "\n"), nil
}

func loadInputs(meta *module.Meta, positions map[string]Position, config *print.Config) ([]*Input, []*Input, []*Input) {
	inputs := make([]*Input, 0, len(meta.Variables))
	required := make([]*Input, 0)
	optional := make([]*Input, 0)

	for name, input := range meta.Variables {
		position := positions[name]
		comments := loadComments(position.Filename, position.Line)

		// Skip over inputs that are marked as being ignored
		if strings.Contains(comments, "terraform-docs-ignore") {
			continue
		}

		// convert CRLF to LF early on (https://github.com/terraform-docs/terraform-docs/issues/305)
		description := strings.ReplaceAll(input.Description, "\r\n", "\n")
		if description == "" && config.Settings.ReadComments {
			description = comments
		}

		isRequired := input.DefaultValue == cty.NilVal

		in := &Input{
			Name:        name,
			Type:        types.String(ctyTypetoString(input.Type)),
			Description: types.String(description),
			Default:     types.String(ctyValueToString(input.DefaultValue)),
			Required:    isRequired,
			Position:    position,
		}

		inputs = append(inputs, in)

		if in.HasDefault() {
			optional = append(optional, in)
		} else {
			required = append(required, in)
		}
	}

	return inputs, required, optional
}

func formatSource(s, v string) (source, version string) {
	substr := "?ref="

	if v != "" {
		return s, v
	}

	pos := strings.LastIndex(s, substr)
	if pos == -1 {
		return s, version
	}

	adjustedPos := pos + len(substr)
	if adjustedPos >= len(s) {
		return s, version
	}

	source = s[0:pos]
	version = s[adjustedPos:]

	return source, version
}

func loadModuleCalls(meta *module.Meta, config *print.Config) []*ModuleCall {
	modules := make([]*ModuleCall, 0, len(meta.ModuleCalls))

	for _, moduleCall := range meta.ModuleCalls {
		var filename string
		var line int

		if moduleCall.RangePtr != nil {
			filename = moduleCall.RangePtr.Filename
			line = moduleCall.RangePtr.Start.Line
		}

		comments := loadComments(filename, line)
		if strings.Contains(comments, "terraform-docs-ignore") {
			continue
		}

		description := ""
		if config.Settings.ReadComments {
			description = comments
		}

		source, version := moduleSourceAndVersion(moduleCall)

		modules = append(modules, &ModuleCall{
			Name:        moduleCall.LocalName,
			Source:      source,
			Version:     version,
			Description: types.String(description),
			Position: Position{
				Filename: filename,
				Line:     line,
			},
		})
	}
	return modules
}

// moduleSourceAndVersion flattens the typed ModuleSourceAddr back into a (source, version) pair compatible with the
// old tfconfig output.
func moduleSourceAndVersion(moduleCall module.DeclaredModuleCall) (string, string) {
	declaredVersion := ""
	if len(moduleCall.Version) > 0 {
		declaredVersion = moduleCall.Version.String()
	}

	switch source := moduleCall.SourceAddr.(type) {
	case tfaddr.Module:
		// registry address version comes from the `version = "..."` arg.
		return source.ForDisplay(), declaredVersion
	case module.LocalSourceAddr:
		return string(source), declaredVersion
	case module.RemoteSourceAddr:
		// remote sources may carry `?ref=...` which should surface as version.
		return formatSource(string(source), declaredVersion)
	case module.UnknownSourceAddr:
		return formatSource(string(source), declaredVersion)
	default:
		// nil SourceAddr falls back to raw string.
		return formatSource(module.RawSourceAddr, declaredVersion)
	}
}

func loadOutputs(meta *module.Meta, positions map[string]Position, config *print.Config) ([]*Output, error) {
	outputs := make([]*Output, 0, len(meta.Outputs))
	values := make(map[string]*output)

	if config.OutputValues.Enabled {
		var err error
		values, err = loadOutputValues(config)
		if err != nil {
			return nil, err
		}
	}

	for name, output := range meta.Outputs {
		position := positions[name]
		comments := loadComments(position.Filename, position.Line)

		// Skip over outputs that are marked as being ignored
		if strings.Contains(comments, "terraform-docs-ignore") {
			continue
		}

		// convert CRLF to LF early on (https://github.com/terraform-docs/terraform-docs/issues/584)
		description := strings.ReplaceAll(output.Description, "\r\n", "\n")
		if description == "" && config.Settings.ReadComments {
			description = comments
		}

		out := &Output{
			Name:        name,
			Description: types.String(description),
			Position:    position,
			ShowValue:   config.OutputValues.Enabled,
		}

		if config.OutputValues.Enabled {
			if value, ok := values[out.Name]; ok {
				out.Sensitive = value.Sensitive
				out.Value = types.ValueOf(value.Value)
			} else {
				out.Value = types.ValueOf("null")
			}

			if out.Sensitive {
				out.Value = types.ValueOf(`<sensitive>`)
			}
		}
		outputs = append(outputs, out)
	}
	return outputs, nil
}

func loadOutputValues(config *print.Config) (map[string]*output, error) {
	var out []byte
	var err error
	if config.OutputValues.From == "" {
		cmd := exec.CommandContext(context.TODO(), "terraform", "output", "-json")
		cmd.Dir = config.ModuleRoot
		if out, err = cmd.Output(); err != nil {
			return nil, fmt.Errorf("caught error while reading the terraform outputs: %w", err)
		}
	} else if out, err = os.ReadFile(config.OutputValues.From); err != nil {
		return nil, fmt.Errorf("caught error while reading the terraform outputs file at %s: %w", config.OutputValues.From, err)
	}
	var terraformOutputs map[string]*output
	err = json.Unmarshal(out, &terraformOutputs)
	if err != nil {
		return nil, err
	}
	return terraformOutputs, err
}

func loadProviders(meta *module.Meta, resources []rawResource, config *print.Config) []*Provider { //nolint:gocyclo
	// NOTE(khos2ow): this function is over our cyclomatic complexity goal.
	// Be wary when adding branches, and look for functionality that could
	// be reasonably moved into an injected dependency.

	type provider struct {
		Name        string   `hcl:"name,label"`
		Version     string   `hcl:"version"`
		Constraints *string  `hcl:"constraints"`
		Hashes      []string `hcl:"hashes"`
	}

	type lockfile struct {
		Provider []provider `hcl:"provider,block"`
	}
	lock := make(map[string]provider)

	if config.Settings.LockFile {
		var lf lockfile
		filename := filepath.Join(config.ModuleRoot, ".terraform.lock.hcl")
		if err := hclsimple.DecodeFile(filename, nil, &lf); err == nil {
			for i := range lf.Provider {
				segments := strings.Split(lf.Provider[i].Name, "/")
				name := segments[len(segments)-1]
				lock[name] = lf.Provider[i]
			}
		}
	}

	// Reverse map: tfaddr.Provider -> local name
	localNames := make(map[tfaddr.Provider]string, len(meta.ProviderReferences))
	for ref, address := range meta.ProviderReferences {
		if ref.Alias == "" {
			localNames[address] = ref.LocalName
		}
	}

	// Helper to look up constraints string by local name
	constraintFor := func(localName string) string {
		for address, constraint := range meta.ProviderRequirements {
			if localNames[address] == localName && len(constraint) > 0 {
				return constraint.String()
			}
		}
		return ""
	}

	discovered := make(map[string]*Provider)

	for _, resource := range resources {
		comments := loadComments(resource.Filename, resource.Line)

		// Skip over resources that are marked as being ignored
		if strings.Contains(comments, "terraform-docs-ignore") {
			continue
		}

		version := ""
		if provider, ok := lock[resource.Name]; ok {
			version = provider.Version
		} else {
			version = constraintFor(resource.Name)
		}

		key := fmt.Sprintf("%s.%s", &resource.Name, resource.ProviderAlias)
		if existing, ok := discovered[key]; ok {
			if resource.Filename < existing.Position.Filename ||
				(resource.Filename == existing.Position.Filename && resource.Line < existing.Position.Line) {
				existing.Position = Position{
					Filename: resource.Filename,
					Line:     resource.Line,
				}
			}
			continue
		}

		discovered[key] = &Provider{
			Name:    resource.ProviderName,
			Alias:   types.String(resource.ProviderAlias),
			Version: types.String(version),
			Position: Position{
				Filename: resource.Filename,
				Line:     resource.Line,
			},
		}
	}

	keys := make([]string, 0, len(discovered))
	for key := range discovered {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	providers := make([]*Provider, 0, len(discovered))
	for _, key := range keys {
		providers = append(providers, discovered[key])
	}

	return providers
}

func loadRequirements(tfmodule *tfconfig.Module) []*Requirement {
	var requirements = make([]*Requirement, 0)
	for _, core := range tfmodule.RequiredCore {
		requirements = append(requirements, &Requirement{
			Name:    "terraform",
			Version: types.String(core),
		})
	}

	names := make([]string, 0, len(tfmodule.RequiredProviders))
	for n := range tfmodule.RequiredProviders {
		names = append(names, n)
	}

	sort.Strings(names)

	for _, name := range names {
		for _, version := range tfmodule.RequiredProviders[name].VersionConstraints {
			requirements = append(requirements, &Requirement{
				Name:    name,
				Version: types.String(version),
			})
		}
	}
	return requirements
}

func loadResources(tfmodule *tfconfig.Module, config *print.Config) []*Resource {
	allResources := []map[string]*tfconfig.Resource{tfmodule.ManagedResources, tfmodule.DataResources}
	discovered := make(map[string]*Resource)

	for _, resource := range allResources {
		for _, r := range resource {
			comments := loadComments(r.Pos.Filename, r.Pos.Line)

			// skip over resources that are marked as being ignored
			if strings.Contains(comments, "terraform-docs-ignore") {
				continue
			}

			var version string
			if rv, ok := tfmodule.RequiredProviders[r.Provider.Name]; ok {
				version = resourceVersion(rv.VersionConstraints)
			}

			var source string
			if len(tfmodule.RequiredProviders[r.Provider.Name].Source) > 0 {
				source = tfmodule.RequiredProviders[r.Provider.Name].Source
			} else {
				source = fmt.Sprintf("%s/%s", "hashicorp", r.Provider.Name)
			}

			rType := strings.TrimPrefix(r.Type, r.Provider.Name+"_")
			key := fmt.Sprintf("%s.%s.%s.%s", r.Provider.Name, r.Mode, rType, r.Name)

			description := ""
			if config.Settings.ReadComments {
				description = comments
			}

			discovered[key] = &Resource{
				Type:           rType,
				Name:           r.Name,
				Mode:           r.Mode.String(),
				ProviderName:   r.Provider.Name,
				ProviderSource: source,
				Version:        types.String(version),
				Description:    types.String(description),
				Position: Position{
					Filename: r.Pos.Filename,
					Line:     r.Pos.Line,
				},
			}
		}
	}

	resourceKeys := make([]string, 0, len(discovered))
	for key := range discovered {
		resourceKeys = append(resourceKeys, key)
	}
	sort.Strings(resourceKeys)

	resources := make([]*Resource, 0, len(discovered))
	for _, key := range resourceKeys {
		resources = append(resources, discovered[key])
	}
	return resources
}

func resourceVersion(constraints []string) string {
	if len(constraints) == 0 {
		return "latest"
	}
	versionParts := strings.Split(constraints[len(constraints)-1], " ")
	switch len(versionParts) {
	case 1:
		if _, err := strconv.Atoi(versionParts[0][0:1]); err != nil {
			if versionParts[0][0:1] == "=" {
				return versionParts[0][1:]
			}
			return "latest"
		}
		return versionParts[0]
	case 2:
		if versionParts[0] == "=" {
			return versionParts[1]
		}
	}
	return "latest"
}

func loadComments(filename string, lineNum int) string {
	lines := reader.Lines{
		FileName: filename,
		LineNum:  lineNum,
		Condition: func(line string) bool {
			return strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//")
		},
		Parser: func(line string) (string, bool) {
			line = strings.TrimSpace(line)
			line = strings.TrimPrefix(line, "#")
			line = strings.TrimPrefix(line, "//")
			line = strings.TrimSpace(line)
			return line, true
		},
	}
	comment, err := lines.Extract()
	if err != nil {
		return "" // absorb the error, we don't need to bubble it up or break the execution
	}
	return strings.Join(comment, " ")
}

func sortItems(tfmodule *Module, config *print.Config) {
	// inputs
	inputs(tfmodule.Inputs).sort(config.Sort.Enabled, config.Sort.By)
	inputs(tfmodule.RequiredInputs).sort(config.Sort.Enabled, config.Sort.By)
	inputs(tfmodule.OptionalInputs).sort(config.Sort.Enabled, config.Sort.By)

	// outputs
	outputs(tfmodule.Outputs).sort(config.Sort.Enabled, config.Sort.By)

	// providers
	providers(tfmodule.Providers).sort(config.Sort.Enabled, config.Sort.By)

	// resources
	resources(tfmodule.Resources).sort(config.Sort.Enabled, config.Sort.By)

	// modules
	modulecalls(tfmodule.ModuleCalls).sort(config.Sort.Enabled, config.Sort.By)
}
