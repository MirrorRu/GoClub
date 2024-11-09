package grpcserver

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const pkgLogName = "grpcServer"

// type AppService interface {
// 	RegisterGrpcServer(server *grpc.Server)
// 	RegisterHttpServer(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
// }

type GRPCService interface {
	RegisterGrpcServer(server *grpc.Server)
}

type grpcServer struct {
	ctx    context.Context
	addr   string
	server *grpc.Server
}

func NewGRPCServer(ctx context.Context, port string) *grpcServer {

	srv := grpc.NewServer()

	reflection.Register(srv)

	return &grpcServer{
		ctx:    ctx,
		addr:   port,
		server: srv,
	}
}

func (s *grpcServer) Start() error {
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

func (s *grpcServer) Stop() error {
	s.server.GracefulStop()
	return nil
}

func (s *grpcServer) RegisterAPI(services []GRPCService) {
	for _, svc := range services {
		svc.RegisterGrpcServer(s.server)
	}
}
