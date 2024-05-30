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

func setupDbInstance() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./db.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect db. crashing.")
	}

	// Migrate the schema
	db.AutoMigrate(&Urls{})

	return db
}

type DbManager struct {
	Db *gorm.DB
}

func SetupDb() DbManager {
	return DbManager{Db: setupDbInstance()}

}

// Gets url from the db.
func (manager *DbManager) internal_get_url(shortend_url string) (*Urls, error) {
	var url Urls
	result := manager.Db.Where("shortend_url = ?", shortend_url).First(&url)
	if result.Error != nil {
		return nil, result.Error
	}
	return &url, nil
}

// sets url in the db.
func (manager *DbManager) internal_set_url(original_url string) (*Urls, error) {
	// create new url.
	new_url := Urls{ShortendUrl: shorten_url(original_url), OriginalUrl: original_url}
	result := manager.Db.Create(&new_url)
	if result.Error != nil {
		return nil, result.Error
	}
	return &new_url, nil
}
