package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func (userEnd *UserEndpoint) DeleteUserEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteUserRequest)
		ok, err := userEnd.userService.DeleteUser(req.Email)
		if err != nil {
			return DeleteUserResponse{ok, err.Error()}, err
		}
		return DeleteUserResponse{ok, ""}, nil
	}
}
