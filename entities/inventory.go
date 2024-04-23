package entities

import "time"

type Inventory struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;"`
	PlayerID  string    `gorm:"type:varchar(64);not null;"`
	ItemID    uint64    `gorm:"not null;type:bigint;"`
	IsDeleted bool      `gorm:"not null;default:false;"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null;"`
}
