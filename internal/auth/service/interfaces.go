package service

import "github.com/yelaco/warble/internal/auth/service/entity"

type UserUsecaseInterface interface {
	CreateUser() (entity.User, error)
	Login() error
	FollowUser() error
	UpgradeWarbleGold() error
	UpdateUser() error
}

type UserRepositoryInterface interface {
	CreateUser(entity.User) (entity.User, error)
	UpdateUser(id, email, username, firstName, lastName, password string) (entity.User, error)
	UpgradeWarbleGold(id int) error
	GetUserByEmail(email string) (entity.User, error)
	GetUserByID(id string) (entity.User, error)
}
