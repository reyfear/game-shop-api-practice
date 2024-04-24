package service

import (
	_itemShopModel "github.com/reyfear/game-shop-api-practice/pkg/itemShop/model"
)

type ItemShopService interface {
	Listing() ([]*_itemShopModel.Item, error)
}
