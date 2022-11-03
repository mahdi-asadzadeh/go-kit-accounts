package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func (userEnd *UserEndpoint) LoginEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(LoginUserRequest)
		token, err := userEnd.userService.LoginUser(req.Email, req.Password)
		if err != nil {
			return LoginUserResponse{"", err.Error()}, err
		}
		return LoginUserResponse{token, ""}, nil
	}
}
