package terraform

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
)

// parseModuleFiles reads all `.tf`/`.tofu` from dir and returns them keyed by base filename, ready for earlydecoder.
func parseModuleFiles(dir string) (map[string]*hcl.File, hcl.Diagnostics) {
	parser := hclparse.NewParser()
	files := map[string]*hcl.File{}
	var diags hcl.Diagnostics

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, hcl.Diagnostics{{
			Severity: hcl.DiagError,
			Summary:  "Cannot read module directory.",
			Detail:   err.Error(),
		}}
	}

	for _, er := range entries {
		if er.IsDir() {
			continue
		}
		name := er.Name()
		path := filepath.Join(dir, name)

		switch {
		case strings.HasSuffix(name, ".tf"), strings.HasSuffix(name, ".tofu"):
			f, d := parser.ParseHCLFile(path)
			diags = append(diags, d...)
			if f != nil {
				files[name] = f
			}
		case strings.HasSuffix(name, ".tf.json"), strings.HasSuffix(name, ".tofu.json"):
			f, d := parser.ParseJSONFile(path)
			diags = append(diags, d...)
			if f != nil {
				files[name] = f
			}
		}
	}
	return files, diags
}
