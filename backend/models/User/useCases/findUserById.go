package useCases

import (
	"api/infra/repository/user"
	"api/models/User"
)

type IFindUserByIdCase interface {
	Handler(id string) (*User.UserModel, int, error)
}

type FindUserById struct {
	userRepository user.Repository
}

func InitFindUserByIdCase(userRepository user.Repository) IFindUserByIdCase {
	return &FindUserById{userRepository: userRepository}
}

func (useCase *FindUserById) Handler(id string) (*User.UserModel, int, error) {
	user, err := useCase.userRepository.FindByID(id)
	if err != nil {
		return &User.UserModel{}, 404, err
	}
	return &user, 200, nil
}
