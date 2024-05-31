package app

import (
	"fmt"
	"log/slog"
	"os"
)

// App Configurations.
type Configuration struct {
	BaseUrl string
	Port    string
}

func (c Configuration) String() string {
	return fmt.Sprintf("BASE_URL=%s", c.BaseUrl)
}

func InitConfig() Configuration {
	base_url, ok := os.LookupEnv("BASE_URL")
	if !ok {
		slog.Warn("BASE_URL env not found, defaulting to http://localhost:3000/")
		base_url = "http://localhost:3000/"
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		slog.Warn("PORT env not found, defaulting to 3000")
		port = "3000"
	}

	c := Configuration{BaseUrl: base_url, Port: port}
	slog.Info("config:")
	slog.Info(c.String())
	return c
}
