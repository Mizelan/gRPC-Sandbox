package proto

// go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
//go:generate protoc --proto_path=.. --go_out=../gen --go_opt=paths=source_relative --go-grpc_out=.. --go-grpc_opt=paths=source_relative proto/v1/entity.proto
