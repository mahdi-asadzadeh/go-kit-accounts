package tests

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/service"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/service/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserService(t *testing.T) {
	var db *gorm.DB
	JWTSECRET := "cHV87ewyuopvdXJh5rt8YXJ0ZWFjaGFuY2llbnRjb3JyZWN0bHlmZWxsb3dhcm1xdWE="
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	db.Debug().AutoMigrate(&models.User{})
	userSer := service.NewUserService(JWTSECRET, db)

	email := uuid.New().String()
	fullName := uuid.New().String()
	password := uuid.New().String()
	var userID uint
	var updateFullName string

	t.Run("Create user", func(t *testing.T) {
		newUser, err := userSer.CreateUser(email, fullName, password)
		userID = newUser.ID

		if err != nil {
			t.Errorf("[Create user]: expected %v received %v", nil, err)
		}
		if newUser.Email != email {
			t.Errorf("got %q, wanted %q", newUser.Email, email)
		}
	})

	t.Run("Get user", func(t *testing.T) {
		user, err := userSer.GetUser(email)

		if err != nil {
			t.Errorf("[Get user]: expected %v received %v", nil, err)
		}
		if user.Email != email {
			t.Errorf("[Get user email]: got %s, wanted %s", user.Email, email)
		}
		if user.ID != userID {
			t.Errorf("[Get user userid]: got %d, wanted %d", user.ID, userID)
		}
		if user.FullName != fullName {
			t.Errorf("[Get user full name]: got %s, wanted %s", user.FullName, fullName)
		}
	})

	t.Run("Update user", func(t *testing.T) {
		updateFullName = "My test"
		user, err := userSer.UpdateUser(email, updateFullName)

		if err != nil {
			t.Errorf("[Update user]: expected %v received %v", nil, err)
		}
		if user.FullName != updateFullName {
			t.Errorf("[Update user full name]: got %q, wanted %q", user.FullName, updateFullName)
		}
	})

	t.Run("Delete user", func(t *testing.T) {
		ok, err := userSer.DeleteUser(email)

		if err != nil {
			t.Errorf("[Delete user]: expected %v received %v", nil, err)
		}
		if ok != true {
			t.Errorf("[Delete user ok]: got %t, wanted %t", ok, true)
		}
	})
}
