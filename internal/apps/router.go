package apps

import (
	"go-boilerplate/config"
	"log"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo, cfg *config.Config) {
	_, err := cfg.ConnectionMySql()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	} else {
		log.Println("Successfully connected to the database")
	}

	e.Group("/api")

}
