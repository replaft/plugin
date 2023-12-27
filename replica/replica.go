package replica

import (
	"context"
	"math"
	"net"

	"google.golang.org/grpc"

	pb "github.com/repligit/plugin/replica/proto"
)

type Server interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct {
	Addr string
}

type server struct {
	cfg *Config
	pb.UnimplementedReplicaProtoServer
}

func New(_ context.Context, cfg *Config) Server {
	return &server{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (s *server) Init(_ context.Context) error {
	return nil
}

func (s *server) Deinit(_ context.Context) error {
	return nil
}

func (s *server) Run(_ context.Context) error {
	options := []grpc.ServerOption{grpc.MaxRecvMsgSize(math.MaxInt32), grpc.MaxSendMsgSize(math.MaxInt32)}

	g := grpc.NewServer(options...)
	pb.RegisterReplicaProtoServer(g, s)

	lis, _ := net.Listen("tcp", s.cfg.Addr)

	return g.Serve(lis)
}
