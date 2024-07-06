MAIN_PACKAGE_DIR := ./cmd/main.go
BIN_DIR := ./bin
PROTO_DIR := ./pkg/proto
PROTO_FILES := $(PROTO_DIR)/auth.proto

GO := go
GOFLAGS := 
PROTOC := protoc
PROTOCFLAGS := --go_out=. --go-grpc_out=.

all: clean protoc build

build:
	$(GO) build -o $(BIN_DIR) ./cmd/main.go 

protoc: $(PROTO_FILES)
	$(PROTOC) $(PROTOCFLAGS) $(PROTO_FILES)

clean:
	rm -rf ./bin/*
	rm -rf pkg/proto/pb/*.go

run: build
	./bin/dump_project

