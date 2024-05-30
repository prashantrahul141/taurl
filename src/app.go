package src

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router     *gin.Engine
	UrlManager UrlsManager
}

// Inits an app.
func (app *App) Init() {
	slog.Info("Init app")
	app.Router.Static("/static", "./assets/static")
	app.Router.StaticFile("/favicon.ico", "./assets/favicon.ico")
	app.Router.LoadHTMLGlob("./assets/templates/*")
}

// mount routes
func (app *App) MountRoutes() {
	app.Router.GET("/", Index)
	app.Router.GET("/get", app.UrlManager.GetUrl)
	app.Router.POST("/set", app.UrlManager.SetUrl)
}

// creates and inits a new app.
func Default() App {
	app := App{Router: gin.Default(), UrlManager: UrlsManager{Db: SetupDbInstance()}}
	app.Init()
	app.MountRoutes()
	return app
}

// starts all services.
func (app *App) Run() {
	slog.Info("Running app.")
	app.Router.Run(":3000")
}
