package karyawans

import (
	"alterra/business/karyawans"
	"alterra/controllers"
	"alterra/controllers/karyawans/requests"
	"alterra/controllers/karyawans/responses"
	"alterra/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type KaryawanController struct {
	KaryawanUseCase karyawans.Usecase
}

func NewKaryawanController(karyawanUseCase karyawans.Usecase) *KaryawanController {
	return &KaryawanController{
		KaryawanUseCase: karyawanUseCase,
	}
}

func (userController *KaryawanController) Login(c echo.Context) error {
	var login karyawans.Domain
	var err error
	var token string
	ctx := c.Request().Context()

	request := requests.KaryawanLogin{}
	err = c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	login, token, err = userController.KaryawanUseCase.Login(ctx, request.Email, request.Password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromKaryawansDomainToLogin(login, token))
}

func (karyawanController KaryawanController) Register(c echo.Context) error {
	request := requests.KaryawanRegister{}
	var err error
	err = c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	ctx := c.Request().Context()
	var data karyawans.Domain
	data, err = karyawanController.KaryawanUseCase.Register(ctx, *request.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromKaryawansDomain(data))
}

func (karyawanController *KaryawanController) GetKaryawans(c echo.Context) error {
	request := c.Request().Context()
	karyawan, err := karyawanController.KaryawanUseCase.GetAll(request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromKaryawansListDomain(karyawan))
}

func (karyawanController *KaryawanController) GetDetailKaryawan(c echo.Context) error {
	request := c.Request().Context()
	id := c.Param("id")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvInt)
	}
	data, err := karyawanController.KaryawanUseCase.GetById(request, uint(convInt))

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromKaryawansDomain(data))
}

func (karyawanController *KaryawanController) UpdateKaryawan(c echo.Context) error {
	id := c.Param("id")
	convId, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := requests.KaryawanRegister{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := karyawanController.KaryawanUseCase.Update(ctx, *req.ToDomain(), convId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromKaryawansDomain(data))
}

func (karyawanController *KaryawanController) DeleteKaryawan(c echo.Context) error {
	id := c.Param("id")
	convId, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	request := c.Request().Context()
	err = karyawanController.KaryawanUseCase.Delete(request, convId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, nil)
}
