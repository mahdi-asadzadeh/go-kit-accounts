package service

import (
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/errors"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/service/models"
)

func (usrSer *UserService) CreateUser(email string, fullName string, password string) (usr *models.User, err error) {
	newUser := models.User{Email: email, FullName: fullName, Password: password}
	newUser.Password = newUser.SetPassword(password)
	err = usrSer.DB.Create(&newUser).Error
	if err != nil {
		return nil, errors.BadRequest400
	}
	return &newUser, nil
}
