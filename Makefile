.PHONY: pre-push go-mod check-commit
pre-push: go-mod check-commit

go-mod:
	go mod tidy -v
	go mod download -x

check-commit:
	go get -d ./...
	go test ./...
	golangci-lint run --fix ./...

.PHONY: protoc-upgrade protoc-update generate go-mod
protoc-upgrade: protoc-update generate go-mod

protoc-update:
	go get -u github.com/golang/protobuf/protoc-gen-go@latest
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

generate:
	# Crypto pb files generation section
	protoc \
		-I=. \
		--go_out=. \
		--go_opt=module=github.com/evenlab/go-kit \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/evenlab/go-kit \
		--proto_path=crypto/proto crypto/proto/*.proto

	# Timestamp pb files generation section
	protoc \
		-I=. \
		--go_out=. \
		--go_opt=module=github.com/evenlab/go-kit \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/evenlab/go-kit \
		--proto_path=timestamp/proto timestamp/proto/*.proto
