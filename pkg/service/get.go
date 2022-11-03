package service

import (
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/errors"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/service/models"
)

func (usrSer *UserService) GetUser(email string) (usr *models.User, err error) {
	var user models.User
	err = usrSer.DB.Model(&user).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, errors.NotFound404
	}
	usr = &user
	return usr, nil
}
