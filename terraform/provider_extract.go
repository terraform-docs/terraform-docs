package terraform

import (
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/zclconf/go-cty/cty"

	"github.com/terraform-docs/terraform-docs/print"
)

type lockedProvider struct {
	Name        string   `hcl:"name,label"`
	Version     string   `hcl:"version"`
	Constraints *string  `hcl:"constraints"`
	Hashes      []string `hcl:"hashes"`
}

type providerLockfile struct {
	Provider []lockedProvider `hcl:"provider,block"`
}

func loadProviderLockfile(config *print.Config) map[string]lockedProvider {
	if !config.Settings.LockFile {
		return nil
	}
	var lockFile providerLockfile
	filename := filepath.Join(config.ModuleRoot, ".terraform.lock.hcl")
	if err := hclsimple.DecodeFile(filename, nil, &lockFile); err != nil {
		return nil
	}
	out := make(map[string]lockedProvider, len(lockFile.Provider))
	for index := range lockFile.Provider {
		segments := strings.Split(lockFile.Provider[index].Name, "/")
		name := segments[len(segments)-1]
		out[name] = lockFile.Provider[index]
	}
	return out
}

type rawProviderBlock struct {
	LocalName string
	Alias     string
	Filename  string
	Line      int
}

func extractProviderBlocks(files map[string]*hcl.File) []rawProviderBlock {
	schema := &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type:       "provider",
				LabelNames: []string{"name"},
			},
		},
	}

	attributeSchema := &hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{
			{
				Name: "alias",
			},
		},
	}

	var output []rawProviderBlock

	for _, file := range files {
		content, _, _ := file.Body.PartialContent(schema)
		for index := range content.Blocks {
			block := content.Blocks[index]
			alias := ""
			inner, _, _ := block.Body.PartialContent(attributeSchema)

			if attribute, ok := inner.Attributes["alias"]; ok {
				value, _ := attribute.Expr.Value(nil)
				if value.Type() == cty.String {
					alias = value.AsString()
				}
			}
			output = append(output, rawProviderBlock{
				LocalName: block.Labels[0],
				Alias:     alias,
				Filename:  block.DefRange.Filename,
				Line:      block.DefRange.Start.Line,
			})
		}
	}
	return output
}
