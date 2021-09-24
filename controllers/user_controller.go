package controllers

import (
	"alterra/configs"
	"alterra/models/response"
	"alterra/models/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterController(c echo.Context) error {
	var userRegister users.UserRegister
	c.Bind(&userRegister)

	// validasi
	if userRegister.Name == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Nama masih kosong",
			Data:    nil,
		})
	}

	// etc

	var userDB users.User
	userDB.Name = userRegister.Name
	userDB.Password = userRegister.Password
	userDB.Address = userRegister.Address
	userDB.Email = userRegister.Email

	result := configs.DB.Create(&userDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error ketika input data user ke DB",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil register",
		Data:    userDB,
	})
}

func LoginController(c echo.Context) error {
	userLogin := users.UserLogin{}
	c.Bind(&userLogin)
	// login

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    userLogin,
	})
}

func DetailUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Gagal konversi userId",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    users.User{Id: userId},
	})
}

func GetUserController(c echo.Context) error {

	users := []users.User{}

	result := configs.DB.Find(&users)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input mendapatkan data user dari DB",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil mendapatkan data user",
		Data:    users,
	})
}