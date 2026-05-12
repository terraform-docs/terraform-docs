/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

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
