package main

import (
<<<<<<< HEAD
	"alterra/configs"
	"alterra/routes"
)

func main() {
	configs.InitDB()
	e := routes.NewRoute()
	e.Start(":8000")
=======
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name    string
	Address string
}

func main() {
	e := echo.New()
	e.GET("v1/users", GetUserController)
	e.Start(":8000")

}
func GetUserController(c echo.Context) error {
	user := User{"wawan", "Pasuruan"}
	return c.JSON(http.StatusOK, user)
>>>>>>> development
}
