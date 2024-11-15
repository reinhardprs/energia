package auth

import "energia/entities"

type AuthRepoInterface interface {
	Login(user entities.User) (entities.User, error)
	Register(user entities.User) (entities.User, error)
}
