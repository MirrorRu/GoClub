rem cls
protoc ^
	--go_out=./internal/api --go_opt=paths=source_relative ^
	--go-grpc_out=./internal/api --go-grpc_opt=paths=source_relative ^
	--grpc-gateway_out=./internal/api ^
    --grpc-gateway_opt=paths=source_relative ^
    --grpc-gateway_opt=generate_unbound_methods=true ^
	--openapiv2_out=./internal/api ^
	-I proto ^
	-I vendor/proto ^
	proto/goclub.proto 

@echo Done
