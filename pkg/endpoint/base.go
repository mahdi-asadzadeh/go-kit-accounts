package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/service"
)

type UserEndpointInterface interface {
	CreateUserEndpoint() endpoint.Endpoint
	UpdateUserEndpoint() endpoint.Endpoint
	DeleteUserEndpoint() endpoint.Endpoint
	GetUserEndpoint() endpoint.Endpoint
	LoginEndpoint() endpoint.Endpoint
}

type UserEndpoint struct {
	userService *service.UserService
}

func NewUserEndpoint(userSer *service.UserService) *UserEndpoint {
	return &UserEndpoint{userService: userSer}
}
