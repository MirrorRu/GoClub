package middleware

import (
	"context"
	"fmt"
	"goclub/common/logger"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

const pkgLogName = "serverMW"

func Logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	raw, _ := protojson.Marshal((req).(proto.Message)) // для превращения protbuf структур в json используем google.golang.org/protobuf/encoding/protojson пакет а не encoding/json
	logger.Info(ctx, "request", "method", info.FullMethod, "req", string(raw))

	if resp, err = handler(ctx, req); err != nil {
		logger.Info(ctx, "handle error", "method", info.FullMethod, "err", err)
		return
	}

	rawResp, _ := protojson.Marshal((resp).(proto.Message))
	logger.Debug(ctx, "response", "method", info.FullMethod, "rawResp", string(rawResp))

	return
}

func WithHTTPLoggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func Metrics(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	// metrics.UpdateRequestsTotal(info.FullMethod)

	// defer metrics.UpdateResponseTime(time.Now().UTC())

	// resp, err = handler(ctx, req)
	// if err != nil {
	// 	st, _ := status.FromError(err)
	// 	metrics.UpdateResponseCode(info.FullMethod, st.Code().String())
	// } else {
	// 	metrics.UpdateResponseCode(info.FullMethod, codes.OK.String())
	// }

	return
}

func Panic(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	const fnLogName = ".Panic"
	defer func() {
		if e := recover(); e != nil {
			logger.Error(ctx, fmt.Sprintf(pkgLogName+fnLogName+" - panic: %v", e), nil)
			err = status.Errorf(codes.Internal, pkgLogName+fnLogName+" - panic: %v", e)
		}
	}()
	resp, err = handler(ctx, req)
	return resp, err
}

func Tracer(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// const fnLogName = ".Tracer"
	// md, _ := metadata.FromIncomingContext(ctx)

	// var traceID string
	// if traceIDs, exists := md["x-trace-id"]; exists {
	// 	traceID = traceIDs[0]
	// }

	// var spanID string
	// if spanIDs, exists := md["x-span-id"]; exists {
	// 	spanID = spanIDs[0]
	// }

	// if traceID != "" {
	// 	var err error
	// 	spanContext := trace.SpanContextConfig{
	// 		TraceFlags: trace.FlagsSampled,
	// 		Remote:     true,
	// 	}
	// 	spanContext.TraceID, err = trace.TraceIDFromHex(traceID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	spanContext.SpanID, err = trace.SpanIDFromHex(spanID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	ctx = trace.ContextWithSpanContext(ctx,
	// 		trace.NewSpanContext(spanContext))
	// }

	// var span trace.Span
	// ctx, span = tracker.StartSpan(ctx, pkgLogName+fnLogName+"_"+info.FullMethod, trace.WithSpanKind(trace.SpanKindServer))
	// defer span.End()

	return //handler(ctx, req)
}

// func Validate(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
// 	if v, ok := req.(interface{ Validate() error }); ok {
// 		if err := v.Validate(); err != nil {
// 			return nil, status.Error(codes.InvalidArgument, err.Error())
// 		}
// 	}
// 	return handler(ctx, req)
// }
