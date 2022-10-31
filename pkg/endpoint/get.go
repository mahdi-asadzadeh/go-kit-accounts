package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func (userEnd *UserEndpoint) GetUserEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUserRequest)
		usr, err := userEnd.userService.GetUser(req.Email)
		if err != nil {
			return GetUserResponse{0, "", "", err.Error()}, err
		}
		return GetUserResponse{usr.ID, usr.Email, usr.FullName, ""}, err
	}
}
