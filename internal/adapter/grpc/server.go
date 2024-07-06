package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/300degree/dumb_project/internal/usecase"
	"github.com/300degree/dumb_project/pkg/proto/pb"
	"google.golang.org/grpc"
)

func Serve() {
	grpcServer := grpc.NewServer()
	service := usecase.NewCustomerUsecase()
	port := 50051

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterAuthServiceServer(grpcServer, service)

	fmt.Printf("gRPC server listening on port %v", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
