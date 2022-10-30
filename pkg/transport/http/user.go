package http

import (
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/endpoint"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHttpUserServer(userEnd *endpoint.UserEndpoint, logger kitlog.Logger) *mux.Router {
	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(EncodeErrorResponse),
	}
	createUserHandler := httptransport.NewServer(
		userEnd.CreateUserEndpoint(),
		DecodeCreateUserRequest,
		EncodeResponse,
		options...,
	)
	deleteUserHandler := httptransport.NewServer(
		userEnd.DeleteUserEndpoint(),
		DecodeDeleteUserRequest,
		EncodeResponse,
		options...,
	)
	getUserHandler := httptransport.NewServer(
		userEnd.GetUserEndpoint(),
		DecodeGetUserRequest,
		EncodeResponse,
		options...,
	)
	updateUserHandler := httptransport.NewServer(
		userEnd.UpdateUserEndpoint(),
		DecodeUpdateUserRequest,
		EncodeResponse,
		options...,
	)
	loginUserHandler := httptransport.NewServer(
		userEnd.LoginEndpoint(),
		DecodeLoginUserRequest,
		EncodeResponse,
		options...,
	)

	router := mux.NewRouter()
	router.Methods("POST").Path("/v1/register").Handler(createUserHandler)
	router.Methods("POST").Path("/v1/login").Handler(loginUserHandler)
	router.Methods("GET").Path("/v1/get").Handler(getUserHandler)
	router.Methods("DELETE").Path("/v1/delete").Handler(deleteUserHandler)
	router.Methods("PUT").Path("/v1/update").Handler(updateUserHandler)
	return router
}
