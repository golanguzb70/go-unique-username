proto-gen:
	protoc -I=./protos --go_out=./ --go-grpc_out=./ ./protos/server.proto
