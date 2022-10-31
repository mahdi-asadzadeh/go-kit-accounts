package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"column:email;UNIQUE; not null"`
	FullName string `gorm:"varchar(255);not null"`
	Password string `gorm:"column:password;not null"`
}

func (user *User) SetPassword(password string) string {
	if len(password) == 0 {
		return ""
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return string(passwordHash)
}

func (user *User) IsValidPassword(password string) error {
	bytePassword := []byte(password)
	byteHashPassword := []byte(user.Password)
	return bcrypt.CompareHashAndPassword(byteHashPassword, bytePassword)
}

func (user *User) GenerateJwtToken(jwtSecret string) string {
	jwt_token := jwt.New(jwt.SigningMethodHS512)
	jwt_token.Claims = jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24 * 90).Unix(),
	}
	token, _ := jwt_token.SignedString([]byte(jwtSecret))
	return token
}
