package endpoints

import "api/adapters/rest/controllers/user"

var userEndpoints = []Route{
	{
		URI:            "/users",
		Method:         "POST",
		Controller:     user.CreateUser,
		Authentication: false,
	},
}
