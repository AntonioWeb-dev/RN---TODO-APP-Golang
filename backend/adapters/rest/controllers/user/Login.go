package user

import (
	"io/ioutil"
	"net/http"

	"api/adapters/rest/controllers"
	"api/models/User/useCases"
)

type loginRequest struct {
	loginCase useCases.Login
}

func InitControllerLogin(loginCase useCases.Login) *loginRequest {
	return &loginRequest{loginCase}
}

func (controller *loginRequest) Handler(w http.ResponseWriter, r *http.Request) {
	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		controllers.Error(w, 422, err)
		return
	}
	tokenResponse, err := controller.loginCase.Handler(bodyReq)
	if err != nil {
		controllers.Error(w, 400, err)
		return
	}
	controllers.JSON(w, 200, tokenResponse)
}
