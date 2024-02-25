.PHONY: generate-todo-protoc generate-user-protoc

generate-todo-protoc:
	protoc \
	--go_out=proto/todo \
	--go_opt=paths=source_relative \
	--go-grpc_out=proto/todo \
	--go-grpc_opt=paths=source_relative \
	todo.proto

generate-user-protoc:
	protoc \
	--go_out=proto/user \
	--go_opt=paths=source_relative \
	--go-grpc_out=proto/user \
	--go-grpc_opt=paths=source_relative \
	user.proto

generate-all-protoc: generate-todo-protoc generate-user-protoc