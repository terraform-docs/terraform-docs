/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package plugin

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	goplugin "github.com/hashicorp/go-plugin"
	"github.com/mitchellh/go-homedir"

	pluginsdk "github.com/terraform-docs/terraform-docs/plugin"
)

// Discover plugins and registers them. The lookup priority of plugins is as
// follow:
//
// 1. `TFDOCS_PLUGIN_DIR` environment variable (if it's set)
// 2. Current directory (./.tfdocs.d/plugins)
// 3. Home directory (~/.tfdocs.d/plugins)
//
// Files under these directories that satisfy the "tfdocs-format-*" naming
// convention are treated as plugins.
func Discover() (*List, error) {
	if dir := os.Getenv("TFDOCS_PLUGIN_DIR"); dir != "" {
		return findPlugins(dir)
	}

	if _, err := os.Stat(localPluginsRoot); !os.IsNotExist(err) {
		return findPlugins(localPluginsRoot)
	}

	dir, err := homedir.Expand(homePluginsRoot)
	if err != nil {
		return nil, err
	}

	return findPlugins(dir)
}

// findPlugins finds plugins in a given 'dir' and registers them.
func findPlugins(dir string) (*List, error) {
	clients := map[string]*goplugin.Client{}
	formatters := map[string]*pluginsdk.Client{}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		name := strings.ReplaceAll(f.Name(), namePrefix, "")
		path, err := getPluginPath(dir, name)
		if err != nil {
			return nil, err
		}

		// Accepting variables here is intentional; we need to determine the
		// path on the fly per directory.
		//
		// nolint:gosec
		cmd := exec.Command(path)

		client := pluginsdk.NewClient(&pluginsdk.ClientOpts{
			Cmd: cmd,
		})

		rpcClient, err := client.Client()
		if err != nil {
			return nil, err
		}

		raw, err := rpcClient.Dispense("formatter")
		if err != nil {
			return nil, err
		}

		formatter := raw.(*pluginsdk.Client)

		if _, ok := clients[name]; ok {
			return nil, fmt.Errorf("plugin %s is already registered", name)
		}

		clients[name] = client
		formatters[name] = formatter
	}

	return &List{formatters: formatters, clients: clients}, nil
}

func getPluginPath(dir string, name string) (string, error) {
	suffix := ""

	if runtime.GOOS == "windows" {
		suffix += ".exe"
	}

	path := filepath.Join(dir, fmt.Sprintf("%s%s%s", namePrefix, name, suffix))

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", os.ErrNotExist
	}

	return path, nil
}
