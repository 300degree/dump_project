package usecase

import (
	"context"

	"github.com/300degree/dumb_project/pkg/proto/pb"
)

type customerUsecaseImpl struct {
	pb.UnimplementedAuthServiceServer
}

func NewCustomerUsecase() pb.AuthServiceServer {
	return &customerUsecaseImpl{}
}

func (impl *customerUsecaseImpl) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	return nil, nil
}

func (impl *customerUsecaseImpl) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	return nil, nil
}

func (impl *customerUsecaseImpl) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	return nil, nil
}
