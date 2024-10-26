gen:
	@protoc \
	 --proto_path=protobuf "protobuf/transfer.proto" \
	 --go_out=service/genproto  --go_opt=paths=source_relative \
	 --go-grpc_out=service/genproto  --go-grpc_opt=paths=source_relative