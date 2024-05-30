package app

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Models the database.
type Urls struct {
	ID          uint
	ShortendUrl string // http://host.com/hash_xyz
	OriginalUrl string // http://areallyreallylongexampleurl.com
	UniqueId    string `gorm:"unique"` // hash_xyz
	CreatedAt   time.Time
}

// Inits a new gorm.DB connection to the local db file.
// Return pointer to the created db instance.
func setupDbInstance() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./db.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect db. crashing.")
	}

	// Migrate the schema
	db.AutoMigrate(&Urls{})

	return db
}

// Wrapper for db, this will be attached to global app.
type DbManager struct {
	// Pointer to the connection instance.
	Db *gorm.DB
	// for in memory caching
	Cache map[string]Urls // UniqueId: UrlsObject
}

// Public init function for DbManager
func SetupDb() DbManager {
	return DbManager{Db: setupDbInstance(), Cache: make(map[string]Urls)}
}

// simple wrapper might need it later.
func (manager *DbManager) get_from_cache(key string) Urls {
	return manager.Cache[key]
}

// Gets url from the db using shortend url.
// considers cache first and also updates it.
func (manager *DbManager) get_url(shortend_url string) (*Urls, error) {
	var url Urls = manager.get_from_cache(shortend_url)

	// we found from cache
	if url.ID != 0 {
		return &url, nil
	}

	result := manager.Db.Where("shortend_url = ?", shortend_url).First(&url)
	if result.Error != nil {
		return nil, result.Error
	}

	manager.Cache[url.UniqueId] = url
	return &url, nil
}

// Gets url from the db using unique id.
// considers cache first and also updates it.
func (manager *DbManager) get_url_from_id(uniqueId string) (*Urls, error) {
	var url Urls = manager.get_from_cache(uniqueId)

	// we found from cache
	if url.ID != 0 {
		return &url, nil
	}

	result := manager.Db.Where("unique_id = ?", uniqueId).First(&url)
	if result.Error != nil {
		return nil, result.Error
	}

	// cache before returning
	manager.Cache[url.UniqueId] = url
	return &url, nil
}

// stores a new url in the db and also in the inmemory cache.
func (manager *DbManager) set_url(original_url string) (*Urls, error) {
	uniqueId := shorten_url(original_url)

	new_url := Urls{ShortendUrl: "http://localhost:3000/" + uniqueId, OriginalUrl: original_url, UniqueId: uniqueId}
	result := manager.Db.Create(&new_url)
	if result.Error != nil {
		return nil, result.Error
	}

	manager.Cache[uniqueId] = new_url
	return &new_url, nil
}
