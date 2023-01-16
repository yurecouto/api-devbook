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
		Auth:     true,
	},
	{
		Uri:      "/usuarios/{usuarioId}",
		Method:   http.MethodGet,
		Function: controllers.GetUser,
		Auth:     true,
	},
	{
		Uri:      "/usuarios/{usuarioId}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		Auth:     true,
	},
	{
		Uri:      "/usuarios/{usuarioId}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		Auth:     true,
	},
	{
		Uri:      "/usuarios/{usuarioId}/follow",
		Method:   http.MethodPost,
		Function: controllers.FollowUser,
		Auth:     true,
	},
	{
		Uri:      "/usuarios/{usuarioId}/unfollow",
		Method:   http.MethodPost,
		Function: controllers.UnFollowUser,
		Auth:     true,
	},
	{
		Uri:      "/usuarios/{usuarioId}/followers",
		Method:   http.MethodGet,
		Function: controllers.FindFollowers,
		Auth:     true,
	},
	{
		Uri:      "/usuarios/{usuarioId}/following",
		Method:   http.MethodGet,
		Function: controllers.FindFollowing,
		Auth:     true,
	},
	{
		Uri:      "/usuarios/{usuarioId}/update-password",
		Method:   http.MethodPost,
		Function: controllers.UpdatePassword,
		Auth:     true,
	},
}
