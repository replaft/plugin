package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"

	"github.com/repligit/plugin/plugin/shared"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("./plugin"),
		Logger:          logger,
	})

	rpcClient, err := client.Client()
	if err != nil {
		client.Kill()
		log.Fatal(err)
	}

	raw, err := rpcClient.Dispense("plugin")
	if err != nil {
		client.Kill()
		log.Fatal(err)
	}

	runner := raw.(shared.Runner)
	fmt.Println(runner.Run())

	client.Kill()
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "plugin",
	MagicCookieValue: "plugin",
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"plugin": &shared.RunnerPlugin{},
}
