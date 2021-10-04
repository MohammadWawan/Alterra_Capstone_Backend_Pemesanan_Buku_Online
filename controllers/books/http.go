package books

import (
	"alterra/business/books"
	"alterra/controllers"
	"alterra/controllers/books/requests"
	"alterra/controllers/books/responses"
	"alterra/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	BookUseCase books.Usecase
}

func NewBookController(bookUseCase books.Usecase) *BookController {
	return &BookController{
		BookUseCase: bookUseCase,
	}
}

func (bookController *BookController) GetBooks(c echo.Context) error {
	ctx := c.Request().Context()
	search := c.QueryParam("q")
	data, err := bookController.BookUseCase.GetListBook(ctx, search)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromBooksListDomain(data))
}

func (bookController *BookController) GetBookById(c echo.Context) error {
	request := c.Request().Context()
	id := c.Param("id")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvInt)
	}
	data, err := bookController.BookUseCase.GetById(request, uint(convInt))

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (bookController *BookController) UpdateBook(c echo.Context) error {
	id := c.Param("id")
	convId, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := requests.InsertBook{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := bookController.BookUseCase.Update(ctx, *req.ToDomain(), convId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (bookController *BookController) InsertBook(c echo.Context) error {
	request := requests.InsertBook{}
	var err error
	err = c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	ctx := c.Request().Context()
	var data books.Domain
	data, err = bookController.BookUseCase.InsertBook(ctx, *request.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (bookController *BookController) DeleteBook(c echo.Context) error {
	id := c.Param("id")
	idUint, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = bookController.BookUseCase.Delete(ctx, idUint)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.BookResponse{
		Id: idUint,
	})
}
