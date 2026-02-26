package main

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"ok": "server is on"})
}

func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("api/v1/healthcheck", healthCheck)

	if err := e.Start(":8080"); err != nil {
		slog.Error("failed to start server", "error", err)
	}

}
