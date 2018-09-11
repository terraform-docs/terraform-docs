package doc

import (
	"bytes"
	"fmt"
	"path"
	"sort"
	"strconv"
	"strings"

	"encoding/json"

	"github.com/buger/jsonparser"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	hclprinter "github.com/hashicorp/hcl/hcl/printer"
)

// Input represents a terraform input variable.
type Input struct {
	Name        string
	Description string
	Default     *Value
	Type        string
}

// Value returns the default value as an interface{} type.
func (i *Input) Value() interface{} {
	if i.Default != nil {
		switch i.Default.Type {
		case "string":
			return i.Default.Value
		case "map":
			if i.Default.Value != "" {
				return i.Default.Value
			}
			return "{}"
		case "list":
			if i.Default.Value != "" {
				return i.Default.Value
			}
			return "[]"
		}
	}

	return "required"
}

// Value represents a terraform value.
type Value struct {
	Type  string
	Value interface{}
}

// Output represents a terraform output.
type Output struct {
	Name        string
	Description string
}

// Doc represents a terraform module doc.
type Doc struct {
	Comment string
	Inputs  []Input
	Outputs []Output
}

type inputsByName []Input

func (a inputsByName) Len() int           { return len(a) }
func (a inputsByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a inputsByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type outputsByName []Output

func (a outputsByName) Len() int           { return len(a) }
func (a outputsByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a outputsByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

// Create creates a new *Doc from the supplied map
// of filenames and *ast.File.
func Create(files map[string]*ast.File) *Doc {
	doc := new(Doc)

	for name, f := range files {
		list := f.Node.(*ast.ObjectList)
		doc.Inputs = append(doc.Inputs, inputs(list)...)
		doc.Outputs = append(doc.Outputs, outputs(list)...)

		filename := path.Base(name)
		comments := f.Comments

		if filename == "main.tf" && len(comments) > 0 {
			doc.Comment = header(comments[0])
		}
	}
	sort.Sort(inputsByName(doc.Inputs))
	sort.Sort(outputsByName(doc.Outputs))
	return doc
}

// Inputs returns all variables from `list`.
func inputs(list *ast.ObjectList) []Input {
	var ret []Input

	for _, item := range list.Items {
		if is(item, "variable") {
			name, _ := strconv.Unquote(item.Keys[1].Token.Text)
			if name == "" {
				name = item.Keys[1].Token.Text
			}
			items := item.Val.(*ast.ObjectType).List.Items
			var desc string
			switch {
			case description(items) != "":
				desc = description(items)
			case item.LeadComment != nil:
				desc = comment(item.LeadComment.List)
			}

			var itemsType = get(items, "type")
			var itemType string
			def := get(items, "default")

			if itemsType == nil || itemsType.Value == "" {
				if def == nil {
					itemType = "string"
				} else {
					itemType = def.Type
				}
			} else {
				itemType = itemsType.Value.(string)
			}

			ret = append(ret, Input{
				Name:        name,
				Description: desc,
				Default:     def,
				Type:        itemType,
			})
		}
	}

	return ret
}

// Outputs returns all outputs from `list`.
func outputs(list *ast.ObjectList) []Output {
	var ret []Output

	for _, item := range list.Items {
		if is(item, "output") {
			name, _ := strconv.Unquote(item.Keys[1].Token.Text)
			if name == "" {
				name = item.Keys[1].Token.Text
			}
			items := item.Val.(*ast.ObjectType).List.Items
			var desc string
			switch {
			case description(items) != "":
				desc = description(items)
			case item.LeadComment != nil:
				desc = comment(item.LeadComment.List)
			}

			ret = append(ret, Output{
				Name:        name,
				Description: strings.TrimSpace(desc),
			})
		}
	}

	return ret
}

// Get `key` from the list of object `items`.
func get(items []*ast.ObjectItem, key string) *Value {
	for _, item := range items {
		if is(item, key) {
			v := new(Value)

			if lit, ok := item.Val.(*ast.LiteralType); ok {
				if value, ok := lit.Token.Value().(string); ok {
					v.Value = value
				} else {
					v.Value = lit.Token.Text
				}
				v.Type = "string"
				return v
			} else {
				if _, ok := item.Val.(*ast.ObjectType); ok {
					v.Type = "map"
					json, err := hclAstNodeToJSON(&item.Val, v.Type)
					if err == nil {
						v.Value = json
					}
					return v
				}

				if _, ok := item.Val.(*ast.ListType); ok {
					v.Type = "list"
					json, err := hclAstNodeToJSON(&item.Val, v.Type)
					if err == nil {
						v.Value = json
					}
					return v
				}
			}

			return nil
		}
	}

	return nil
}

// description returns a description from items or an empty string.
func description(items []*ast.ObjectItem) string {
	if v := get(items, "description"); v != nil {
		return v.Value.(string)
	}

	return ""
}

// Is returns true if `item` is of `kind`.
func is(item *ast.ObjectItem, kind string) bool {
	if len(item.Keys) > 0 {
		return item.Keys[0].Token.Text == kind
	}

	return false
}

// Unquote the given string.
func unquote(s string) string {
	s, _ = strconv.Unquote(s)
	return s
}

// Comment cleans and returns a comment.
func comment(l []*ast.Comment) string {
	var line string
	var ret string

	for _, t := range l {
		line = strings.TrimSpace(t.Text)
		line = strings.TrimPrefix(line, "#")
		line = strings.TrimPrefix(line, "//")
		ret += strings.TrimSpace(line) + "\n"
	}

	return ret
}

// Header returns the header comment from the list
// or an empty comment. The head comment must start
// at line 1 and start with `/**`.
func header(c *ast.CommentGroup) (comment string) {
	if len(c.List) == 0 {
		return comment
	}

	if c.Pos().Line != 1 {
		return comment
	}

	cm := strings.TrimSpace(c.List[0].Text)

	if strings.HasPrefix(cm, "/**") {
		lines := strings.Split(cm, "\n")

		if len(lines) < 2 {
			return comment
		}

		lines = lines[1 : len(lines)-1]
		for _, l := range lines {
			l = strings.TrimSpace(l)
			switch {
			case strings.TrimPrefix(l, "* ") != l:
				l = strings.TrimPrefix(l, "* ")
			default:
				l = strings.TrimPrefix(l, "*")
			}
			comment += l + "\n"
		}
	}

	return comment
}

// hclAstNodeToJSON returns a JSON representation of an HCL AST node
func hclAstNodeToJSON(node *ast.Node, nodeType string) (interface{}, error) {
	var result interface{}

	// Marshals the HCL Ast Node Object into HCL text
	config := hclprinter.Config{
		SpacesWidth: 2,
	}

	buffer := bytes.NewBufferString("value = ")
	if err := config.Fprint(buffer, *node); err != nil {
		return nil, err
	}

	// Unmarshals HCL Text into a Golang data structure
	err := hcl.Unmarshal(buffer.Bytes(), &result)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse HCL: %s", err)
	}

	// Marshals the Golang data structure into JSON text
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

	// Unmarshal JSON text into a Golang data structure
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("Unable to unmarshal JSON: %s", err)
	}

	return result, nil
}
