package service

import (
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/errors"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/service/models"
)

func (usrSer *UserService) UpdateUser(email string, fullName string) (usr *models.User, err error) {
	getUser, err := usrSer.GetUser(email)
	if err != nil {
		return nil, errors.NotFound404
	}
	err = usrSer.DB.Model(&models.User{}).Where("email = ?", email).Update("full_name", fullName).Error
	if err != nil {
		return nil, errors.BadRequest400
	}
	getUser.FullName = fullName
	usr = getUser
	return usr, nil
}
