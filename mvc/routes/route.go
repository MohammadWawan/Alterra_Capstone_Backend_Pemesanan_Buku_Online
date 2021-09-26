package routes

import (
	"alterra/constants"
	"alterra/controllers"
	"alterra/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.BodyDump(middlewares.Log))

	jwt := middleware.JWT([]byte(constants.SECRET_JWT))

	ev1 := e.Group("api/v1/")
	ev1.GET("users", controllers.GetUserController, jwt)
	ev1.POST("users/login", controllers.LoginController)
	ev1.POST("users/register", controllers.RegisterController)
	ev1.GET("users/:userId", controllers.DetailUserController)

	ev1.GET("karyawans", controllers.GetKaryawanController)
	ev1.POST("karyawans/login", controllers.KaryawanLoginController)
	ev1.POST("karyawans/register", controllers.KaryawanRegisterController)
	ev1.GET("karyawans/:karyawanId", controllers.DetailKaryawanController)

	ev1.GET("addBooks", controllers.AddBookController)
	return e
}
