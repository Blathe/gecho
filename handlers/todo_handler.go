package handlers

import (
	"net/http"

	"github.com/blathe/gecho/models"
	"github.com/labstack/echo"
)

func HandleLoadTodos(c echo.Context) error {
	data := map[string][]models.Todo{
		"Todos": {
			{Title: "Test", Complete: false},
			{Title: "Walk the dog", Complete: true},
		},
	}
	return c.Render(http.StatusOK, "todo_items", data)
}
