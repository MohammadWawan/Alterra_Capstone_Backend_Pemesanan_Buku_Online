package categories

import (
	"alterra/business/categories"
	"alterra/controllers"
	"alterra/controllers/categories/request"
	"alterra/controllers/categories/responses"
	"alterra/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	CategoryUseCase categories.Usecase
}

func NewCategoryController(categoryUseCase categories.Usecase) *CategoryController {
	return &CategoryController{
		CategoryUseCase: categoryUseCase,
	}
}

func (categoryController *CategoryController) GetCategories(c echo.Context) error {
	ctx := c.Request().Context()
	search := c.QueryParam("q")
	data, err := categoryController.CategoryUseCase.GetListCategory(ctx, search)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromCategoriesListDomain(data))
}

func (categoryController *CategoryController) GetCategoryById(c echo.Context) error {
	request := c.Request().Context()
	id := c.Param("id")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvInt)
	}
	data, err := categoryController.CategoryUseCase.GetById(request, uint(convInt))

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (categoryController *CategoryController) UpdateCategory(c echo.Context) error {
	id := c.Param("id")
	convId, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := request.InsertCategory{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := categoryController.CategoryUseCase.Update(ctx, *req.ToDomain(), convId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (categoryController *CategoryController) InsertCategory(c echo.Context) error {
	request := request.InsertCategory{}
	var err error
	err = c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	ctx := c.Request().Context()
	var data categories.Domain
	data, err = categoryController.CategoryUseCase.InsertCategory(ctx, *request.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (categoryController *CategoryController) DeleteCategory(c echo.Context) error {
	id := c.Param("id")
	idUint, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = categoryController.CategoryUseCase.Delete(ctx, idUint)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.CategoryResponse{
		Id: idUint,
	})
}
