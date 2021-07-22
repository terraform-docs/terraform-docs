/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package terraform

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsimple"

	terraformsdk "github.com/terraform-docs/plugin-sdk/terraform"
	"github.com/terraform-docs/terraform-config-inspect/tfconfig"
	"github.com/terraform-docs/terraform-docs/internal/reader"
	"github.com/terraform-docs/terraform-docs/internal/types"
)

// Module represents a Terraform module. It consists of
//
// - Header       ('header' json key):        Module header found in shape of multi line '*.tf' comments or an entire file
// - Footer       ('footer' json key):        Module footer found in shape of multi line '*.tf' comments or an entire file
// - Inputs       ('inputs' json key):        List of input 'variables' extracted from the Terraform module .tf files
// - ModuleCalls  ('modules' json key):       List of 'modules' extracted from the Terraform module .tf files
// - Outputs      ('outputs' json key):       List of 'outputs' extracted from Terraform module .tf files
// - Providers    ('providers' json key):     List of 'providers' extracted from resources used in Terraform module
// - Requirements ('requirements' json key):  List of 'requirements' extracted from the Terraform module .tf files
// - Resources    ('resources' json key):     List of 'resources' extracted from the Terraform module .tf files
type Module struct {
	XMLName xml.Name `json:"-" toml:"-" xml:"module" yaml:"-"`

	Header       string         `json:"header" toml:"header" xml:"header" yaml:"header"`
	Footer       string         `json:"footer" toml:"footer" xml:"footer" yaml:"footer"`
	Inputs       []*Input       `json:"inputs" toml:"inputs" xml:"inputs>input" yaml:"inputs"`
	ModuleCalls  []*ModuleCall  `json:"modules" toml:"modules" xml:"modules>module" yaml:"modules"`
	Outputs      []*Output      `json:"outputs" toml:"outputs" xml:"outputs>output" yaml:"outputs"`
	Providers    []*Provider    `json:"providers" toml:"providers" xml:"providers>provider" yaml:"providers"`
	Requirements []*Requirement `json:"requirements" toml:"requirements" xml:"requirements>requirement" yaml:"requirements"`
	Resources    []*Resource    `json:"resources" toml:"resources" xml:"resources>resource" yaml:"resources"`

	RequiredInputs []*Input `json:"-" toml:"-" xml:"-" yaml:"-"`
	OptionalInputs []*Input `json:"-" toml:"-" xml:"-" yaml:"-"`
}

// HasHeader indicates if the module has header.
func (m *Module) HasHeader() bool {
	return len(m.Header) > 0
}

// HasFooter indicates if the module has footer.
func (m *Module) HasFooter() bool {
	return len(m.Footer) > 0
}

// HasInputs indicates if the module has inputs.
func (m *Module) HasInputs() bool {
	return len(m.Inputs) > 0
}

// HasModuleCalls indicates if the module has modulecalls.
func (m *Module) HasModuleCalls() bool {
	return len(m.ModuleCalls) > 0
}

// HasOutputs indicates if the module has outputs.
func (m *Module) HasOutputs() bool {
	return len(m.Outputs) > 0
}

// HasProviders indicates if the module has providers.
func (m *Module) HasProviders() bool {
	return len(m.Providers) > 0
}

// HasRequirements indicates if the module has requirements.
func (m *Module) HasRequirements() bool {
	return len(m.Requirements) > 0
}

// HasResources indicates if the module has resources.
func (m *Module) HasResources() bool {
	return len(m.Resources) > 0
}

// Convert internal Module to its equivalent in plugin-sdk
func (m *Module) Convert() terraformsdk.Module {
	return terraformsdk.NewModule(
		terraformsdk.WithHeader(m.Header),
		terraformsdk.WithFooter(m.Footer),
		terraformsdk.WithInputs(inputs(m.Inputs).convert()),
		terraformsdk.WithModuleCalls(modulecalls(m.ModuleCalls).convert()),
		terraformsdk.WithOutputs(outputs(m.Outputs).convert()),
		terraformsdk.WithProviders(providers(m.Providers).convert()),
		terraformsdk.WithRequirements(requirements(m.Requirements).convert()),
		terraformsdk.WithResources(resources(m.Resources).convert()),
		terraformsdk.WithRequiredInputs(inputs(m.RequiredInputs).convert()),
		terraformsdk.WithOptionalInputs(inputs(m.OptionalInputs).convert()),
	)
}

// LoadWithOptions returns new instance of Module with all the inputs and
// outputs discovered from provided 'path' containing Terraform config
func LoadWithOptions(options *Options) (*Module, error) {
	tfmodule, err := loadModule(options.Path)
	if err != nil {
		return nil, err
	}

	module, err := loadModuleItems(tfmodule, options)
	if err != nil {
		return nil, err
	}
	sortItems(module, options.SortBy)
	return module, nil
}

func loadModule(path string) (*tfconfig.Module, error) {
	module, diag := tfconfig.LoadModule(path)
	if diag != nil && diag.HasErrors() {
		return nil, diag
	}
	return module, nil
}

