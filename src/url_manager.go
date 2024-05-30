package src

// Gets url from the db.
func (app *App) internal_get_url(shortend_url string) (*Urls, error) {
	var url Urls
	result := app.Db.Where("shortend_url = ?", shortend_url).First(&url)
	if result.Error != nil {
		return nil, result.Error
	}
	return &url, nil
}

// sets url in the db.
func (app *App) internal_set_url(original_url string) (*Urls, error) {
	// create new url.
	new_url := Urls{ShortendUrl: shorten_url(original_url), OriginalUrl: original_url}
	result := app.Db.Create(&new_url)
	if result.Error != nil {
		return nil, result.Error
	}
	return &new_url, nil
}

func shorten_url(original_url string) string {
	// TODO: algorithm to calculate a unique shortend url.
	return "http" + original_url
}
