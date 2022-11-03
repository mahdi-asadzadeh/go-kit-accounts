package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/mahdi-asadzadeh/go-kit-accounts/clients/types"
	transporthttp "github.com/mahdi-asadzadeh/go-kit-accounts/pkg/transport/http"
)

func TestUserTransportHTTP(t *testing.T) {
	httpServer := transporthttp.NewHttpUserServer(testCommon.UserEnd, nil)
	srv := httptest.NewServer(httpServer)

	// Implement user data
	email := uuid.New().String()
	fullName := uuid.New().String()
	password := uuid.New().String()
	var newFullName string
	var userID uint

	t.Run("create-user", func(t *testing.T) {
		input := types.CreateUserInput{Email: email, FullName: fullName, Password: password}
		json_data, _ := json.Marshal(&input)
		req, _ := http.NewRequest("POST", srv.URL+"/v1/register", bytes.NewBuffer(json_data))
		resp, _ := http.DefaultClient.Do(req)

		// Check StatusCode
		if resp.StatusCode != 200 {
			t.Errorf("%s %s %s: want %d have %d", "POST", srv.URL+"/v1/register", json_data, 200, resp.StatusCode)
		}

		// Check response
		var result types.CreateUserResponse
		json.NewDecoder(resp.Body).Decode(&result)
		if result.Email != email {
			t.Errorf("want %s have %s", email, result.Email)
		}
		if result.FullName != fullName {
			t.Errorf("want %s, have %s", fullName, result.FullName)
		}
		userID = result.ID
	})

	t.Run("update-user", func(t *testing.T) {
		newFullName = "My test"
		input := types.UpdateUserInput{Email: email, FullName: newFullName}
		json_data, _ := json.Marshal(&input)
		req, _ := http.NewRequest("PUT", srv.URL+"/v1/update", bytes.NewBuffer(json_data))
		resp, _ := http.DefaultClient.Do(req)

		// Check StatusCode
		if resp.StatusCode != 200 {
			t.Errorf("%s %s %s: want %d have %d", "POST", srv.URL+"/v1/update", json_data, 200, resp.StatusCode)
		}

		// Check response
		var result types.UpdateUserResponse
		json.NewDecoder(resp.Body).Decode(&result)
		if result.FullName != newFullName {
			t.Errorf("want %s, have %s", newFullName, result.FullName)
		}
	})

	t.Run("get-user", func(t *testing.T) {
		req, _ := http.NewRequest("GET", srv.URL+"/v1/get?email="+email, nil)
		resp, _ := http.DefaultClient.Do(req)

		// Check StatusCode
		if resp.StatusCode != 200 {
			t.Errorf("%s %s : want %d have %d", "POST", srv.URL+"/v1/update", 200, resp.StatusCode)
		}

		// Check response
		var result types.GetUserResponse
		json.NewDecoder(resp.Body).Decode(&result)
		if result.ID != userID {
			t.Errorf("want %d, have %d", userID, result.ID)
		}
		if result.Email != email {
			t.Errorf("want %s, have %s", email, result.Email)
		}
		if result.FullName != newFullName {
			t.Errorf("want %s, have %s", newFullName, result.FullName)
		}
	})

	t.Run("login", func(t *testing.T) {
		input := types.LoginUserInput{Email: email, Password: password}
		json_data, _ := json.Marshal(&input)
		req, _ := http.NewRequest("POST", srv.URL+"/v1/login", bytes.NewBuffer(json_data))
		resp, _ := http.DefaultClient.Do(req)

		// Check StatusCode
		if resp.StatusCode != 200 {
			t.Errorf("%s %s %s: want %d have %d", "POST", srv.URL+"/v1/login", json_data, 200, resp.StatusCode)
		}

		// Check response
		var result types.LoginUserResponse
		json.NewDecoder(resp.Body).Decode(&result)
		if result.Token == "" {
			t.Errorf("have %s", result.Token)
		}
	})

	t.Run("delete-user", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", srv.URL+"/v1/delete?email="+email, nil)
		resp, _ := http.DefaultClient.Do(req)

		// Check StatusCode
		if resp.StatusCode != 200 {
			t.Errorf("%s %s : want %d have %d", "POST", srv.URL+"/v1/delete?email="+email, 200, resp.StatusCode)
		}

		// Check response
		var result types.DeleteUserResponse
		json.NewDecoder(resp.Body).Decode(&result)
		if result.Ok != true {
			t.Errorf("want %t, have %t", true, result.Ok)
		}
	})
}
