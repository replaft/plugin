package shared

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Runner interface {
	Run() string
}

type RunnerRPC struct{ client *rpc.Client }

func (r *RunnerRPC) Run() string {
	var resp string

	if err := r.client.Call("Plugin.Run", new(interface{}), &resp); err != nil {
		return ""
	}

	return resp
}

type RunnerRPCServer struct {
	Impl Runner
}

func (r *RunnerRPCServer) Run(_ interface{}, resp *string) error {
	*resp = r.Impl.Run()
	return nil
}

type RunnerPlugin struct {
	Impl Runner
}

func (r *RunnerPlugin) Server(_ *plugin.MuxBroker) (interface{}, error) {
	return &RunnerRPCServer{Impl: r.Impl}, nil
}

func (RunnerPlugin) Client(_ *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RunnerRPC{client: c}, nil
}
