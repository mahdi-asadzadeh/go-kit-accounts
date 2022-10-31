package service

import (
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/errors"
)

func (usrSer *UserService) LoginUser(email string, password string) (token string, err error) {
	user, err := usrSer.GetUser(email)
	if err != nil {
		return "", err
	}
	err = user.IsValidPassword(password)
	if err != nil {
		return "", errors.BadRequest400
	}
	token = user.GenerateJwtToken(usrSer.JwtSecret)
	return token, err
}
