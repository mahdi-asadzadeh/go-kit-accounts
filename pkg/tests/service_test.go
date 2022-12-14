package tests

import (
	"testing"

	"github.com/google/uuid"
)

func TestUserService(t *testing.T) {
	// Implement user data
	email := uuid.New().String()
	fullName := uuid.New().String()
	password := uuid.New().String()
	var userID uint
	var updateFullName string

	userSer := testCommon.UserSer

	t.Run("create-user", func(t *testing.T) {
		newUser, err := userSer.CreateUser(email, fullName, password)
		userID = newUser.ID

		if err != nil {
			t.Errorf("[Create user]: expected %v received %v", nil, err)
		}
		if newUser.Email != email {
			t.Errorf("got %s, wanted %s", newUser.Email, email)
		}
	})

	t.Run("get-user", func(t *testing.T) {
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

	t.Run("update-user", func(t *testing.T) {
		updateFullName = "My test"
		user, err := userSer.UpdateUser(email, updateFullName)

		if err != nil {
			t.Errorf("[Update user]: expected %v received %v", nil, err)
		}
		if user.FullName != updateFullName {
			t.Errorf("[Update user full name]: got %s, wanted %s", user.FullName, updateFullName)
		}
	})

	t.Run("delete-user", func(t *testing.T) {
		ok, err := userSer.DeleteUser(email)

		if err != nil {
			t.Errorf("[Delete user]: expected %v received %v", nil, err)
		}
		if ok != true {
			t.Errorf("[Delete user ok]: got %t, wanted %t", ok, true)
		}
	})
}
