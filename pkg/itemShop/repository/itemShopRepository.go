package repository

import "github.com/reyfear/game-shop-api-practice/entities"

type ItemShopRepository interface {
	Listing() ([]*entities.Item, error)
}
