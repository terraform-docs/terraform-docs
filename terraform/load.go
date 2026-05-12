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

func ctyTypetoString(input cty.Type) string {
	if input == cty.NilType || input == cty.DynamicPseudoType {
		return ""
	}
	return typeexpr.TypeString(input)
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

func ctyValueToTypesValue(value cty.Value) types.Value {
	if value == cty.NilVal {
		return new(types.Nil)
	}
	if value.IsNull() {
		return new(types.Nil)
	}
	var raw interface{}
	if err := json.Unmarshal([]byte(ctyValueToString(value)), &raw); err != nil {
		return new(types.Nil)
	}
	return types.ValueOf(raw)
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

	variablePositions := extractBlockPositions(files, "variable")
	outputPositions := extractBlockPositions(files, "output")
	rawResources := extractResources(files)
	inputs, required, optional := loadInputs(meta, variablePositions, config)
	moduleCalls := loadModuleCalls(meta, files, config)
	outputs, err := loadOutputs(meta, outputPositions, config)
	if err != nil {
		return nil, err
	}
	providers := loadProviders(meta, rawResources, files, config)
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

func loadSection(config *print.Config, file, section string) (string, error) {
	if section == "" {
		return "", errors.New("section is missing")
	}
	if ok, err := isFileFormatSupported(file, section); !ok {
		return "", err
	}

	filename, err := resolveSectionFile(config.ModuleRoot, file, section)
	if err != nil || filename == "" {
		return "", err
	}

	if format := getFileFormat(file); format != ".tf" && format != ".tofu" {
		return readPlainFile(filename)
	}
	return extractTerraformDocComment(filename)
}

// resolveSectionFile returns the absolute path of the section file, or
// ("", nil) when the implicit default header file is missing (which is
// silently allowed to keep the legacy workflow intact).
func resolveSectionFile(root, file, section string) (string, error) {
	filename := filepath.Join(root, file)
	info, err := os.Stat(filename)
	if err == nil && !info.IsDir() {
		return filename, nil
	}
	if section == "header" && file == "main.tf" {
		return "", nil
	}
	return "", err
}

func readPlainFile(filename string) (string, error) {
	content, err := os.ReadFile(filepath.Clean(filename))
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func extractTerraformDocComment(filename string) (string, error) {
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
		defaultValue := ctyValueToTypesValue(input.DefaultValue)

		in := &Input{
			Name:        name,
			Type:        types.TypeOf(ctyTypetoString(input.Type), defaultValue.Raw()),
			Description: types.String(description),
			Default:     defaultValue,
			Required:    isRequired,
			Deprecated:  types.String(input.Deprecated),
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
			Deprecated:  types.String(output.Deprecated),
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

func loadProviders(meta *module.Meta, resources []rawResource, files map[string]*hcl.File, config *print.Config) []*Provider {
	lock := loadProviderLockfile(config)
	constraintFor := buildConstraintLookup(meta)

	discovered := make(map[string]*Provider)

	for _, block := range extractProviderBlocks(files) {
		seedProviderFromBlock(discovered, block, lock, constraintFor)
	}

	for index := range resources {
		addProviderFromResource(discovered, &resources[index], lock, constraintFor)
	}
	return sortedProviders(discovered)
}

func buildConstraintLookup(meta *module.Meta) func(string) string {
	constraints := make(map[string]string, len(meta.ProviderReferences))
	for reference, address := range meta.ProviderReferences {
		if reference.Alias != "" {
			continue
		}
		if cs, ok := meta.ProviderRequirements[address]; ok && len(cs) > 0 {
			constraints[reference.LocalName] = cs.String()
		}
	}
	return func(localName string) string {
		return constraints[localName]
	}
}

func seedProviderFromBlock(
	discovered map[string]*Provider,
	block rawProviderBlock,
	lock map[string]lockedProvider,
	constraintFor func(string) string,
) {
	if _, ignored := isIgnored(block.Filename, block.Line); ignored {
		return
	}

	version := constraintFor(block.LocalName)
	if entry, ok := lock[block.LocalName]; ok {
		version = entry.Version
	}

	position := Position{Filename: block.Filename, Line: block.Line}
	key := fmt.Sprintf("%s.%s", block.LocalName, block.Alias)

	if existing, ok := discovered[key]; ok {
		if position.Filename < existing.Position.Filename ||
			(position.Filename == existing.Position.Filename && position.Line < existing.Position.Line) {
			existing.Position = position
		}
		return
	}

	discovered[key] = &Provider{
		Name:     block.LocalName,
		Alias:    types.String(block.Alias),
		Version:  types.String(version),
		Position: position,
	}
}

func addProviderFromResource(
	discovered map[string]*Provider,
	resource *rawResource,
	lock map[string]lockedProvider,
	constraintFor func(string) string,
) {
	if _, ignored := isIgnored(resource.Filename, resource.Line); ignored {
		return
	}

	version := constraintFor(resource.ProviderName)
	if entry, ok := lock[resource.ProviderName]; ok {
		version = entry.Version
	}

	position := Position{Filename: resource.Filename, Line: resource.Line}
	key := fmt.Sprintf("%s.%s", resource.ProviderName, resource.ProviderAlias)

	if existing, ok := discovered[key]; ok {
		if position.Filename < existing.Position.Filename ||
			(position.Filename == existing.Position.Filename && position.Line < existing.Position.Line) {
			existing.Position = position
		}
		return
	}

	discovered[key] = &Provider{
		Name:     resource.ProviderName,
		Alias:    types.String(resource.ProviderAlias),
		Version:  types.String(version),
		Position: position,
	}
}

func sortedProviders(discovered map[string]*Provider) []*Provider {
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

	if !ok || attribute.Namespace == "-" {
		return fmt.Sprintf("hashicorp/%s", resource.ProviderName), resourceVersion(nil)
	}

	if vs, ok := meta.ProviderRequirements[attribute]; ok && len(vs) > 0 {
		version = resourceVersion([]string{
			vs.String(),
		})
	}
	return attribute.ForDisplay(), version
}


