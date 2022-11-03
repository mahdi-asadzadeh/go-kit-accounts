package tests

import (
	"fmt"
	"os"

	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/endpoint"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/service"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/service/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TestCommon struct {
	UserSer *service.UserService
	UserEnd *endpoint.UserEndpoint
}

var testCommon *TestCommon

func init() {
	os.Remove("gorm.db")
  
	// Implement user service
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

	newTestCommon := TestCommon{
		UserSer: userSer,
		UserEnd: userEnd,
	}
	testCommon = &newTestCommon
}
