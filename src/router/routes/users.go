package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Rota{
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Function:     controllers.FindUsers,
		AuthRequired: false,
	},
	{
		URI:          "/user/{id}",
		Method:       http.MethodGet,
		Function:     controllers.FindUser,
		AuthRequired: false,
	},
	{
		URI:          "/user/{id}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		AuthRequired: false,
	},
	{
		URI:          "/user/{id}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		AuthRequired: false,
	},
}
