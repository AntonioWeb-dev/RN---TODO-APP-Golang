package user

import (
	"errors"
	"net/http"

	"api/helpers/response"
	"api/models/User/useCases"

	"github.com/gorilla/mux"
)

type FindUserByIdRequest struct {
	findUserByIdCase useCases.IFindUserByIdCase
}

func InitControllerFindUserById(findUserByIdCase useCases.IFindUserByIdCase) *FindUserByIdRequest {
	return &FindUserByIdRequest{findUserByIdCase: findUserByIdCase}
}

func (controller *FindUserByIdRequest) Handler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := r.Header.Get("userId")
	if userId != params["id"] {
		response.Error(w, http.StatusForbidden, errors.New("Forbidden"))
		return
	}
	user, stts, err := controller.findUserByIdCase.Handler(params["id"])
	if err != nil {
		response.Error(w, stts, err)
		return
	}
	response.JSON(w, stts, user)
}
