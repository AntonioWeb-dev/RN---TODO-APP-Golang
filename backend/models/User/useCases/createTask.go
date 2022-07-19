package useCases

import (
	"encoding/json"

	"api/infra/repository/user"
	"api/models/Task"
)

type ICreateTask interface {
	Handler(userId string, data []byte) (error, int)
}

type createTaskCase struct {
	userRepo user.Repository
}

func InitCreateTaskCase(userRepo user.Repository) ICreateTask {
	return &createTaskCase{userRepo: userRepo}
}

func (useCase *createTaskCase) Handler(userId string, data []byte) (error, int) {
	var task *Task.TaskModel
	if err := json.Unmarshal(data, &task); err != nil {
		return err, 422
	}
	err := useCase.userRepo.CreateTask(userId, task)
	if err != nil {
		return err, 500
	}
	return nil, 201
}
