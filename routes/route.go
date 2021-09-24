package routes

import (
	"alterra/controllers"

	"github.com/labstack/echo/v4"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	ev1 := e.Group("api/v1/")
	ev1.GET("users", controllers.GetUserController)
	ev1.POST("users/login", controllers.LoginController)
	ev1.POST("users/register", controllers.RegisterController)
	ev1.GET("users/:userId", controllers.DetailUserController)

	ev1.GET("karyawans", controllers.GetKaryawanController)
	ev1.POST("karyawans/login", controllers.KaryawanLoginController)
	ev1.POST("karyawans/register", controllers.KaryawanRegisterController)
	ev1.GET("karyawans/:karyawanId", controllers.DetailKaryawanController)
	return e
}
