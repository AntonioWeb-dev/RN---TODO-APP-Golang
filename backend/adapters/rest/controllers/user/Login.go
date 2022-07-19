package user

import (
	"io/ioutil"
	"net/http"

	"api/helpers/response"
	"api/models/User/useCases"
)

type loginRequest struct {
	loginCase useCases.ILogin
}

func InitControllerLogin(loginCase useCases.ILogin) *loginRequest {
	return &loginRequest{loginCase}
}

func (controller *loginRequest) Handler(w http.ResponseWriter, r *http.Request) {
	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, 422, err)
		return
	}
	tokenResponse, err := controller.loginCase.Handler(bodyReq)
	if err != nil {
		response.Error(w, 400, err)
		return
	}
	response.JSON(w, 200, tokenResponse)
}
