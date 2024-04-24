package server

import (
	_itemShopController "github.com/reyfear/game-shop-api-practice/pkg/itemShop/controller"
	_itemShopRepository "github.com/reyfear/game-shop-api-practice/pkg/itemShop/repository"
	_itemShopService "github.com/reyfear/game-shop-api-practice/pkg/itemShop/service"
)

func (s *echoServer) initItemShopRouter() {
	router := s.app.Group("/v1/item-shop")

	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemShopService := _itemShopService.NewItemShopServiceImpl(itemShopRepository)
	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router.GET("/listing", itemShopController.Listing)
}
