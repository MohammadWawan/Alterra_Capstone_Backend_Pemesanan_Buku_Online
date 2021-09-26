package middlewares

import (
	"alterra/configs"
	"alterra/models/users"

	"github.com/labstack/echo/v4"
)

func Basicauth(email, password string, c echo.Context) (bool, error) {
	result := configs.DB.First(&users.User{}, "email = ? AND password = ?", email, password)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
