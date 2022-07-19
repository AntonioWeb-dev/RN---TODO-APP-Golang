package useCases

import (
	"api/infra/repository/user"
	"api/models/User"
)

type IFindAllUsersCase interface {
	Handler() ([]*User.UserModel, error)
}

type FindAllUsers struct {
	userRepository user.Repository
}

func InitFindAllUsersCase(userRepository user.Repository) IFindAllUsersCase {
	return &FindAllUsers{userRepository}
}
func (useCase *FindAllUsers) Handler() ([]*User.UserModel, error) {
	users, err := useCase.userRepository.FindAll()
	if err != nil {
		return []*User.UserModel{}, err
	}
	return users, nil
}
