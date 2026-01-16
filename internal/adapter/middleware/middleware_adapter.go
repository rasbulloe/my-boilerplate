package middleware

import (
	"go-boilerplate/config"
	"go-boilerplate/internal/adapter/handler/response"
	"go-boilerplate/internal/core/service"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type MiddlewareAdapterInterface interface {
	// Define methods for middleware adapter
	CheckToken() echo.MiddlewareFunc
}

type middlewareAdapter struct {
	// Add any dependencies required by the middleware adapter here
	cfg        *config.Config
	jwtService service.JwtServiceInterface
}

func (m *middlewareAdapter) CheckToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Implement token checking logic here
			// For example, extract token from Authorization header and validate it
			respErr := response.DefaultResponse{}
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				respErr.Message = "Missing Authorization header"
				respErr.Data = nil
				return c.JSON(http.StatusUnauthorized, respErr)
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			token, err := m.jwtService.ValidateToken(tokenString)
			if err != nil {
				log.Printf("Middleware[1] CheckToken: %s", err.Error())
				respErr.Message = "Invalid token"
				respErr.Data = nil
				return c.JSON(http.StatusUnauthorized, respErr)
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				respErr.Message = "Invalid token claims"
				respErr.Data = nil
				return c.JSON(http.StatusUnauthorized, respErr)
			}

			userID := int64(claims["user_id"].(float64))

			// Token is valid, proceed to the next handler
			c.Set("token", tokenString)
			c.Set("user_id", userID)
			return next(c)
		}
	}
}

func NewMiddlewareAdapter(cfg *config.Config, jwtService service.JwtServiceInterface) MiddlewareAdapterInterface {
	return &middlewareAdapter{
		cfg:        cfg,
		jwtService: jwtService,
	}
}
