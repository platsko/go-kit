// Copyright © 2020-2021 The EVEN Solutions Developers Team
// Protobuf code generation descriptors

package main

// Crypto pb files generation section
//go:generate protoc -I=. --go_out=. --go_opt=module=github.com/platsko/go-kit --go-grpc_out=. --go-grpc_opt=module=github.com/platsko/go-kit --proto_path=crypto/proto crypto/proto/*.proto

// Timestamp pb files generation section
//go:generate protoc -I=. --go_out=. --go_opt=module=github.com/platsko/go-kit --go-grpc_out=. --go-grpc_opt=module=github.com/platsko/go-kit --proto_path=timestamp/proto timestamp/proto/*.proto

