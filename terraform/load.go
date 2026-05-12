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

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/ext/typeexpr"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/opentofu/opentofu-schema/earlydecoder"
	"github.com/opentofu/opentofu-schema/module"
	tfaddr "github.com/opentofu/registry-address"
	"github.com/zclconf/go-cty/cty"
	ctyjson "github.com/zclconf/go-cty/cty/json"

	"github.com/terraform-docs/terraform-docs/internal/reader"
	"github.com/terraform-docs/terraform-docs/internal/types"
	"github.com/terraform-docs/terraform-docs/print"
)

const ignoreMarker = "terraform-docs-ignore"

// LoadWithOptions returns new instance of Module with all the inputs and
// outputs discovered from provided 'path' containing Terraform config
func LoadWithOptions(config *print.Config) (*Module, error) {
	meta, files, err := loadModule(config.ModuleRoot)
	if err != nil {
		return nil, err
	}

	module, err := loadModuleItems(meta, files, config)
	if err != nil {
		return nil, err
	}

	sortItems(module, config)
	return module, nil
}

func loadModule(path string) (*module.Meta, map[string]*hcl.File, error) {
	files, diags := parseModuleFiles(path)
	if diags.HasErrors() {
		return nil, nil, diags
	}
	// earlydecoder may emit diagnostics for constructs it cannot fully
	// resolve (e.g. legacy `type = "string"`, `var.foo` references in
	// places it does not evaluate, etc.). These are non-fatal for our
	// metadata extraction: it still populates meta with what it could
	// parse. Mirror the lenient behavior of the previous
	// terraform-config-inspect loader and only fail if no metadata at
	// all came back.meta, _ := earlydecoder.LoadModule(path, files)
	meta, _ := earlydecoder.LoadModule(path, files)
	if meta == nil {
		return nil, nil, fmt.Errorf("failed to load module from %q", path)
	}
	return meta, files, nil
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

func loadModuleItems(meta *module.Meta, files map[string]*hcl.File, config *print.Config) (*Module, error) {
	header, err := loadHeader(config)
	if err != nil {
		return nil, err
	}

	footer, err := loadFooter(config)
	if err != nil {
		return nil, err
	}

	variablePositions := extractBlockPositions(files, "variables")
	outputPositions := extractBlockPositions(files, "output")
	rawResources := extractResources(files)
	inputs, required, optional := loadInputs(meta, variablePositions, config)
	moduleCalls := loadModuleCalls(meta, files, config)
	outputs, err := loadOutputs(meta, outputPositions, config)
	if err != nil {
		return nil, err
	}
	providers := loadProviders(meta, rawResources, config)
	requirements := loadRequirements(meta)
	resources := loadResources(meta, rawResources, config)

	return &Module{
		Header:         header,
		Footer:         footer,
		Inputs:         inputs,
		ModuleCalls:    moduleCalls,
		Outputs:        outputs,
		Providers:      providers,
		Requirements:   requirements,
		Resources:      resources,
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

	for name := range meta.Variables {
		input := meta.Variables[name]
		position := positions[name]
		comments, ignored := isIgnored(position.Filename, position.Line)

		if ignored {
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

func loadModuleCalls(meta *module.Meta, files map[string]*hcl.File, config *print.Config) []*ModuleCall {
	modules := make([]*ModuleCall, 0, len(meta.ModuleCalls))

	versionOverrides := resolveModuleVersions(files, meta)

	for localName := range meta.ModuleCalls {
		moduleCall := meta.ModuleCalls[localName]
		var filename string
		var line int

		if moduleCall.RangePtr != nil {
			filename = moduleCall.RangePtr.Filename
			line = moduleCall.RangePtr.Start.Line
		}

		comments, ignored := isIgnored(filename, line)
		if ignored {
			continue
		}

		description := ""
		if config.Settings.ReadComments {
			description = comments
		}

		source, version := moduleSourceAndVersion(&moduleCall)
		if version == "" {
			if v, ok := versionOverrides[localName]; ok && v != "" {
				version = v
			}
		}

		modules = append(modules, &ModuleCall{
			Name:        localName,
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
func moduleSourceAndVersion(moduleCall *module.DeclaredModuleCall) (string, string) {
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
		return formatSource(moduleCall.RawSourceAddr, declaredVersion)
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

	for name := range meta.Outputs {
		output := meta.Outputs[name]
		position := positions[name]
		comments, ignored := isIgnored(position.Filename, position.Line)

		if ignored {
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
			for index := range lf.Provider {
				segments := strings.Split(lf.Provider[index].Name, "/")
				name := segments[len(segments)-1]
				lock[name] = lf.Provider[index]
			}
		}
	}

	// Reverse map: tfaddr.Provider -> local name
	localNames := make(map[tfaddr.Provider]string, len(meta.ProviderReferences))
	for reference := range meta.ProviderReferences {
		if reference.Alias == "" {
			localNames[meta.ProviderReferences[reference]] = reference.LocalName
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

	for index := range resources {
		resource := &resources[index]
		_, ignored := isIgnored(resource.Filename, resource.Line)

		if ignored {
			continue
		}

		version := ""
		if provider, ok := lock[resource.ProviderName]; ok {
			version = provider.Version
		} else {
			version = constraintFor(resource.ProviderName)
		}

		key := fmt.Sprintf("%s.%s", resource.ProviderName, resource.ProviderAlias)
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

func loadRequirements(meta *module.Meta) []*Requirement {
	requirements := make([]*Requirement, 0)

	type providerEntry struct {
		localName string
		address   tfaddr.Provider
	}

	// terraform / opentofu core
	for _, coreConstraint := range meta.CoreRequirements {
		requirements = append(requirements, &Requirement{
			Name:    "terraform",
			Version: types.String(coreConstraint.String()),
		})
	}

	// reverse map: tfaddr.Provider -> local name (un-aliased ref)
	localNamesByProvider := make(map[tfaddr.Provider]string, len(meta.ProviderReferences))
	for providerReference := range meta.ProviderReferences {
		if providerReference.Alias == "" {
			localNamesByProvider[meta.ProviderReferences[providerReference]] = providerReference.LocalName
		}
	}

	// stable ordering by local name
	providerEntries := make([]providerEntry, 0, len(meta.ProviderRequirements))
	for providerAddr := range meta.ProviderRequirements {
		localName, hasLocalName := localNamesByProvider[providerAddr]
		if !hasLocalName {
			localName = providerAddr.Type // fallback: bare type
		}
		providerEntries = append(providerEntries, providerEntry{localName: localName, address: providerAddr})
	}
	sort.Slice(providerEntries, func(i, j int) bool {
		return providerEntries[i].localName < providerEntries[j].localName
	})

	for index := range providerEntries {
		entry := &providerEntries[index]
		for _, versionConstraint := range meta.ProviderRequirements[entry.address] {
			requirements = append(requirements, &Requirement{
				Name:    entry.localName,
				Version: types.String(versionConstraint.String()),
			})
		}
	}
	return requirements
}

func loadResources(meta *module.Meta, resources []rawResource, config *print.Config) []*Resource {
	// build localName -> tfaddr.Provider once
	localToAttribute := make(map[string]tfaddr.Provider)
	for reference := range meta.ProviderReferences {
		if reference.Alias == "" {
			localToAttribute[reference.LocalName] = meta.ProviderReferences[reference]
		}
	}

	discovered := make(map[string]*Resource)
	for index := range resources {
		resource := &resources[index]
		comments, ignored := isIgnored(resource.Filename, resource.Line)

		if ignored {
			continue
		}

		source, version := resolveProviderSource(resource, meta, localToAttribute)

		resourceType := strings.TrimPrefix(resource.Type, resource.ProviderName+"_")
		key := fmt.Sprintf("%s.%s.%s.%s", resource.ProviderName, resource.Mode, resourceType, resource.Name)
		description := ""
		if config.Settings.ReadComments {
			description = comments
		}

		discovered[key] = &Resource{
			Type:           resourceType,
			Name:           resource.Name,
			Mode:           resource.Mode,
			ProviderName:   resource.ProviderName,
			ProviderSource: source,
			Version:        types.String(version),
			Description:    types.String(description),
			Position: Position{
				Filename: resource.Filename,
				Line:     resource.Line,
			},
		}

	}

	resourceKeys := make([]string, 0, len(discovered))
	for key := range discovered {
		resourceKeys = append(resourceKeys, key)
	}
	sort.Strings(resourceKeys)

	allResources := make([]*Resource, 0, len(discovered))
	for _, key := range resourceKeys {
		allResources = append(allResources, discovered[key])
	}
	return allResources
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

func resolveProviderSource(resource *rawResource, meta *module.Meta, localToAttribute map[string]tfaddr.Provider) (string, string) {
	attribute, ok := localToAttribute[resource.ProviderName]
	version := ""

	if !ok {
		return fmt.Sprintf("hashicorp/%s", resource.ProviderName), ""
	}

	if vs, ok := meta.ProviderRequirements[attribute]; ok && len(vs) > 0 {
		version = resourceVersion([]string{
			vs.String(),
		})
	}
	return attribute.ForDisplay(), version
}

func isIgnored(filename string, line int) (comments string, ignored bool) {
	comments = loadComments(filename, line)
	return comments, strings.Contains(comments, ignoreMarker)
}
