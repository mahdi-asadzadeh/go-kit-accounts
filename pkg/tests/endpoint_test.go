package tests

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/endpoint"
)

func TestUserEndpoint(t *testing.T) {
	// Implement user data
	email := uuid.New().String()
	fullName := uuid.New().String()
	password := uuid.New().String()

	userEnd := testCommon.UserEnd

	t.Run("create-user", func(t *testing.T) {
		createUser := userEnd.CreateUserEndpoint()
		req := endpoint.CreateUserRequest{email, fullName, password}
		_, err := createUser(context.Background(), req)
		if err != nil {
			t.Errorf("[Create user endpoint]: expected %v received %v", nil, err)
		}
	})

	t.Run("update-user", func(t *testing.T) {
		updateUser := userEnd.UpdateUserEndpoint()
		req := endpoint.UpdateUserRequest{email, "My test"}
		_, err := updateUser(context.Background(), req)
		if err != nil {
			t.Errorf("[Update user endpoint]: expected %v received %v", nil, err)
		}
	})

	t.Run("get-user", func(t *testing.T) {
		getUser := userEnd.GetUserEndpoint()
		req := endpoint.GetUserRequest{email}
		_, err := getUser(context.Background(), req)
		if err != nil {
			t.Errorf("[Get user endpoint]: expected %v received %v", nil, err)
		}
	})

	t.Run("login-user", func(t *testing.T) {
		loginUser := userEnd.LoginEndpoint()
		req := endpoint.LoginUserRequest{email, password}
		_, err := loginUser(context.Background(), req)
		if err != nil {
			t.Errorf("[Login user endpoint]: expected %v received %v", nil, err)
		}
	})

	t.Run("delete-user", func(t *testing.T) {
		deleteUser := userEnd.DeleteUserEndpoint()
		req := endpoint.DeleteUserRequest{email}
		_, err := deleteUser(context.Background(), req)
		if err != nil {
			t.Errorf("[Delete user endpoint]: expected %v received %v", nil, err)
		}
	})
}
