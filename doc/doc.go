package doc

import (
	"path"
	"strconv"
	"strings"

	"github.com/hashicorp/hcl/hcl/ast"
)

// Input represents a terraform input variable.
type Input struct {
	Name        string
	Description string
	Default     *Value
}

// Value returns the default value as a string.
func (i *Input) Value() string {
	if i.Default != nil {
		switch i.Default.Type {
		case "string":
			return i.Default.Literal
		case "map":
			return "<map>"
		}
	}

	return "required"
}

// Value represents a terraform value.
type Value struct {
	Type    string
	Literal string
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

	return doc
}

// Inputs returns all variables from `list`.
func inputs(list *ast.ObjectList) []Input {
	var ret []Input

	for _, item := range list.Items {
		if is(item, "variable") {
			name, _ := strconv.Unquote(item.Keys[1].Token.Text)
			items := item.Val.(*ast.ObjectType).List.Items
			desc, _ := strconv.Unquote(description(items))
			def := get(items, "default")
			ret = append(ret, Input{
				Name:        name,
				Description: desc,
				Default:     def,
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

			var desc string
			if c := item.LeadComment; c != nil {
				desc = comment(c.List)
			}

			ret = append(ret, Output{
				Name:        name,
				Description: desc,
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
				v.Literal = lit.Token.Text
				v.Type = "string"
				return v
			}

			if _, ok := item.Val.(*ast.ObjectType); ok {
				v.Type = "map"
				return v
			}

			return nil
		}
	}

	return nil
}

// description returns a description from items or an empty string.
func description(items []*ast.ObjectItem) string {
	if v := get(items, "description"); v != nil {
		return v.Literal
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
			l = strings.TrimPrefix(l, "*")
			comment += l + "\n"
		}
	}

	return comment
}
