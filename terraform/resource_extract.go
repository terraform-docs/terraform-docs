package terraform

import "github.com/hashicorp/hcl/v2"

type rawResource struct {
	Mode     string // "managed" || "data"
	Type     string
	Name     string
	Filename string
	Line     int
}

func extractResources(files map[string]*hcl.File) []rawResource {
	var out []rawResource

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

	for name, file := range files {
		content, _, _ := file.Body.PartialContent(schema)
		for _, block := range content.Blocks {
			var mode string = "managed"
			if block.Type == "data" {
				mode = "data"
			}
			out = append(out, rawResource{
				Mode:     mode,
				Type:     block.Labels[0],
				Name:     block.Labels[1],
				Filename: name,
				Line:     block.DefRange.Start.Line,
			})
		}
	}
	return out
}
