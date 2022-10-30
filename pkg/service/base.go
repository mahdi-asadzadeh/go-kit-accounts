package service

import (
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/service/models"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	LoginUser(email string, password string) (string, error)
	CreateUser(email string, fullName string, password string) (*models.User, error)
	GetUser(email string) (*models.User, error)
	UpdateUser(email string, fullName string) (*models.User, error)
	DeleteUser(email string) (bool, error)
}

type UserService struct {
	JwtSecret string
	DB        *gorm.DB
}

func NewUserService(jwtSecret string, db *gorm.DB) *UserService {
	return &UserService{JwtSecret: jwtSecret, DB: db}
}
