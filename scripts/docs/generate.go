/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/rquadling/terraform-docs/cmd"
	"github.com/rquadling/terraform-docs/format"
	"github.com/rquadling/terraform-docs/print"
	"github.com/rquadling/terraform-docs/terraform"
)

// These are practically a copy/paste of https://github.com/spf13/cobra/blob/master/doc/md_docs.go
// The reason we've decided to bring them over and not use them directly
// from cobra module was that we wanted to inject custom "Example" section
// with generated output based on the "examples" folder.

var (
	baseWeight = 950
)

func main() {
	if err := generate(cmd.NewCommand(), baseWeight, "terraform-docs"); err != nil {
		log.Fatal(err)
	}
}

func ignore(cmd *cobra.Command) bool {
	switch {
	case !cmd.IsAvailableCommand():
		return true
	case cmd.IsAdditionalHelpTopicCommand():
		return true
	case cmd.Annotations["kind"] == "":
		return true
	case cmd.Annotations["kind"] != "formatter":
		return true
	}
	return false
}

func generate(cmd *cobra.Command, weight int, basename string) error {
	for _, c := range cmd.Commands() {
		if ignore(c) {
			continue
		}
		b := extractFilename(c.CommandPath())
		baseWeight++
		if err := generate(c, baseWeight, b); err != nil {
			return err
		}
	}

	filename := filepath.Join("docs", "reference", basename+".md")
	f, err := os.Create(filepath.Clean(filename))
	if err != nil {
		return err
	}
	defer f.Close() //nolint:errcheck,gosec

	if _, err := f.WriteString(""); err != nil {
		return err
	}
	if err := generateMarkdown(cmd, weight, f); err != nil {
		return err
	}
	return nil
}

type reference struct {
	Name             string
	Command          string
	Description      string
	Parent           string
	Synopsis         string
	Runnable         bool
	HasChildren      bool
	UseLine          string
	Options          string
	InheritedOptions string
	Usage            string
	Example          string
	Subcommands      []command
	Weight           int
}

type command struct {
	Name     string
	Link     string
	Children []command
}

func generateMarkdown(cmd *cobra.Command, weight int, w io.Writer) error {
	cmd.InitDefaultHelpCmd()
	cmd.InitDefaultHelpFlag()

	command := cmd.CommandPath()
	name := strings.ReplaceAll(command, "terraform-docs ", "")

	short := cmd.Short
	long := cmd.Long

	if len(long) == 0 {
		long = short
	}

	parent := "reference"
	if cmd.Parent() != nil {
		parent = cmd.Parent().Name()
	}

	ref := &reference{
		Name:        name,
		Command:     command,
		Description: short,
		Parent:      parent,
		Synopsis:    long,
		Runnable:    cmd.Runnable(),
		HasChildren: len(cmd.Commands()) > 0,
		UseLine:     cmd.UseLine(),
		Weight:      weight,
	}

	// Options
	if f := cmd.NonInheritedFlags(); f.HasAvailableFlags() {
		ref.Options = f.FlagUsages()
	}

	// Inherited Options
	if f := cmd.InheritedFlags(); f.HasAvailableFlags() {
		ref.InheritedOptions = f.FlagUsages()
	}

	if ref.HasChildren {
		subcommands(ref, cmd.Commands())
	} else {
		example(ref) //nolint:errcheck,gosec
	}

	file := "format.tmpl"
	paths := []string{filepath.Join("scripts", "docs", file)}

	t := template.Must(template.New(file).ParseFiles(paths...))

	return t.Execute(w, ref)
}

func example(ref *reference) error {
	flag := " --footer-from footer.md"
	if ref.Name == "pretty" {
		flag += " --no-color"
	}

	ref.Usage = fmt.Sprintf("%s%s ./examples/", ref.Command, flag)

	config := print.DefaultConfig()
	config.ModuleRoot = "./examples"
	config.Formatter = ref.Name
	config.Settings.Color = false
	config.Sections.Show = append(config.Sections.Show, "all")
	config.Sections.Footer = true
	config.FooterFrom = "footer.md"
	config.Parse()

	tfmodule, err := terraform.LoadWithOptions(config)
	if err != nil {
		log.Fatal(err)
	}

	formatter, err := format.New(config)
	if err != nil {
		return err
	}

	if err := formatter.Generate(tfmodule); err != nil {
		return err
	}

	segments := strings.Split(formatter.Content(), "\n")
	buf := new(bytes.Buffer)
	for _, s := range segments {
		if s == "" {
			buf.WriteString("\n")
		} else {
			fmt.Fprintf(buf, "    %s\n", s)
		}
	}
	ref.Example = buf.String()

	return nil
}

func subcommands(ref *reference, children []*cobra.Command) {
	subs := []command{}
	for _, child := range children {
		if ignore(child) {
			continue
		}
		subchild := []command{}
		for _, c := range child.Commands() {
			if ignore(c) {
				continue
			}
			cname := c.CommandPath()
			link := extractFilename(cname)
			subchild = append(subchild, command{Name: cname, Link: link})
		}
		cname := child.CommandPath()
		link := extractFilename(cname)
		subs = append(subs, command{Name: cname, Link: link, Children: subchild})
	}
	ref.Subcommands = subs
}

func extractFilename(s string) string {
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, "terraform-docs-", "")
	return s
}
