package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/endpoint"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/service"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/service/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserEndpoint(t *testing.T) {
	// Implement user service
	email := uuid.New().String()
	fullName := uuid.New().String()
	password := uuid.New().String()
	// var userID uint
	// var updateFullName string
	var db *gorm.DB

	JWTSECRET := "cHV87ewyuopvdXJh5rt8YXJ0ZWFjaGFuY2llbnRjb3JyZWN0bHlmZWxsb3dhcm1xdWE="
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	db.Debug().AutoMigrate(&models.User{})
	userSer := service.NewUserService(JWTSECRET, db)

	// Implement user endpoint
	userEnd := endpoint.NewUserEndpoint(userSer)

	t.Run("Create user endpoint", func(t *testing.T) {
		createUser := userEnd.CreateUserEndpoint()
		req := endpoint.CreateUserRequest{email, fullName, password}
		_, err := createUser(context.Background(), req)
		if err != nil {
			t.Errorf("[Create user endpoint]: expected %v received %v", nil, err)
		}
	})

	t.Run("Update user endpoint", func(t *testing.T) {
		updateUser := userEnd.UpdateUserEndpoint()
		req := endpoint.UpdateUserRequest{email, "My test"}
		_, err := updateUser(context.Background(), req)
		if err != nil {
			t.Errorf("[Update user endpoint]: expected %v received %v", nil, err)
		}
	})

	t.Run("Get user endpoint", func(t *testing.T) {
		getUser := userEnd.GetUserEndpoint()
		req := endpoint.GetUserRequest{email}
		_, err := getUser(context.Background(), req)
		if err != nil {
			t.Errorf("[Get user endpoint]: expected %v received %v", nil, err)
		}
	})

	t.Run("Login user endpoint", func(t *testing.T) {
		loginUser := userEnd.LoginEndpoint()
		req := endpoint.LoginUserRequest{email, password}
		_, err := loginUser(context.Background(), req)
		if err != nil {
			t.Errorf("[Login user endpoint]: expected %v received %v", nil, err)
		}
	})

	t.Run("Delete user endpoint", func(t *testing.T) {
		deleteUser := userEnd.DeleteUserEndpoint()
		req := endpoint.DeleteUserRequest{email}
		_, err := deleteUser(context.Background(), req)
		if err != nil {
			t.Errorf("[Delete user endpoint]: expected %v received %v", nil, err)
		}
	})
}
