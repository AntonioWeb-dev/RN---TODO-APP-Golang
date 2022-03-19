package endpoints

import (
	"api/src/controllers"
)

var userEndpoints = []Route{
	{
		URI:            "/users",
		Method:         "POST",
		Controller:     controllers.CreateUser,
		Authentication: false,
	},
}
