package models

import (
	"database/sql"
	"fmt"
	"sync"
)

type Todo struct {
	Id       int
	Title    string
	Complete bool
}

type Todos struct {
	mu sync.Mutex
	db *sql.DB
}

// Creates a new todo list with the passed database connection.
func CreateTodoList(db *sql.DB) *Todos {
	list := Todos{
		mu: sync.Mutex{},
		db: db,
	}
	return &list
}

// Gets all todos and returns a slice of todos.
func (t *Todos) GetAllTodos() ([]Todo, error) {
	var todos []Todo
	rows, err := t.db.Query("SELECT * FROM todos ORDER BY id;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		todo := Todo{}
		err := rows.Scan(&todo.Id, &todo.Complete, &todo.Title)
		if err != nil {
			fmt.Println("Error scanning row...")
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (t *Todos) GetByID(id int) (Todo, error) {
	row := t.db.QueryRow("SELECT * FROM todos WHERE id=?", id)

	todo := Todo{}
	var err error
	if err = row.Scan(&todo.Id, &todo.Complete, &todo.Title); err == sql.ErrNoRows {
		fmt.Println(err)
		return Todo{}, err
	}
	return todo, nil
}

func (t *Todos) UpdateComplete(id int) error {
	_, err := t.db.Exec("UPDATE todos SET complete = NOT(complete) WHERE id=?;", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Insert a new Todo into the database.
func (t *Todos) Insert(todo Todo) (int, error) {
	res, err := t.db.Exec("INSERT INTO todos VALUES(NULL, ?, ?);", todo.Complete, todo.Title)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}

	return int(id), nil
}

// Deletes a Todo from the database.
func (t *Todos) Delete(todo_id int) error {
	res, err := t.db.Exec("DELETE FROM todos WHERE id=?;", todo_id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var amt int64
	if amt, err = res.RowsAffected(); err != nil {
		fmt.Printf("Rows deleted: %v\n", amt)
		return err
	}

	return nil
}
