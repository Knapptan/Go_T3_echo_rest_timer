package main

import (
	"echo_rest_timer/internal/config"
	"echo_rest_timer/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	cfg, err := config.Load()
	if err != nil {
		e.Logger.Fatal("Config load error", err)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.SetLevel(cfg.Environment.LogLevel())

	port := ":" + cfg.Port
	e.GET("/status", handlers.NewHandlerStatus(cfg.Date))
	e.Logger.Debug(cfg.Date.Day())
	e.Logger.Fatal(e.Start(port))

}
