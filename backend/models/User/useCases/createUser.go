package useCases

import (
	"api/infra/repository/user"
	"api/models/User"
	"encoding/json"
)

type ICreateUserCase interface {
	Handler(data []byte) (error, int)
}

type CreateUser struct {
	userRepository user.Repository
}

func InitCreateUserCase(userRepository user.Repository) ICreateUserCase {
	return &CreateUser{userRepository}
}

func (useCase *CreateUser) Handler(data []byte) (error, int) {
	var userToCreate User.UserModel
	if err := json.Unmarshal(data, &userToCreate); err != nil {
		return err, 422
	}
	if err := userToCreate.Prepare("create"); err != nil {
		return err, 400
	}
	err := useCase.userRepository.CreateUser(&userToCreate)
	if err != nil {
		return err, 500
	}
	return nil, 200
}
