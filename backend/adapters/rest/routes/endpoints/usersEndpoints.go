package endpoints

import (
	"api/adapters/rest/controllers/task"
	"api/adapters/rest/controllers/user"
	"api/models/User"
	"api/models/User/useCases"
)

func InitUserRoutes(userRepository User.Repository) []Route {
	createUserRequest := user.InitControllerCreateUser(useCases.InitCreateUserCase(userRepository))
	findAllUsersRequest := user.InitControllerFindAllUsers(useCases.InitFindAllUsersCase(userRepository))
	findUserByIdRequest := user.InitControllerFindUserById(useCases.InitFindUserByIdCase(userRepository))
	createTaskRequest := task.InitControllerCreateTask(useCases.InitCreateTaskCase(userRepository))
	loginRequest := user.InitControllerLogin(useCases.InitLoginCase(userRepository))

	userEndpoints := []Route{
		{
			URI:            "/users",
			Method:         "POST",
			Controller:     createUserRequest.Handler,
			Authentication: false,
		},
		{
			URI:            "/users",
			Method:         "GET",
			Controller:     findAllUsersRequest.Handler,
			Authentication: false,
		},
		{
			URI:            "/users/{id}",
			Method:         "GET",
			Controller:     findUserByIdRequest.Handler,
			Authentication: false,
		},
		{
			URI:            "/users/login",
			Method:         "POST",
			Controller:     loginRequest.Handler,
			Authentication: false,
		},
		{
			URI:            "/users/{id}/tasks",
			Method:         "POST",
			Controller:     createTaskRequest.Handler,
			Authentication: false,
		},
	}
	return userEndpoints
}
