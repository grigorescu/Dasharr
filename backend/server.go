package main

import (
	"backend/handlers"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	apiKey := os.Getenv("API_KEY")

	apiKeyMiddleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			providedKey := c.Request().Header.Get("X-API-Key")
			if providedKey != apiKey {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid API key")
			}
			return next(c)
		}
	}

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	e.GET("/initdb", handlers.InitDB, apiKeyMiddleware)
	e.GET("/update", handlers.Update, apiKeyMiddleware)
	e.GET("/stats", handlers.GetStats, apiKeyMiddleware)
	e.GET("/config", handlers.GetConfig, apiKeyMiddleware)
	e.POST("/saveCredentials", handlers.SaveCredentials, apiKeyMiddleware)
	e.GET("/savedCredentials", handlers.SavedCredentials, apiKeyMiddleware)
	e.GET("/prowlarrConfig", handlers.GetProwlarrIndexerIds, apiKeyMiddleware)
	e.Logger.Fatal(e.Start(":1323"))
}
