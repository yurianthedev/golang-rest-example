package repository

import "github.com/yurianxdev/rest-example/api/model"

type UserRepository interface {
	ListUsers() ([]model.User, error)
	CreateUser(model.User) (model.User, error)
}
