package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *App) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "Taurl"})
}

type RedirectReq struct {
	Id string `uri:"id" binding:"required"`
}

func (app *App) Redirect(c *gin.Context) {
	var req RedirectReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	url, err := app.Db.get_url_from_id(req.Id)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.tmpl", gin.H{"title": "404"})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, url.OriginalUrl)
}
