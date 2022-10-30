package grpc

import (
	"context"

	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/endpoint"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/transport/grpc/pb"
)

func EncodeCreateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.CreateUserResponse)
	return &pb.CreateUserResponse{Id: int64(resp.ID), Email: resp.Email, FullName: resp.FullName, Err: resp.Err}, nil
}

func EncodeUpdateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.UpdateUserResponse)
	return &pb.UpdateUserResponse{Id: int64(resp.ID), Email: resp.Email, FullName: resp.FullName, Err: resp.Err}, nil
}

func EncodeDeleteUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.DeleteUserResponse)
	return &pb.DeleteUserResponse{Ok: resp.Ok, Err: resp.Err}, nil
}

func EncodeLoginResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.LoginUserResponse)
	return &pb.LoginUserResponse{Token: resp.Token, Err: resp.Err}, nil
}

func EncodeGetResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetUserResponse)
	return &pb.GetUserResponse{Id: int64(resp.ID), Email: resp.Email, FullName: resp.FullName, Err: resp.Err}, nil
}
