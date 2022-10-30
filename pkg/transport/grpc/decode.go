package grpc

import (
	"context"

	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/endpoint"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/transport/grpc/pb"
)

func DecodeCreateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateUserRequest)
	return endpoint.CreateUserRequest{req.Email, req.FullName, req.Password}, nil
}

func DecodeDeleteUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteUserRequest)
	return endpoint.DeleteUserRequest{req.Email}, nil
}

func DecodeUpdateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateUserRequest)
	return endpoint.UpdateUserRequest{req.Email, req.FullName}, nil
}

func DecodeGetUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetUserRequest)
	return endpoint.GetUserRequest{req.Email}, nil
}

func DecodeLoginRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.LoginUserRequest)
	return endpoint.LoginUserRequest{req.Email, req.Password}, nil
}
