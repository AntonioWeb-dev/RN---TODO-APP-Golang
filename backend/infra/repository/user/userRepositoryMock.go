package user

import (
	"api/models/Task"
	"api/models/User"
)

type User_mongo_mock interface {
	CreateUser(user *User.UserModel) error
	FindAll() ([]*User.UserModel, error)
	FindByID(id string) (User.UserModel, error)
	CreateTask(userId string, taskData *Task.TaskModel) error
	FindByEmail(email string) (User.UserModel, error)
}

type repository_mock struct{}

func InitUserRepositoryMock() User_mongo_mock {
	return &repository_mock{}
}

func (m *repository_mock) CreateUser(user *User.UserModel) error {
	return nil
}

func (m *repository_mock) FindAll() ([]*User.UserModel, error) {
	return []*User.UserModel{}, nil
}

func (m *repository_mock) FindByID(userId string) (User.UserModel, error) {
	return User.UserModel{
		ID:       "623f525b9ae2a09765214a81",
		Name:     "junior",
		Username: "user",
		Tasks:    nil,
		Email:    "antoniovs420@gmail.com",
	}, nil
}
func (m *repository_mock) CreateTask(userId string, taskData *Task.TaskModel) error {
	return nil
}
func (m *repository_mock) FindByEmail(email string) (User.UserModel, error) {
	return User.UserModel{}, nil
}
