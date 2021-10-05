/*
Copyright 2021 The terraform-docs Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package plugin

import (
	goplugin "github.com/hashicorp/go-plugin"

	"github.com/terraform-docs/terraform-docs/print"
	"github.com/terraform-docs/terraform-docs/terraform"
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
