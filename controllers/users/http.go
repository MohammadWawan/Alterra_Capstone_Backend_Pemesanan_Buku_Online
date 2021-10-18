package users

import (
	"alterra/business/users"
	"alterra/controllers"
	"alterra/controllers/users/requests"
	"alterra/controllers/users/responses"
	"alterra/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase users.Usecase
}

func NewUserController(userUseCase users.Usecase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (userController *UserController) Login(c echo.Context) error {
	var login users.Domain
	var err error
	var token string
	ctx := c.Request().Context()

	request := requests.UserLogin{}
	err = c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	login, token, err = userController.UserUseCase.Login(ctx, request.Email, request.Password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromUsersDomainToLogin(login, token))
}

func (userController *UserController) Register(c echo.Context) error {
	request := requests.UserRegister{}
	var err error
	err = c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	ctx := c.Request().Context()
	var data users.Domain
	data, err = userController.UserUseCase.Register(ctx, *request.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromUsersRegisterDomain(data))
}

func (userController *UserController) GetUsers(c echo.Context) error {
	request := c.Request().Context()
	user, err := userController.UserUseCase.GetAll(request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromUsersListDomain(user))
}

func (userController *UserController) GetDetailUser(c echo.Context) error {
	request := c.Request().Context()
	id := c.Param("id")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvInt)
	}
	data, err := userController.UserUseCase.GetById(request, uint(convInt))

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromUsersRegisterDomain(data))
}

func (userController *UserController) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	convId, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := requests.UserRegister{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := userController.UserUseCase.Update(ctx, *req.ToDomain(), convId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromUsersRegisterDomain(data))
}

func (userController *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	convId, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	request := c.Request().Context()
	err = userController.UserUseCase.Delete(request, convId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, nil)
}
