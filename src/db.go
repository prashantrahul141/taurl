package src

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Urls struct {
	ID          uint
	ShortendUrl string
	OriginalUrl string
	CreatedAt   time.Time
}

func SetupDbInstance() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./db.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect db. crashing.")
	}

	// Migrate the schema
	db.AutoMigrate(&Urls{})

	return db
}
