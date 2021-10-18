package wishlists

import (
	"alterra/business/wishlists"
	"alterra/controllers"
	"alterra/controllers/wishlists/requests"
	"alterra/controllers/wishlists/responses"
	"alterra/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type WishlistController struct {
	WishlistUseCase wishlists.Usecase
}

func NewWishlistController(wishlistUseCase wishlists.Usecase) *WishlistController {
	return &WishlistController{
		WishlistUseCase: wishlistUseCase,
	}
}

func (wishlistController *WishlistController) GetWishlists(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := wishlistController.WishlistUseCase.GetListWishlist(ctx, 1, 1)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromWishlistListDomain(data))
}

func (wishlistController *WishlistController) GetWishlistById(c echo.Context) error {
	request := c.Request().Context()
	id := c.Param("id")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvInt)
	}
	data, err := wishlistController.WishlistUseCase.GetById(request, uint(convInt))

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (wishlistController *WishlistController) UpdateWishlist(c echo.Context) error {
	id := c.Param("id")
	convId, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := requests.InsertWishlist{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := wishlistController.WishlistUseCase.Update(ctx, *req.ToDomain(), convId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (wishlistController *WishlistController) InsertWishlist(c echo.Context) error {
	request := requests.InsertWishlist{}
	var err error
	err = c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	ctx := c.Request().Context()
	var data wishlists.Domain
	data, err = wishlistController.WishlistUseCase.InsertWishlist(ctx, request.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromDomain(data))
}

func (wishlistController *WishlistController) DeleteWishlist(c echo.Context) error {
	id := c.Param("id")
	idUint, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = wishlistController.WishlistUseCase.Delete(ctx, idUint)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.WishlistResponse{
		Id: idUint,
	})
}
