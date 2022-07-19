package user

import (
	"io/ioutil"
	"net/http"

	"api/helpers/response"
	"api/models/User/useCases"
)

type CreateUserRequest struct {
	createUserCase useCases.ICreateUserCase
}

func InitControllerCreateUser(createUserCase useCases.ICreateUserCase) *CreateUserRequest {
	return &CreateUserRequest{createUserCase}
}

func (controller *CreateUserRequest) Handler(w http.ResponseWriter, r *http.Request) {
	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, 400, err)
		return
	}
	err, stts := controller.createUserCase.Handler(bodyReq)
	if err != nil {
		response.Error(w, stts, err)
		return
	}

	response.JSON(w, http.StatusCreated, nil)
}
