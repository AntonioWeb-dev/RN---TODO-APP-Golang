package user

import (
	"net/http"

	"api/adapters/rest/controllers"
	"api/models/User/useCases"
)

type FindAllUsersRequest struct {
	findAllUsersCase *useCases.FindAllUsers
}

func InitControllerFindAllUsers(findAllUsersCase *useCases.FindAllUsers) *FindAllUsersRequest {
	return &FindAllUsersRequest{findAllUsersCase}
}

func (controller *FindAllUsersRequest) Handler(w http.ResponseWriter, r *http.Request) {
	users, err := controller.findAllUsersCase.Handler()
	if err != nil {
		controllers.Error(w, http.StatusInternalServerError, err)
		return
	}
	controllers.JSON(w, http.StatusCreated, users)
}
