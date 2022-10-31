package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/endpoint"
)

func DecodeCreateUserRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	var request endpoint.CreateUserRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeDeleteUserRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	email := req.URL.Query().Get("email")
	request := endpoint.DeleteUserRequest{Email: email}
	return request, nil
}

func DecodeGetUserRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	email := req.URL.Query().Get("email")
	request := endpoint.GetUserRequest{Email: email}
	return request, nil
}

func DecodeLoginUserRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	var request endpoint.LoginUserRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeUpdateUserRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	var request endpoint.UpdateUserRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
