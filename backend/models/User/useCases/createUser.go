package useCases

import (
	"api/models/User"
	"encoding/json"
)

type CreateUser struct {
	userRepository User.Repository
}

func InitCreateUserCase(userRepository User.Repository) *CreateUser {
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
