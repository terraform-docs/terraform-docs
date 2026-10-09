/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package plugin

import (
	"encoding/gob"
	"net/rpc"

	goplugin "github.com/hashicorp/go-plugin"

	"github.com/rquadling/terraform-docs/internal/types"
)

// Ensure formatter fully satisfy plugin interface.
var _ goplugin.Plugin = &formatter{}

// handshakeConfig is used for UX. ProcotolVersion will be updated by incompatible changes.
var handshakeConfig = goplugin.HandshakeConfig{
	ProtocolVersion:  7,
	MagicCookieKey:   "TFDOCS_PLUGIN",
	MagicCookieValue: "A7U5oTDDJwdL6UKOw6RXATDa86NEo4xLK3rz7QqegT1N4EY66qb6UeAJDSxLwtXH",
}

// formatter is a wrapper to satisfy the interface of go-plugin.
type formatter struct {
	name    string
	version string
	printer printFunc
}

func newFormatter(name string, version string, printer printFunc) *formatter {
	return &formatter{
		name:    name,
		version: version,
		printer: printer,
	}
}

func (f *formatter) Name() string {
	return f.name
}

func (f *formatter) Version() string {
	return f.version
}

func (f *formatter) Execute(args *ExecuteArgs) (string, error) {
	return f.printer(args.Config, args.Module)
}

// Server returns an RPC server acting as a plugin.
func (f *formatter) Server(b *goplugin.MuxBroker) (interface{}, error) {
	return &Server{impl: f, broker: b}, nil
}

// Client returns an RPC client for the host.
func (*formatter) Client(b *goplugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &Client{rpcClient: c, broker: b}, nil
}

func init() {
	gob.Register(new(types.Bool))
	gob.Register(new(types.Empty))
	gob.Register(new(types.List))
	gob.Register(new(types.Map))
	gob.Register(new(types.Nil))
	gob.Register(new(types.Number))
	gob.Register(new(types.String))
}
