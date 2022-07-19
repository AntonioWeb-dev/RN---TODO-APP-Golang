package user

import (
	"net/http"

	"api/helpers/response"
	"api/models/User/useCases"
)

type FindAllUsersRequest struct {
	findAllUsersCase useCases.IFindAllUsersCase
}

func InitControllerFindAllUsers(findAllUsersCase useCases.IFindAllUsersCase) *FindAllUsersRequest {
	return &FindAllUsersRequest{findAllUsersCase}
}

func (controller *FindAllUsersRequest) Handler(w http.ResponseWriter, r *http.Request) {
	users, err := controller.findAllUsersCase.Handler()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusCreated, users)
}
