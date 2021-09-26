package controllers

import (
	"alterra/configs"
	"alterra/models/karyawans"
	"alterra/models/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func KaryawanRegisterController(c echo.Context) error {
	var karyawanRegister karyawans.KaryawanRegister
	c.Bind(&karyawanRegister)

	// validasi
	if karyawanRegister.Name == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Nama masih kosong",
			Data:    nil,
		})
	}

	// etc

	var karyawanDB karyawans.Karyawan
	karyawanDB.Name = karyawanRegister.Name
	karyawanDB.Email = karyawanRegister.Email
	karyawanDB.Password = karyawanRegister.Password

	result := configs.DB.Create(&karyawanDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error ketika input data karyawan ke DB",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil register",
		Data:    karyawanDB,
	})
}

func KaryawanLoginController(c echo.Context) error {
	karyawanLogin := karyawans.KaryawanLogin{}
	c.Bind(&karyawanLogin)
	// login

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    karyawanLogin,
	})
}

func DetailKaryawanController(c echo.Context) error {
	karyawanId, err := strconv.Atoi(c.Param("karyawanId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Gagal konversi karyawanId",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
		Data:    karyawans.Karyawan{Id: karyawanId},
	})
}

func GetKaryawanController(c echo.Context) error {

	karyawans := []karyawans.Karyawan{}

	result := configs.DB.Find(&karyawans)

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
		Data:    karyawans,
	})
}
