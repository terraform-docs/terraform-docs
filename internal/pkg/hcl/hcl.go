package hcl

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/buger/jsonparser"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	hclprinter "github.com/hashicorp/hcl/hcl/printer"
)

// ParseAstNode returns Go data structure from an HCL AST node.
func ParseAstNode(node *ast.Node, nodeType string) (interface{}, error) {
	var result interface{}

	// Marshals the HCL Ast Node Object into HCL text
	config := hclprinter.Config{
		SpacesWidth: 2,
	}

	buffer := bytes.NewBufferString("value = ")
	if err := config.Fprint(buffer, *node); err != nil {
		return nil, err
	}

	// Unmarshals HCL Text into a Go data structure
	err := hcl.Unmarshal(buffer.Bytes(), &result)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse HCL: %s", err)
	}

	// Marshals the Go data structure into JSON text
	data, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("Unable to marshal JSON: %s", err)
	}

	// Extract the desired value from JSON text
	path := []string{"value"}
	if nodeType == "map" {
		path = append(path, "[0]")
	}

	data, _, _, err = jsonparser.Get(data, path...)
	if err != nil {
		return nil, fmt.Errorf("Unable to extract value from JSON: %s", err)
	}

	// Unmarshal JSON text into a Go data structure
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("Unable to unmarshal JSON: %s", err)
	}

	return result, nil
}
