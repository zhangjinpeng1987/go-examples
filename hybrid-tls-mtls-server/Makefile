.PHONY: build

build:
	@GO111MODULE=on go build -o bin/main ./main.go

proto_go:
	@echo "Generating proto code for golang"
	@protoc \
		--go_out=./pkg/proto \
		--go_opt paths=source_relative \
		--plugin protoc-gen-go="${GOPATH}/bin/protoc-gen-go" \
		--go-grpc_out=./pkg/proto \
		--go-grpc_opt paths=source_relative \
		--plugin protoc-gen-go-grpc="${GOPATH}/bin/protoc-gen-go-grpc" \
		proto/*.proto
	@echo "Done"
