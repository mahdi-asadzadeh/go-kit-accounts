package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func (userEnd *UserEndpoint) UpdateUserEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateUserRequest)
		usr, err := userEnd.userService.UpdateUser(req.Email, req.FullName)
		if err != nil {
			return UpdateUserResponse{0, "", "", err.Error()}, err
		}
		return UpdateUserResponse{usr.ID, usr.Email, usr.FullName, ""}, err
	}
}
