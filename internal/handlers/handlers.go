package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func NewHandlerStatus(date time.Time) echo.HandlerFunc {
	return func(c echo.Context) error {
		timeLess := time.Until(date)
		if timeLess < 0 {
			reqStr := fmt.Sprintf("Количество дней осталось до %v: %d", date, 0)
			return c.String(http.StatusOK, reqStr)
		}
		reqStr := fmt.Sprintf("Количество дней осталось до %v: %d", date, int(timeLess.Hours())/24)
		return c.String(http.StatusOK, reqStr)
	}
}
