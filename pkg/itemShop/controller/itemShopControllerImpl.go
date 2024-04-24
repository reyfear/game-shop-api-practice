package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_itemShopService "github.com/reyfear/game-shop-api-practice/pkg/itemShop/service"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {
	itemModelList, err := c.itemShopService.Listing()
	if err != nil {
		return pctx.String(http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, itemModelList)
}
