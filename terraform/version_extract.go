/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package terraform

import (
	"maps"

	"github.com/hashicorp/hcl/v2"
	"github.com/opentofu/opentofu-schema/module"
	"github.com/zclconf/go-cty/cty"
)

// maxLocalsResolutionPasses caps the iterative passes used when locals
// reference one another. Eight is plenty for any realistic chain.
const maxLocalsResolutionPasses = 8

// resolveModuleVersions walks every `module "..." {}` block in the parsed
// HCL files and tries to evaluate the `version` attribute against an
// EvalContext built from `locals` blocks and variable defaults. This lets
// us recover module versions that the upstream early decoder cannot
// resolve when the value is an expression (e.g. `local.modules.null` or
// `var.module_version`) instead of a string literal.
//
// The returned map is keyed by the module's local name and contains the
// evaluated string value of `version`.
func resolveModuleVersions(files map[string]*hcl.File, meta *module.Meta) map[string]string {
	evalContext := buildModuleEvalContext(files, meta)

	output := make(map[string]string)
	schema := &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{Type: "module", LabelNames: []string{"name"}},
		},
	}
	for _, file := range files {
		content, _, _ := file.Body.PartialContent(schema)
		for _, block := range content.Blocks {
			if version, ok := evalModuleVersion(block, evalContext); ok {
				output[block.Labels[0]] = version
			}
		}
	}
	return output
}

// evalModuleVersion returns the evaluated string value of the `version`
// attribute on a single `module` block, or false if it cannot be resolved.
func evalModuleVersion(block *hcl.Block, evalContext *hcl.EvalContext) (string, bool) {
	if len(block.Labels) == 0 {
		return "", false
	}
	attributes, _ := block.Body.JustAttributes()
	attribute, ok := attributes["version"]
	if !ok {
		return "", false
	}
	value, diags := attribute.Expr.Value(evalContext)
	if diags.HasErrors() || value.IsNull() || !value.IsKnown() {
		return "", false
	}
	if value.Type() != cty.String {
		return "", false
	}
	return value.AsString(), true
}

// buildModuleEvalContext constructs an hcl.EvalContext exposing `local.*`
// (from `locals` blocks) and `var.*` (from variable default values) so we
// can evaluate simple expressions used as module version constraints.
func buildModuleEvalContext(files map[string]*hcl.File, meta *module.Meta) *hcl.EvalContext {
	context := &hcl.EvalContext{
		Variables: map[string]cty.Value{},
	}

	if varValues := collectVariableDefaults(meta); len(varValues) > 0 {
		context.Variables["var"] = cty.ObjectVal(varValues)
	}

	rawLocals := collectLocalExpressions(files)
	if resolved := resolveLocals(rawLocals, context.Variables); len(resolved) > 0 {
		context.Variables["local"] = cty.ObjectVal(resolved)
	}

	return context
}

// collectVariableDefaults returns variable name -> default cty value for
// every variable in the module that has a known default.
func collectVariableDefaults(meta *module.Meta) map[string]cty.Value {
	if meta == nil || len(meta.Variables) == 0 {
		return nil
	}
	values := make(map[string]cty.Value, len(meta.Variables))
	for name := range meta.Variables {
		variable := meta.Variables[name]
		if variable.DefaultValue != cty.NilVal && variable.DefaultValue.IsKnown() {
			values[name] = variable.DefaultValue
		}
	}
	return values
}

// collectLocalExpressions gathers every attribute from every `locals { ... }`
// block across the module's files into a single map of name -> expression.
func collectLocalExpressions(files map[string]*hcl.File) map[string]hcl.Expression {
	rawLocals := map[string]hcl.Expression{}
	schema := &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{{Type: "locals"}},
	}
	for _, file := range files {
		content, _, _ := file.Body.PartialContent(schema)
		for _, block := range content.Blocks {
			attributes, _ := block.Body.JustAttributes()
			for name, attribute := range attributes {
				rawLocals[name] = attribute.Expr
			}
		}
	}
	return rawLocals
}

// resolveLocals iteratively evaluates the local expressions so that
// locals referencing other locals can be resolved. It stops when a pass
// makes no progress or after maxLocalsResolutionPasses iterations.
func resolveLocals(rawLocals map[string]hcl.Expression, baseVars map[string]cty.Value) map[string]cty.Value {
	resolved := map[string]cty.Value{}
	if len(rawLocals) == 0 {
		return resolved
	}
	for index := 0; index < maxLocalsResolutionPasses; index++ {
		passVars := make(map[string]cty.Value, len(baseVars)+1)
		maps.Copy(passVars, baseVars)
		if len(resolved) > 0 {
			passVars["local"] = cty.ObjectVal(resolved)
		}
		passContext := &hcl.EvalContext{Variables: passVars}

		progress := false
		for name, expression := range rawLocals {
			if _, ok := resolved[name]; ok {
				continue
			}
			value, diags := expression.Value(passContext)
			if diags.HasErrors() || !value.IsKnown() {
				continue
			}
			resolved[name] = value
			progress = true
		}
		if !progress {
			break
		}
	}
	return resolved
}
