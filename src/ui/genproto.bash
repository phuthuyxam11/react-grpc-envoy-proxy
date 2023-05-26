rm -f ./src/protobuf/pb/*

protoc -I=./src/protobuf/proto ./src/protobuf/proto/*.proto \
--js_out=import_style=commonjs:./src/protobuf/pb \
--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./src/protobuf/pb
