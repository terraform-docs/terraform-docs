package plugin

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	goplugin "github.com/hashicorp/go-plugin"
	"github.com/mitchellh/go-homedir"

	pluginsdk "github.com/terraform-docs/plugin-sdk/plugin"
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

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		name := strings.Replace(f.Name(), namePrefix, "", -1)
		path, err := getPluginPath(dir, name)
		if err != nil {
			return nil, err
		}

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
