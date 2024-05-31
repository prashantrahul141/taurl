package main

import "taurl/app"

var GlobalConfig app.Configuration = app.InitConfig()

func main() {
	app := app.Default(GlobalConfig)
	app.Run(GlobalConfig.Port)
}
