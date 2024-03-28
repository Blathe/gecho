package db

import "github.com/blathe/gecho/models"

// This will act as our "database"
type TodoDatabase struct {
	Todos []models.Todo
}
