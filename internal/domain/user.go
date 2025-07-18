package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

type AuthUseCase interface {
	Register(email, password string) error
	Login(email, password string) (string, error)
	RefreshToken(token string) (string, error)
	Logout(token string) error
}

type UserRepository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
}