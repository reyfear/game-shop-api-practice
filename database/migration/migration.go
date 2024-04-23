package main

import (
	"github.com/reyfear/game-shop-api-practice/config"
	"github.com/reyfear/game-shop-api-practice/database"
	"github.com/reyfear/game-shop-api-practice/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := database.NewPostgresDatabase(conf.Database)

	tx := db.ConnectionGetting().Begin()
	playerMigration(tx)
	adminMigration(tx)
	itemMigration(tx)
	playerCoinMigration(tx)
	inventoryMigration(tx)
	purchaseHistoryMigration(tx)

	tx.Commit()
	if tx.Error != nil {
		tx.Rollback()
		panic(tx.Error)
	}
}

func playerMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Player{})
}

func adminMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Admin{})
}

func itemMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Item{})
}

func playerCoinMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.PlayerCoin{})
}

func inventoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Inventory{})
}

func purchaseHistoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.PurchaseHistory{})
}
