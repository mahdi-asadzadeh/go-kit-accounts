package grpc

import (
	"context"

	kitlog "github.com/go-kit/kit/log"
	gr "github.com/go-kit/kit/transport/grpc"

	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/endpoint"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/transport/grpc/pb"
)

type gRPCServer struct {
	pb.UnimplementedUSerServiceServer
	createUser gr.Handler
	updateUser gr.Handler
	deleteUser gr.Handler
	getUser    gr.Handler
	login      gr.Handler
}

func NewGrpcUserServer(userEnd *endpoint.UserEndpoint, logger kitlog.Logger) pb.USerServiceServer {
	return &gRPCServer{
		createUser: gr.NewServer(
			userEnd.CreateUserEndpoint(),
			DecodeCreateUserRequest,
			EncodeCreateUserResponse,
		),
		updateUser: gr.NewServer(
			userEnd.UpdateUserEndpoint(),
			DecodeUpdateUserRequest,
			EncodeUpdateUserResponse,
		),
		deleteUser: gr.NewServer(
			userEnd.DeleteUserEndpoint(),
			DecodeDeleteUserRequest,
			EncodeDeleteUserResponse,
		),
		getUser: gr.NewServer(
			userEnd.GetUserEndpoint(),
			DecodeGetUserRequest,
			EncodeGetResponse,
		),
		login: gr.NewServer(
			userEnd.LoginEndpoint(),
			DecodeLoginRequest,
			EncodeLoginResponse,
		),
	}
}
// func (s *gRPCServer) mustEmbedUnimplementedUSerServiceServer(a string) int {
// 	panic("fff")
// 	return 1
// }

func (s *gRPCServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, res, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.CreateUserResponse), nil
}

func (s *gRPCServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	_, res, err := s.updateUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.UpdateUserResponse), nil
}

func (s *gRPCServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, res, err := s.deleteUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.DeleteUserResponse), nil
}

func (s *gRPCServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	_, res, err := s.getUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.GetUserResponse), nil
}

func (s *gRPCServer) Login(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	_, res, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.LoginUserResponse), nil
}
