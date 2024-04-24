package main

import (
	"github.com/reyfear/game-shop-api-practice/config"
	"github.com/reyfear/game-shop-api-practice/database"
	"github.com/reyfear/game-shop-api-practice/server"
)

func main() {
	conf := config.ConfigGetting()
	db := database.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db.ConnectionGetting())

	server.Start()
}
