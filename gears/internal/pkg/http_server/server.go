package httpserver

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
	"time"

	"goclub/gears/internal/logger"
	"goclub/gears/internal/pkg/middleware"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const pkgLogName = "httpServer"

type API interface {
	RegisterHttpHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
}

type Server interface {
	Start() error
	Stop() error
	RegisterAPI(api []API) error
}

type server struct {
	ctx        context.Context
	mux        *http.ServeMux
	gwmux      *runtime.ServeMux
	conn       *grpc.ClientConn
	httpServer *http.Server
}

func NewServer(ctx context.Context, httpAddr, grpcAddr string) (Server, error) {
	const fnLogName = ".NewServer"
	var err error

	s := &server{
		ctx: ctx,
		mux: http.NewServeMux(),
		gwmux: runtime.NewServeMux(
			runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
				md := make(metadata.MD, len(request.Header))
				metadata.Pairs()
				for k, v := range request.Header {
					md[strings.ToLower(k)] = v
				}
				return md
			}),
		),
	}

	s.conn, err = grpc.NewClient(
		grpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, fmt.Errorf(pkgLogName+fnLogName+" - failed to deal: %w", err)
	}

	s.httpServer = &http.Server{
		Addr:              httpAddr,
		Handler:           middleware.WithHTTPLoggingMiddleware(ctx, s.mux),
		ReadHeaderTimeout: 1 * time.Second,
		ReadTimeout:       3 * time.Second,
	}

	s.mux.HandleFunc("/swagger.json", s.handleSwagger)

	fs := http.FileServer(http.Dir("pkg/swagger-ui"))
	s.mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", fs))

	s.mux.Handle("/metrics", promhttp.Handler())

	s.mux.Handle("/", s.gwmux)
	s.mux.HandleFunc("/debug/pprof/", pprof.Index)
	s.mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	s.mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	s.mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	s.mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	return s, nil
}

func (s *server) Start() error {
	const fnLogName = ".Start"
	err := s.httpServer.ListenAndServe()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf(pkgLogName+fnLogName+" - failed to start http server: %w", err)
		}
	}
	return nil
}

func (s *server) Stop() error {
	ctx, cancel := context.WithTimeout(s.ctx, 3*time.Second)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}

func (s *server) RegisterAPI(api []API) error {
	for _, singleAPI := range api {
		err := singleAPI.RegisterHttpHandler(s.ctx, s.gwmux, s.conn)
		if err != nil {
			return nil
		}
	}
	return nil
}

func (s *server) handleSwagger(w http.ResponseWriter, req *http.Request) {
	const fnLogName = ".handleSwagger"
	file, err := os.Open("pkg/swagger/swagger.json")
	if err != nil {
		http.Error(w, pkgLogName+fnLogName+" - swagger not found", http.StatusNotFound)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	if _, err := io.Copy(w, reader); err != nil {
		logger.Error(req.Context(), pkgLogName+fnLogName, err)
	}
}
