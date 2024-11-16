package httpserver

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"goclub/common/logger"
	"io"
	"net/http"
	"net/http/pprof"
	"os"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const pkgLogName = "httpServer"

type HTTPAPI interface {
	RegisterHttpServer(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
}

type HTTPServer interface {
	Start() error
	Stop() error
	RegisterAPI(api []HTTPAPI) error
}

type server struct {
	ctx        context.Context
	mux        *http.ServeMux
	gwmux      *runtime.ServeMux
	conn       *grpc.ClientConn
	httpServer *http.Server
}

func NewServer(ctx context.Context, httpAddr, grpcAddr string) (HTTPServer, error) {
	const fnLogName = ".NewServer"
	var err error

	s := &server{
		ctx:   ctx,
		mux:   http.NewServeMux(),
		gwmux: runtime.NewServeMux(
		// TODO select headers to transfer
		// runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
		// 	md := make(metadata.MD, len(request.Header))
		// 	metadata.Pairs()
		// 	for k, v := range request.Header {
		// 		md[strings.ToLower(k)] = v
		// 	}
		// 	fmt.Println(md)
		// 	return md
		// }),
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
		Handler:           s.mux,
		ReadHeaderTimeout: 1 * time.Second,
		ReadTimeout:       3 * time.Second,
	}

	s.mux.HandleFunc("GET /swagger.json", s.handleSwagger)

	fs := http.FileServer(http.Dir("pkg/swagger-ui"))
	s.mux.HandleFunc("GET /", func(response http.ResponseWriter, request *http.Request) {
		// Приветствие первой страницы
		logger.Debug(ctx, "Hello()", "PATH", request.URL.Path)
		response.Write([]byte("Привет! " + request.URL.Path))
	})

	s.mux.Handle("GET /swagger-ui/", http.StripPrefix("/swagger-ui", fs))

	s.mux.Handle("GET /metrics", promhttp.Handler())

	s.mux.Handle("GET /api/", s.gwmux)
	s.mux.Handle("POST /api/", s.gwmux)
	s.mux.Handle("GET /svc", s.gwmux)
	s.mux.Handle("POST /svc", s.gwmux)
	s.mux.HandleFunc("GET /debug/pprof/", pprof.Index)
	s.mux.HandleFunc("GET /debug/pprof/cmdline", pprof.Cmdline)
	s.mux.HandleFunc("GET /debug/pprof/profile", pprof.Profile)
	s.mux.HandleFunc("GET /debug/pprof/symbol", pprof.Symbol)
	s.mux.HandleFunc("GET /debug/pprof/trace", pprof.Trace)

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

func (s *server) RegisterAPI(api []HTTPAPI) error {
	for _, singleAPI := range api {
		err := singleAPI.RegisterHttpServer(s.ctx, s.gwmux, s.conn)
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
