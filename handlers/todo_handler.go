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

func HandleToggleTodo(todos *db.TodoDatabase, id int) *models.Todo {
	for i := range todos.Todos {
		if todos.Todos[i].Id == id {
			todos.Todos[i].Complete = !todos.Todos[i].Complete
			return &todos.Todos[i]
		}
	}
	return nil
}

func HandleDeleteTodo(todos *models.Todos, todo_id int) (int, error) {
	err := todos.Delete(todo_id)
	if err != nil {
		return 0, err
	}

	return todo_id, nil
}
