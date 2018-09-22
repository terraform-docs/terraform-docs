package doc

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/segmentio/terraform-docs/internal/pkg/fs"
)

// Doc represents a Terraform module.
type Doc struct {
	Comment string
	Inputs  []Input
	Outputs []Output
}

// Value represents a Terraform value.
type Value struct {
	Type    string
	Literal string
}

// HasComment indicates if the document has a comment.
func (d *Doc) HasComment() bool {
	return len(d.Comment) > 0
}

// HasInputs indicates if the document has inputs.
func (d *Doc) HasInputs() bool {
	return len(d.Inputs) > 0
}

// HasOutputs indicates if the document has outputs.
func (d *Doc) HasOutputs() bool {
	return len(d.Outputs) > 0
}

type inputsByName []Input

func (a inputsByName) Len() int           { return len(a) }
func (a inputsByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a inputsByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type outputsByName []Output

func (a outputsByName) Len() int           { return len(a) }
func (a outputsByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a outputsByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

// CreateFromPaths creates a new document from a list of file or directory paths.
func CreateFromPaths(paths []string) (*Doc, error) {
	names := make([]string, 0)

	for _, path := range paths {
		if fs.DirectoryExists(path) {
			matches, err := filepath.Glob(fmt.Sprintf("%s/*.tf", path))
			if err != nil {
				log.Fatal(err)
			}

			names = append(names, matches...)
		} else if fs.FileExists(path) {
			names = append(names, path)
		}
	}

	files := make(map[string]*ast.File)

	for _, name := range names {
		bytes, err := ioutil.ReadFile(name)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		ast, err := hcl.ParseBytes(bytes)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		files[name] = ast
	}

	return Create(files), nil
}

// Create creates a new document from a map of filenames to *ast.Files.
func Create(files map[string]*ast.File) *Doc {
	doc := new(Doc)

	for name, file := range files {
		objects := file.Node.(*ast.ObjectList)

		doc.Inputs = append(doc.Inputs, getInputs(objects)...)
		doc.Outputs = append(doc.Outputs, getOutputs(objects)...)

		filename := path.Base(name)
		comments := file.Comments
		if filename == "main.tf" && len(comments) > 0 {
			doc.Comment = header(comments[0])
		}
	}

	sort.Sort(inputsByName(doc.Inputs))
	sort.Sort(outputsByName(doc.Outputs))
	return doc
}

// getInputs returns a list of inputs from an ast.ObjectList.
func getInputs(list *ast.ObjectList) []Input {
	var result []Input

	for _, item := range list.Items {
		if isItemOfKindVariable(item) {
			result = append(result, Input{
				Name:        getItemName(item),
				Description: getItemDescription(item),
				Default:     getItemDefault(item),
				Type:        getItemType(item),
			})
		}
	}

	return result
}

// getOutputs returns a list of outputs from an ast.ObjectList.
func getOutputs(list *ast.ObjectList) []Output {
	var result []Output

	for _, item := range list.Items {
		if isItemOfKindOutput(item) {
			result = append(result, Output{
				Name:        getItemName(item),
				Description: getItemDescription(item),
			})
		}
	}

	return result
}

func getItemByKey(items []*ast.ObjectItem, key string) *Value {
	for _, item := range items {
		if isItemOfKind(item, key) {
			result := new(Value)

			if literal, ok := item.Val.(*ast.LiteralType); ok {
				result.Type = "string"

				if value, ok := literal.Token.Value().(string); ok {
					result.Literal = value
				} else {
					result.Literal = literal.Token.Text
				}

				return result
			}

			if _, ok := item.Val.(*ast.ObjectType); ok {
				result.Type = "map"
				return result
			}

			if _, ok := item.Val.(*ast.ListType); ok {
				result.Type = "list"
				return result
			}

			return nil
		}
	}

	return nil
}

func getItemDefault(item *ast.ObjectItem) *Value {
	items := item.Val.(*ast.ObjectType).List.Items
	return getItemByKey(items, "default")
}

func getItemDescription(item *ast.ObjectItem) string {
	var result string

	items := item.Val.(*ast.ObjectType).List.Items

	var description = getItemByKey(items, "description")
	if description != nil {
		result = description.Literal
	}

	if result == "" {
		result = getItemDescriptionFromComment(item.LeadComment.List)
	}

	return result
}

func getItemDescriptionFromComment(comments []*ast.Comment) string {
	var result string

	for _, comment := range comments {
		line := strings.TrimSpace(comment.Text)
		line = strings.TrimPrefix(line, "#")
		line = strings.TrimPrefix(line, "//")
		result += strings.TrimSpace(line)
	}

	return result
}

func getItemName(item *ast.ObjectItem) string {
	name, err := strconv.Unquote(item.Keys[1].Token.Text)
	if err != nil {
		log.Fatal(err)
	}

	if name == "" {
		name = item.Keys[1].Token.Text
	}

	return name
}

func getItemType(item *ast.ObjectItem) string {
	var result string

	items := item.Val.(*ast.ObjectType).List.Items

	var itemsType = getItemByKey(items, "type")
	if itemsType == nil || itemsType.Literal == "" {
		result = "string"
	} else {
		result = itemsType.Literal
	}

	return result
}

func isItemOfKind(item *ast.ObjectItem, kind string) bool {
	if len(item.Keys) > 0 {
		return item.Keys[0].Token.Text == kind
	}

	return false
}

func isItemOfKindOutput(item *ast.ObjectItem) bool {
	return isItemOfKind(item, "output")
}

func isItemOfKindVariable(item *ast.ObjectItem) bool {
	return isItemOfKind(item, "variable")
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
