package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func IndexHandler(c echo.Context) error {
	data := map[string]interface{}{
		"name": "Gecho",
	}
	return c.Render(http.StatusOK, "index.html", data)
}
