package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func (userEnd *UserEndpoint) CreateUserEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateUserRequest)
		usr, err := userEnd.userService.CreateUser(req.Email, req.FullName, req.Password)
		if err != nil {
			return CreateUserResponse{0, "", "", err.Error()}, err
		}
		return CreateUserResponse{ID: usr.ID, Email: usr.Email, FullName: usr.FullName, Err: ""}, nil
	}
}
