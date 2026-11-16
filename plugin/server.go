/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package plugin

import (
	goplugin "github.com/hashicorp/go-plugin"

	"github.com/rquadling/terraform-docs/print"
	"github.com/rquadling/terraform-docs/terraform"
)

// Server is an RPC Server acting as a plugin.
type Server struct {
	impl   *formatter
	broker *goplugin.MuxBroker
}

type printFunc func(*print.Config, *terraform.Module) (string, error)

// ServeOpts is an option for serving a plugin.
type ServeOpts struct {
	Name    string
	Version string
	Printer printFunc
}

// Serve is a wrapper of plugin.Serve. This is entrypoint of all plugins.
func Serve(opts *ServeOpts) {
	goplugin.Serve(&goplugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins: goplugin.PluginSet{
			"formatter": newFormatter(opts.Name, opts.Version, opts.Printer),
		},
	})
}

// Name returns the version of the plugin.
func (s *Server) Name(args interface{}, resp *string) error {
	*resp = s.impl.Name()
	return nil
}

// Version returns the version of the plugin.
func (s *Server) Version(args interface{}, resp *string) error {
	*resp = s.impl.Version()
	return nil
}

// Execute returns the generated output.
func (s *Server) Execute(args *ExecuteArgs, resp *string) error {
	r, err := s.impl.Execute(args)
	*resp = r
	return err
}
