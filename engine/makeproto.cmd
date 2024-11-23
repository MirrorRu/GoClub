rem cls
@rem go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
@rem go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
@rem go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
@rem go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

protoc ^
	--go_out=./internal/api --go_opt=paths=source_relative ^
	--go-grpc_out=./internal/api --go-grpc_opt=paths=source_relative ^
	--grpc-gateway_out=./internal/api ^
    --grpc-gateway_opt=paths=source_relative ^
    --grpc-gateway_opt=generate_unbound_methods=true ^
	--openapiv2_out=./internal/api ^
	-I proto ^
	-I vendor/proto ^
	proto/*.proto 

@echo Done
