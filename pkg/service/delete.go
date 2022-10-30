package service

import (
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/service/models"
)

func (usrSer *UserService) DeleteUser(email string) (ok bool, err error) {
	var user models.User
	err = usrSer.DB.Where("email = ?", email).Delete(&user).Error
	if err != nil {
		ok = false
	} else {
		ok = true
	}
	return ok, err
}
