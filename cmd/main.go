package main

import (
	"echo_rest_timer/internal/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	cfg, err := config.Load()
	if err != nil {
		e.Logger.Fatal("Config load error", err)
	}

	e.Logger.SetLevel(cfg.Environment.LogLevel())

	port := ":" + cfg.Port
	e.GET("/", func(c echo.Context) error {
		e.Logger.Debug("Use e.GET /,func(c echo.Context)")
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Debug(cfg.Date.Day())
	e.Logger.Fatal(e.Start(port))

}
