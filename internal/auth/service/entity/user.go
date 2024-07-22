package entity

import (
	"time"

	"github.com/yelaco/warble/util"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string
	Email        string
	Username     string
	FirstName    string
	LastName     string
	Password     string
	IsWarbleGold bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUser(email, userName, firstName, lastName, password string) User {
	return User{
		ID:           util.GenrateRandomID(),
		Email:        email,
		Username:     userName,
		FirstName:    firstName,
		LastName:     lastName,
		Password:     password,
		IsWarbleGold: false,
		CreatedAt:    time.Now(),
	}
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
