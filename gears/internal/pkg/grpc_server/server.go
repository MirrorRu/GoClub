package grpcserver

import (
	"context"
	"fmt"
	"net"

	"goclub/gears/internal/pkg/middleware"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const pkgLogName = "grpcServer"

type API interface {
	RegisterGrpcServer(server *grpc.Server)
}

type Server struct {
	ctx        context.Context
	addr       string
	grpcServer *grpc.Server
}

func NewServer(ctx context.Context, port string) *Server {

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			middleware.Panic,
			middleware.Tracer,
			// middleware.Logger,
			// middleware.Validate,
			// middleware.Metrics,
		),
	)

	reflection.Register(grpcServer)

	return &Server{
		ctx:        ctx,
		addr:       port,
		grpcServer: grpcServer,
	}
}

func (s *Server) Start() error {
	const fnLogName = ".Start"
	listner, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf(pkgLogName+fnLogName+" - failed to create listner on address %s: %w", s.addr, err)
	}

	err = s.grpcServer.Serve(listner)
	if err != nil {
		return fmt.Errorf(pkgLogName+fnLogName+" - failed to start grpc server: %w", err)
	}

	return nil
}

func (s *Server) Stop() error {
	s.grpcServer.GracefulStop()
	return nil
}

func (s *Server) RegisterAPI(APIs []API) {
	for _, api := range APIs {
		api.RegisterGrpcServer(s.grpcServer)
	}
}
