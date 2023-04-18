package rpc

//go:generate protoc -I .. --go_out=paths=source_relative:. --go_opt=Mregistry.proto=github.com/fuddle-io/fuddle-rpc/go --go-grpc_out=paths=source_relative:. registry.proto registryv2.proto
