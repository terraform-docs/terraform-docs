package terraform

import "github.com/hashicorp/hcl/v2"

func extractBlockPositions(files map[string]*hcl.File, blockType string) map[string]Position {
	output := make(map[string]Position)
	schema := &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type:       blockType,
				LabelNames: []string{"names"},
			},
		},
	}
	for name, file := range files {
		content, _, _ := file.Body.PartialContent(schema)
		for _, block := range content.Blocks {
			if len(block.Labels) == 0 {
				continue
			}

			label := block.Labels[0]

			if _, exists := output[label]; exists {
				continue
			}

			output[label] = Position{
				Filename: name,
				Line:     block.DefRange.Start.Line,
			}
		}
	}
	return output
}
