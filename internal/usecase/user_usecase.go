package usecase

import "miniDriver/internal/model"

type UserUseCase interface {
	Register(user *model.User) error
	Login(username, password string) (string, error)
}
