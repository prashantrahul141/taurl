package src

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type UrlReq struct {
	Url string `binding:"required"`
}

// api endpoint to get an already existing url, can return 404
func (app *App) ApiGetUrl(c *gin.Context) {
	// validate req body.
	var json UrlReq
	if err := c.ShouldBindQuery(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	url, err := app.internal_get_url(json.Url)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Url was not found."})
		return
	}

	c.JSON(http.StatusOK, url)
}

// creates a new Url using given original url
func (app *App) ApiSetUrl(c *gin.Context) {
	// validate req body.
	var json UrlReq
	if err := c.ShouldBindBodyWithJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// if the given url is not a valid url.
	_, err := url.ParseRequestURI(json.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse url."})
		return
	}

	new_url, err := app.internal_set_url(json.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse url."})
		return
	}

	c.JSON(http.StatusCreated, new_url)
}
