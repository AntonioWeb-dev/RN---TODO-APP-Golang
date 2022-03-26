package useCases

import (
	"api/models/User"
)

type FindAllUsers struct {
	userRepository User.Repository
}

func InitFindAllUsersCase(userRepository User.Repository) *FindAllUsers {
	return &FindAllUsers{userRepository}
}
func (useCase *FindAllUsers) Handler() ([]*User.UserModel, error) {
	users, err := useCase.userRepository.FindAll()
	if err != nil {
		return []*User.UserModel{}, err
	}
	return users, nil
}
