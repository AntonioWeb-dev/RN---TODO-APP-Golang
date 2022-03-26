package useCases

import (
	"encoding/json"

	"api/helpers/hash"
	"api/helpers/jwt"
	"api/models/User"
)

type Login interface {
	Handler(data []byte) (loginStruct, error)
}

type loginCase struct {
	userRepo User.Repository
}

type loginStruct struct {
	Token string `json:"token,omitempty"`
	Email string `json:"email,omitempty"`
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
}

func InitLoginCase(userRepo User.Repository) Login {
	return &loginCase{userRepo: userRepo}
}

func (useCase *loginCase) Handler(data []byte) (loginStruct, error) {
	var user User.UserModel
	if err := json.Unmarshal(data, &user); err != nil {
		return loginStruct{}, nil
	}
	passwordNotHashed := user.Password
	user, err := useCase.userRepo.FindByEmail(user.Email)
	if err != nil {
		return loginStruct{}, err
	}
	if err := hash.CompareHash(passwordNotHashed, user.Password); err != nil {
		return loginStruct{}, err
	}
	token, err := jwt.GenerateJWT(user.ID, "secretLogin")
	if err != nil {
		return loginStruct{}, err
	}

	loginStructToReturn := loginStruct{
		Token: token,
		Email: user.Email,
		ID:    user.ID,
		Name:  user.Name,
	}
	return loginStructToReturn, nil
}
