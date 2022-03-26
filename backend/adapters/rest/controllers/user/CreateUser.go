package user

import (
	"io/ioutil"
	"net/http"

	"api/adapters/rest/controllers"
	"api/models/User/useCases"
)

type CreateUserRequest struct {
	createUserCase *useCases.CreateUser
}

func InitControllerCreateUser(createUserCase *useCases.CreateUser) *CreateUserRequest {
	return &CreateUserRequest{createUserCase}
}

func (controller *CreateUserRequest) Handler(w http.ResponseWriter, r *http.Request) {
	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		controllers.Error(w, 400, err)
		return
	}
	err, stts := controller.createUserCase.Handler(bodyReq)
	if err != nil {
		controllers.Error(w, stts, err)
		return
	}

	controllers.JSON(w, http.StatusCreated, nil)
}