func loadModuleItems(tfmodule *tfconfig.Module, options *Options) (*Module, error) {
	header, err := loadHeader(options)
	if err != nil {
		return nil, err
	}

	footer, err := loadFooter(options)
	if err != nil {
		return nil, err
	}

	inputs, required, optional := loadInputs(tfmodule)
	modulecalls := loadModulecalls(tfmodule)
	outputs, err := loadOutputs(tfmodule, options)
	if err != nil {
		return nil, err
	}
	providers := loadProviders(tfmodule, options)
	requirements := loadRequirements(tfmodule)
	resources := loadResources(tfmodule)

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
	case ".adoc", ".md", ".tf", ".txt":
		return true, nil
	}
	return false, fmt.Errorf("only .adoc, .md, .tf, and .txt formats are supported to read %s from", section)
}

func loadHeader(options *Options) (string, error) {
	if !options.ShowHeader {
		return "", nil
	}
	return loadSection(options, options.HeaderFromFile, "header")
}

func loadFooter(options *Options) (string, error) {
	if !options.ShowFooter {
		return "", nil
	}
	return loadSection(options, options.FooterFromFile, "footer")
}

func loadSection(options *Options, file string, section string) (string, error) { //nolint:gocyclo
	// NOTE(khos2ow): this function is over our cyclomatic complexity goal.
	// Be wary when adding branches, and look for functionality that could
	// be reasonably moved into an injected dependency.

	if section == "" {
		return "", errors.New("section is missing")
	}
	filename := filepath.Join(options.Path, file)
	if ok, err := isFileFormatSupported(file, section); !ok {
		return "", err
	}
	if info, err := os.Stat(filename); os.IsNotExist(err) || info.IsDir() {
		if section == "header" && file == "main.tf" {
			return "", nil // absorb the error to not break workflow for default value of header and missing 'main.tf'
		}
		return "", err // user explicitly asked for a file which doesn't exist
	}
	if getFileFormat(file) != ".tf" {
		content, err := ioutil.ReadFile(filepath.Clean(filename))
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

func loadInputs(tfmodule *tfconfig.Module) ([]*Input, []*Input, []*Input) {
	var inputs = make([]*Input, 0, len(tfmodule.Variables))
	var required = make([]*Input, 0, len(tfmodule.Variables))
	var optional = make([]*Input, 0, len(tfmodule.Variables))

	for _, input := range tfmodule.Variables {
		// convert CRLF to LF early on (https://github.com/terraform-docs/terraform-docs/issues/305)
		inputDescription := strings.ReplaceAll(input.Description, "\r\n", "\n")
		if inputDescription == "" {
			inputDescription = loadComments(input.Pos.Filename, input.Pos.Line)
		}

		i := &Input{
			Name:        input.Name,
			Type:        types.TypeOf(input.Type, input.Default),
			Description: types.String(inputDescription),
			Default:     types.ValueOf(input.Default),
			Required:    input.Required,
			Position: Position{
				Filename: input.Pos.Filename,
				Line:     input.Pos.Line,
			},
		}

		inputs = append(inputs, i)

		if i.HasDefault() {
			optional = append(optional, i)
		} else {
			required = append(required, i)
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

func loadModulecalls(tfmodule *tfconfig.Module) []*ModuleCall {
	var modules = make([]*ModuleCall, 0)
	var source, version string

	for _, m := range tfmodule.ModuleCalls {
		source, version = formatSource(m.Source, m.Version)
		modules = append(modules, &ModuleCall{
			Name:    m.Name,
			Source:  source,
			Version: version,
			Position: Position{
				Filename: m.Pos.Filename,
				Line:     m.Pos.Line,
			},
		})
	}
	return modules
}

func loadOutputs(tfmodule *tfconfig.Module, options *Options) ([]*Output, error) {
	outputs := make([]*Output, 0, len(tfmodule.Outputs))
	values := make(map[string]*output)
	if options.OutputValues {
		var err error
		values, err = loadOutputValues(options)
		if err != nil {
			return nil, err
		}
	}
	for _, o := range tfmodule.Outputs {
		description := o.Description
		if description == "" {
			description = loadComments(o.Pos.Filename, o.Pos.Line)
		}
		output := &Output{
			Name:        o.Name,
			Description: types.String(description),
			Position: Position{
				Filename: o.Pos.Filename,
				Line:     o.Pos.Line,
			},
			ShowValue: options.OutputValues,
		}
		if options.OutputValues {
			output.Sensitive = values[output.Name].Sensitive
			if values[output.Name].Sensitive {
				output.Value = types.ValueOf(`<sensitive>`)
			} else {
				output.Value = types.ValueOf(values[output.Name].Value)
			}
		}
		outputs = append(outputs, output)
	}
	return outputs, nil
}

func loadOutputValues(options *Options) (map[string]*output, error) {
	var out []byte
	var err error
	if options.OutputValuesPath == "" {
		cmd := exec.Command("terraform", "output", "-json")
		cmd.Dir = options.Path
		if out, err = cmd.Output(); err != nil {
			return nil, fmt.Errorf("caught error while reading the terraform outputs: %w", err)
		}
	} else if out, err = ioutil.ReadFile(options.OutputValuesPath); err != nil {
		return nil, fmt.Errorf("caught error while reading the terraform outputs file at %s: %w", options.OutputValuesPath, err)
	}
	var terraformOutputs map[string]*output
	err = json.Unmarshal(out, &terraformOutputs)
	if err != nil {
		return nil, err
	}
	return terraformOutputs, err
}

func loadProviders(tfmodule *tfconfig.Module, options *Options) []*Provider {
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

	if options.UseLockFile {
		var lf lockfile

		filename := filepath.Join(options.Path, ".terraform.lock.hcl")
		if err := hclsimple.DecodeFile(filename, nil, &lf); err == nil {
			for i := range lf.Provider {
				segments := strings.Split(lf.Provider[i].Name, "/")
				name := segments[len(segments)-1]
				lock[name] = lf.Provider[i]
			}
		}
	}

	resources := []map[string]*tfconfig.Resource{tfmodule.ManagedResources, tfmodule.DataResources}
	discovered := make(map[string]*Provider)

	for _, resource := range resources {
		for _, r := range resource {
			var version = ""
			if l, ok := lock[r.Provider.Name]; ok {
				version = l.Version
			} else if rv, ok := tfmodule.RequiredProviders[r.Provider.Name]; ok && len(rv.VersionConstraints) > 0 {
				version = strings.Join(rv.VersionConstraints, " ")
			}

			key := fmt.Sprintf("%s.%s", r.Provider.Name, r.Provider.Alias)
			discovered[key] = &Provider{
				Name:    r.Provider.Name,
				Alias:   types.String(r.Provider.Alias),
				Version: types.String(version),
				Position: Position{
					Filename: r.Pos.Filename,
					Line:     r.Pos.Line,
				},
			}
		}
	}

	providers := make([]*Provider, 0, len(discovered))
	for _, provider := range discovered {
		providers = append(providers, provider)
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

func loadResources(tfmodule *tfconfig.Module) []*Resource {
	allResources := []map[string]*tfconfig.Resource{tfmodule.ManagedResources, tfmodule.DataResources}
	discovered := make(map[string]*Resource)

	for _, resource := range allResources {
		for _, r := range resource {
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
			discovered[key] = &Resource{
				Type:           rType,
				Name:           r.Name,
				Mode:           r.Mode.String(),
				ProviderName:   r.Provider.Name,
				ProviderSource: source,
				Version:        types.String(version),
				Position: Position{
					Filename: r.Pos.Filename,
					Line:     r.Pos.Line,
				},
			}
		}
	}

	resources := make([]*Resource, 0, len(discovered))
	for _, resource := range discovered {
		resources = append(resources, resource)
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

func sortItems(tfmodule *Module, sortby *SortBy) { //nolint:gocyclo
	// NOTE(khos2ow): this function is over our cyclomatic complexity goal.
	// Be wary when adding branches, and look for functionality that could
	// be reasonably moved into an injected dependency.

	// inputs
	switch {
	case sortby.Type:
		sort.Sort(inputsSortedByType(tfmodule.Inputs))
		sort.Sort(inputsSortedByType(tfmodule.RequiredInputs))
		sort.Sort(inputsSortedByType(tfmodule.OptionalInputs))
	case sortby.Required:
		sort.Sort(inputsSortedByRequired(tfmodule.Inputs))
		sort.Sort(inputsSortedByRequired(tfmodule.RequiredInputs))
		sort.Sort(inputsSortedByRequired(tfmodule.OptionalInputs))
	case sortby.Name:
		sort.Sort(inputsSortedByName(tfmodule.Inputs))
		sort.Sort(inputsSortedByName(tfmodule.RequiredInputs))
		sort.Sort(inputsSortedByName(tfmodule.OptionalInputs))
	default:
		sort.Sort(inputsSortedByPosition(tfmodule.Inputs))
		sort.Sort(inputsSortedByPosition(tfmodule.RequiredInputs))
		sort.Sort(inputsSortedByPosition(tfmodule.OptionalInputs))
	}

	// outputs
	if sortby.Name || sortby.Required || sortby.Type {
		sort.Sort(outputsSortedByName(tfmodule.Outputs))
	} else {
		sort.Sort(outputsSortedByPosition(tfmodule.Outputs))
	}

	// providers
	if sortby.Name || sortby.Required || sortby.Type {
		sort.Sort(providersSortedByName(tfmodule.Providers))
	} else {
		sort.Sort(providersSortedByPosition(tfmodule.Providers))
	}

	// resources (always sorted)
	sort.Sort(resourcesSortedByType(tfmodule.Resources))

	// modules
	switch {
	case sortby.Name || sortby.Required:
		sort.Sort(modulecallsSortedByName(tfmodule.ModuleCalls))
	case sortby.Type:
		sort.Sort(modulecallsSortedBySource(tfmodule.ModuleCalls))
	default:
		sort.Sort(modulecallsSortedByPosition(tfmodule.ModuleCalls))
	}
}
