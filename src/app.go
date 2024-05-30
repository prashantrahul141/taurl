package src

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	Db     DbManager
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
	// public user interface endpoints.
	app.Router.GET("/", app.Index)

	// endpoint for actual redirection.
	app.Router.GET("/:id", app.Redirect)

	// api routes
	api := app.Router.Group("/api")
	{
		api.GET("/get", app.ApiGetUrl)
		api.GET("/get_from_id", app.ApiGetUrlFromId)
		api.POST("/set", app.ApiSetUrl)
	}
}

// creates and inits a new app.
func Default() App {
	app := App{Router: gin.Default(), Db: SetupDb()}
	app.Init()
	app.MountRoutes()
	return app
}

// starts all services.
func (app *App) Run() {
	slog.Info("Running app.")
	app.Router.Run(":3000")
}
