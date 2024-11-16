package grpcserver

import (
	"context"
	"fmt"
	"goclub/engine/internal/grpc_server/middleware"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const pkgLogName = "grpcServer"

type GRPCAPI interface {
	RegisterGrpcServer(server *grpc.Server)
}

type GRPCServer interface {
	Start() error
	Stop() error
	RegisterAPI(api []GRPCAPI)
}

type server struct {
	ctx    context.Context
	addr   string
	server *grpc.Server
}

func NewGRPCServer(ctx context.Context, port string) *server {

	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.Panic,
			//middleware.Tracer,
			middleware.Logger,
		),
	)

	reflection.Register(srv)

	return &server{
		ctx:    ctx,
		addr:   port,
		server: srv,
	}
}

func (s *server) Start() error {
	const fnLogName = ".Start"
	listner, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf(pkgLogName+fnLogName+" - failed to create listner on address %s: %w", s.addr, err)
	}

	err = s.server.Serve(listner)
	if err != nil {
		return fmt.Errorf(pkgLogName+fnLogName+" - failed to start grpc server: %w", err)
	}

	return nil
}

func (s *server) Stop() error {
	s.server.GracefulStop()
	return nil
}

func (s *server) RegisterAPI(services []GRPCAPI) {
	for _, svc := range services {
		svc.RegisterGrpcServer(s.server)
	}
}
