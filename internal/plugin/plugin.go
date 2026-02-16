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

	pluginsdk "github.com/rquadling/terraform-docs/plugin"
)

// namePrefix is the mandatory prefix for name of the plugin file. What
// comes after this is considered to be identifier of the plugin and in
// the overall ecosystem should be unique (as much as possible.)
const namePrefix = "tfdocs-format-"

// homePluginsRoot is the root directory of the plugins
var homePluginsRoot = "~/.tfdocs.d/plugins"
var localPluginsRoot = "./.tfdocs.d/plugins"

// List is an object caching discovered plugins and their corresponding
// clients. Basically, it is a wrapper for go-plugin and provides an API
// to handle them collectively.
type List struct {
	formatters map[string]*pluginsdk.Client
	clients    map[string]*goplugin.Client
}

// All returns all registered plugins.
func (l *List) All() []*pluginsdk.Client {
	all := make([]*pluginsdk.Client, 0)
	for _, f := range l.formatters {
		all = append(all, f)
	}
	return all
}

// Get plugin by its name.
func (l *List) Get(name string) (*pluginsdk.Client, bool) {
	client, ok := l.formatters[name]
	return client, ok
}

// Clean is a helper for ending plugin processes.
func (l *List) Clean() {
	for _, client := range l.clients {
		client.Kill()
	}
}
