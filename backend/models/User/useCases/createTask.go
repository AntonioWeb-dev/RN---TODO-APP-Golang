package useCases

import (
	"encoding/json"

	"api/models/User"
)

type CreateTask interface {
	Handler(userId string, data []byte) (error, int)
}

type createTaskCase struct {
	userRepo User.Repository
}

func InitCreateTaskCase(userRepo User.Repository) CreateTask {
	return &createTaskCase{userRepo: userRepo}
}

func (useCase *createTaskCase) Handler(userId string, data []byte) (error, int) {
	var task *User.Task
	if err := json.Unmarshal(data, &task); err != nil {
		return err, 422
	}
	err := useCase.userRepo.CreateTask(userId, task)
	if err != nil {
		return err, 500
	}
	return nil, 201
}
