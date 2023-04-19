.PHONY: proto
proto:
	protoc --go_out=./product --go_opt=paths=source_relative --go-grpc_out=./product --go-grpc_opt=paths=source_relative protos/product.proto
	protoc --go_out=./summer --go_opt=paths=source_relative --go-grpc_out=./summer --go-grpc_opt=paths=source_relative protos/product.proto