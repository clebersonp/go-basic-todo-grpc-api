generate-protoc:
	protoc \
	--go_out=model \
	--go_opt=paths=source_relative \
	--go-grpc_out=model \
	--go-grpc_opt=paths=source_relative \
	todo.proto