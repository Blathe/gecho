package handlers

import (
	"github.com/blathe/gecho/db"
	"github.com/blathe/gecho/models"
)

func IndexHandler(db *db.TodoDatabase) map[string][]models.Todo {
	todos := map[string][]models.Todo{
		"Todos": db.Todos,
	}
	return todos
}
