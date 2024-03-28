package handlers

import (
	"github.com/blathe/gecho/db"
	"github.com/blathe/gecho/models"
)

func HandleLoadTodos(todos *db.TodoDatabase) map[string][]models.Todo {
	data := map[string][]models.Todo{
		"Todos": todos.Todos,
	}
	return data
}

func HandleAddTodo(todos *db.TodoDatabase, newTodo *models.Todo) {
	todos.Todos = append(todos.Todos, *newTodo)
}

func HandleToggleTodo(todos *db.TodoDatabase, id int, newStatus bool) {
	for _, v := range todos.Todos {
		if v.Id == id {
			v.Complete = newStatus
		}
	}
}
