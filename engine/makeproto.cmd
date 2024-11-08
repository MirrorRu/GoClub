rem cls
protoc ^
	--go_out=./grpc_api --go_opt=paths=source_relative ^
	--go-grpc_out=./grpc_api --go-grpc_opt=paths=source_relative ^
	--grpc-gateway_out=./grpc_api ^
    --grpc-gateway_opt=paths=source_relative ^
    --grpc-gateway_opt=generate_unbound_methods=true ^
	--openapiv2_out=./grpc_api ^
	-I proto ^
	-I vendor/proto ^
	proto/goclub.proto 

@echo Done
