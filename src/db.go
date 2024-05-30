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
	UniqueId    string
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

// Gets url from the db using shortend url.
func (manager *DbManager) get_url(shortend_url string) (*Urls, error) {
	var url Urls
	result := manager.Db.Where("shortend_url = ?", shortend_url).First(&url)
	if result.Error != nil {
		return nil, result.Error
	}
	return &url, nil
}

// Gets url from the db using unique id.
func (manager *DbManager) get_url_from_id(uniqueId string) (*Urls, error) {
	var url Urls
	result := manager.Db.Where("unique_id = ?", uniqueId).First(&url)
	if result.Error != nil {
		return nil, result.Error
	}
	return &url, nil
}

// sets url in the db.
func (manager *DbManager) set_url(original_url string) (*Urls, error) {
	// create new url.
	uniqueId := shorten_url(original_url)
	new_url := Urls{ShortendUrl: "http://localhost:3000/" + uniqueId, OriginalUrl: original_url, UniqueId: uniqueId}
	result := manager.Db.Create(&new_url)
	if result.Error != nil {
		return nil, result.Error
	}
	return &new_url, nil
}
