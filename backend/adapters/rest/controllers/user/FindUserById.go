package user

import (
	"net/http"

	"api/adapters/rest/controllers"
	"api/models/User/useCases"

	"github.com/gorilla/mux"
)

type FindUserByIdRequest struct {
	findUserByIdCase *useCases.FindUserById
}

func InitControllerFindUserById(findUserByIdCase *useCases.FindUserById) *FindUserByIdRequest {
	return &FindUserByIdRequest{findUserByIdCase: findUserByIdCase}
}

func (controller *FindUserByIdRequest) Handler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, stts, err := controller.findUserByIdCase.Handler(params["id"])
	if err != nil {
		controllers.Error(w, stts, err)
		return
	}
	controllers.JSON(w, stts, user)
}
