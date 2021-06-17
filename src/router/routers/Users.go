package routers

import (
	"api/src/controllers"
	"net/http"
)

var usersRoute = []Route{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Func:   controllers.CreateUser,
		Auth:   false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Func:   controllers.GetUsers,
		Auth:   false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Func:   controllers.GetUser,
		Auth:   false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodPut,
		Func:   controllers.UpdateUser,
		Auth:   false,
	},
	{
		URI:    "/users/{useriD}",
		Method: http.MethodDelete,
		Func:   controllers.DeleteUser,
		Auth:   false,
	},
}
