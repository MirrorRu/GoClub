@mkdir vendor\proto\google
@goto next
rmdir /S /Q temp\protobuf.tmp
git clone --depth 1 --no-checkout https://github.com/protocolbuffers/internal/protobuf temp/internal/protobuf.tmp
cd temp/internal/protobuf.tmp
git checkout HEAD -- src*.proto
cd ../..
rmdir /S /Q vendor\proto\google\protobuf
move temp\protobuf.tmp\src\google\protobuf vendor\proto\google

@REM https://github.com/googleapis/googleapis
rmdir /S /Q temp\googleapis.tmp
git clone --depth 1 --no-checkout https://github.com/googleapis/googleapis temp/googleapis.tmp
cd temp/googleapis.tmp
git checkout HEAD -- google/api*.proto
cd ../..
rmdir /S /Q vendor\proto\goolge\api
move temp\googleapis.tmp\google\api vendor\proto\google

@REM https://github.com/grpc-ecosystem/grpc-gateway
rmdir /S /Q temp\grpc-gateway.tmp
git clone --depth 1 --no-checkout https://github.com/grpc-ecosystem/grpc-gateway temp/grpc-gateway.tmp
cd temp/grpc-gateway.tmp
git checkout HEAD -- protoc-gen-openapiv2/*.proto
cd ../..
rmdir /S /Q vendor\proto\protoc-gen-openapiv2
move temp\grpc-gateway.tmp\protoc-gen-openapiv2 vendor\proto

:next
@REM https://github.com/bufbuild/protoc-gen-validate
rmdir /S /Q temp\protoc-gen-validate.tmp
git clone --depth 1 --no-checkout https://github.com/bufbuild/protoc-gen-validate temp/protoc-gen-validate.tmp
cd temp/protoc-gen-validate.tmp
git checkout HEAD -- validate/*.proto
cd ../..
rmdir /S /Q vendor\proto\protoc-gen-validate
move temp\protoc-gen-validate.tmp\validate vendor\proto

