package apps

import (
	"go-boilerplate/config"
	mw "go-boilerplate/internal/adapter/middleware"
	"go-boilerplate/internal/core/service"
	"go-boilerplate/utils/validator"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter(e *echo.Echo, cfg *config.Config) {
	_, err := cfg.ConnectionMySql()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	} else {
		log.Println("Successfully connected to the database")
	}

	customValidator := validator.NewValidator()
	en.RegisterDefaultTranslations(customValidator.Validator, customValidator.Translator)
	e.Validator = customValidator

	jwtService := service.NewJwtService(cfg)

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Group("/api")

	e.GET("/api/check", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	mid := mw.NewMiddlewareAdapter(cfg, jwtService)
	adminGroup := e.Group("/admin", mid.CheckToken())
	adminGroup.GET("/check", func(c echo.Context) error {
		return c.String(http.StatusOK, "admin access granted")
	})

}
