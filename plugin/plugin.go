package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"

	"github.com/repligit/plugin/plugin/shared"
)

type RunnerImpl struct {
	logger hclog.Logger
}

func (r *RunnerImpl) Run() string {
	r.logger.Debug("message from RunnerImpl.Run")
	return "Hello!"
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "plugin",
	MagicCookieValue: "plugin",
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	runner := &RunnerImpl{
		logger: logger,
	}
	var pluginMap = map[string]plugin.Plugin{
		"plugin": &shared.RunnerPlugin{Impl: runner},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
