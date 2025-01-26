.PHONY: protos

protos:
	protoc -I . --go_out=. ./currency.proto
	protoc -I . --go-grpc_out=. ./currency.proto