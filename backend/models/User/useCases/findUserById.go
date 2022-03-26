package useCases

import (
	"api/models/User"
)

type FindUserById struct {
	userRepository User.Repository
}

func InitFindUserByIdCase(userRepository User.Repository) *FindUserById {
	return &FindUserById{userRepository: userRepository}
}

func (useCase *FindUserById) Handler(id string) (*User.UserModel, int, error) {
	user, err := useCase.userRepository.FindByID(id)
	if err != nil {
		return &User.UserModel{}, 404, err
	}
	return &user, 200, nil
}
