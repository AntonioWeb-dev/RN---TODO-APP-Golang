package task

import (
	"io/ioutil"
	"net/http"

	"api/helpers/response"
	"api/models/User/useCases"

	"github.com/gorilla/mux"
)

type CreateTaskRequest struct {
	createTaskCase useCases.ICreateTask
}

func InitControllerCreateTask(useCase useCases.ICreateTask) *CreateTaskRequest {
	return &CreateTaskRequest{createTaskCase: useCase}
}

func (controller *CreateTaskRequest) Handler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Error(w, 400, err)
		return
	}
	err, stts := controller.createTaskCase.Handler(params["id"], bodyReq)
	if err != nil {
		response.Error(w, stts, err)
		return
	}
	response.JSON(w, 201, nil)
}
