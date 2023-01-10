package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		Uri:      "/usuarios",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
		Auth:     false,
	},
	{
		Uri:      "/usuarios",
		Method:   http.MethodGet,
		Function: controllers.GetUsers,
		Auth:     false,
	},
	{
		Uri:      "/usuarios/{usuarioId}",
		Method:   http.MethodGet,
		Function: controllers.GetUser,
		Auth:     false,
	},
	{
		Uri:      "/usuarios/{usuarioId}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		Auth:     false,
	},
	{
		Uri:      "/usuarios/{usuarioId}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		Auth:     false,
	},
}