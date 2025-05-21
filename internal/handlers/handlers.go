package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func NewHandlerStatus(date time.Time) echo.HandlerFunc {
	return func(c echo.Context) error {
		now := time.Now()
		timeLess := date.Sub(now)
		if timeLess < 0 {
			timeLess = 0
		}
		return c.String(http.StatusOK, timeLess.String())
	}
}
