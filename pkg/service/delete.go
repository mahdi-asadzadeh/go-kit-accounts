package service

import (
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/errors"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/service/models"
)

func (usrSer *UserService) DeleteUser(email string) (bool, error) {
	var user models.User
	err := usrSer.DB.Where("email = ?", email).Delete(&user).Error
	if err != nil {
		return false, errors.NotFound404
	}
	return true, nil
}
