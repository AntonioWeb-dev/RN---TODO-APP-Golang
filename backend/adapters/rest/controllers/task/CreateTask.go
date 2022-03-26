package task

import (
	"io/ioutil"
	"net/http"

	"api/adapters/rest/controllers"
	"api/models/User/useCases"

	"github.com/gorilla/mux"
)

type CreateTaskRequest struct {
	createTaskCase useCases.CreateTask
}

func InitControllerCreateTask(useCase useCases.CreateTask) *CreateTaskRequest {
	return &CreateTaskRequest{createTaskCase: useCase}
}

func (controller *CreateTaskRequest) Handler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		controllers.Error(w, 400, err)
		return
	}
	err, stts := controller.createTaskCase.Handler(params["id"], bodyReq)
	if err != nil {
		controllers.Error(w, stts, err)
		return
	}
	controllers.JSON(w, 201, nil)
}
