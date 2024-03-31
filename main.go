package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/blathe/gecho/models"
	"github.com/blathe/gecho/utils"
	"github.com/labstack/echo"

	_ "github.com/mattn/go-sqlite3"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	//Create a new instance of the echo framework
	e := echo.New()
	e.Static("static", "static")

	//Connect to our sqlite database
	db, err := sql.Open("sqlite3", "todos.db")
	if err != nil {
		fmt.Println(err)
	}

	//Create our todo list with our DB connection.
	todo_list := models.CreateTodoList(db)

	//Define the location of our templates
	t := template.Must(template.ParseGlob("views/*.html"))
	renderer := &TemplateRenderer{
		templates: t,
	}
	e.Renderer = renderer

	//Routes
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})
	e.GET("/todos", func(c echo.Context) error {
		todos, err := todo_list.GetAllTodos()
		if err != nil {
			return c.NoContent(400)
		}

		data := map[string][]models.Todo{
			"Todos": todos,
		}
		return c.Render(http.StatusOK, "todo_items", &data)
	})
	e.POST("/todos", func(c echo.Context) error {
		if c.FormValue("todo-title") == "" {
			return c.NoContent(400)
		}

		new_todo := models.Todo{
			Id:       0,
			Title:    c.FormValue("todo-title"),
			Complete: false,
		}

		id, err := todo_list.Insert(new_todo)
		if err != nil {
			return c.NoContent(400)
		}

		new_todo.Id = id
		return c.Render(http.StatusOK, "todo_item", new_todo)
	})
	e.PUT("/todos/:id/complete", func(c echo.Context) error {
		id, err := utils.StringToInt(c.Param("id"))
		if err != nil {
			fmt.Println("error with ID")
			return c.NoContent(400)
		}
		err = todo_list.ToggleComplete(id)
		if err != nil {
			fmt.Println(err)
			return c.NoContent(400)
		}

		return c.NoContent(200)
	})
	e.DELETE("/todos/:id", func(c echo.Context) error {
		id, err := utils.StringToInt(c.Param("id"))
		fmt.Println(id)
		if err != nil {
			fmt.Println(err)
			return c.NoContent(400)
		}

		err = todo_list.Delete(id)
		if err != nil {
			fmt.Println(err)
			return c.NoContent(400)
		}

		return c.NoContent(200)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
