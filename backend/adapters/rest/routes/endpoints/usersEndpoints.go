package endpoints

import (
	"api/adapters/rest/controllers/user"
	"api/models/User"
	"api/models/User/useCases"
)

func InitUserRoutes(userRepository User.Repository) []Route {
	createUserRequest := user.InitControllerCreateUser(useCases.InitCreateUserCase(userRepository))
	findAllUsersRequest := user.InitControllerFindAllUsers(useCases.InitFindAllUsersCase(userRepository))

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
	}
	return userEndpoints
}
