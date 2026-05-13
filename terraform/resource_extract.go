/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package terraform

import (
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

type rawResource struct {
	Mode          string // "managed" || "data"
	Type          string
	Name          string
	Filename      string
	Line          int
	ProviderName  string
	ProviderAlias string
}

func extractResources(files map[string]*hcl.File) []rawResource {
	var out []rawResource

	bodySchema := &hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{
			{
				Name: "provider",
			},
		},
	}

	schema := &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
			},
			{
				Type:       "data",
				LabelNames: []string{"type", "name"},
			},
		},
	}

	for _, file := range files {
		content, _, _ := file.Body.PartialContent(schema)
		for _, block := range content.Blocks {
			out = append(out, buildResource(block, bodySchema))
		}
	}
	return out
}

func buildResource(block *hcl.Block, bodySchema *hcl.BodySchema) rawResource {
	mode := "managed"
	if block.Type == "data" {
		mode = "data"
	}

	providerName := strings.SplitN(block.Labels[0], "_", 2)[0]
	providerAlias := ""

	inner, _, _ := block.Body.PartialContent(bodySchema)
	if attribute, ok := inner.Attributes["provider"]; ok {
		providerName, providerAlias = parseProvider(attribute, providerName)
	}

	return rawResource{
		Mode:          mode,
		Type:          block.Labels[0],
		Name:          block.Labels[1],
		Filename:      block.DefRange.Filename,
		Line:          block.DefRange.Start.Line,
		ProviderName:  providerName,
		ProviderAlias: providerAlias,
	}
}

func parseProvider(attribute *hcl.Attribute, defaultName string) (string, string) {
	// provider = aws.useast1 -> Traversal: [aws, useast1]
	if traversal, diags := hcl.AbsTraversalForExpr(attribute.Expr); !diags.HasErrors() {
		return providerFromTraversal(traversal, defaultName)
	}
	// provider = "aws.useast1" (legacy string-literal form)
	if value, diags := attribute.Expr.Value(nil); !diags.HasErrors() && value.Type() == cty.String {
		return providerFromString(value.AsString(), defaultName)
	}
	return defaultName, ""
}

func providerFromTraversal(traversal hcl.Traversal, defaultName string) (string, string) {
	name := defaultName
	alias := ""
	if len(traversal) >= 1 {
		if root, ok := traversal[0].(hcl.TraverseRoot); ok {
			name = root.Name
		}
	}
	if len(traversal) >= 2 {
		if attr, ok := traversal[1].(hcl.TraverseAttr); ok {
			alias = attr.Name
		}
	}
	return name, alias
}

func providerFromString(value string, _ string) (string, string) {
	parts := strings.SplitN(value, ".", 2)
	name := parts[0]
	alias := ""
	if len(parts) == 2 {
		alias = parts[1]
	}
	return name, alias
}
