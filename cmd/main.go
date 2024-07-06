package main

import (
	"github.com/300degree/dumb_project/internal/adapter/grpc"
	"github.com/300degree/dumb_project/internal/adapter/http"
)

func main() {
	go http.Serve()
	grpc.Serve()
}
