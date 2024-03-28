package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/blathe/gecho/db"
	"github.com/blathe/gecho/handlers"
	"github.com/blathe/gecho/models"
	"github.com/blathe/gecho/utils"
	"github.com/labstack/echo"
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
	e := echo.New()
	e.Static("static", "static")

	db := db.TodoDatabase{
		Todos: []models.Todo{
			{Id: 1, Title: "Learn Go", Complete: false},
			{Id: 2, Title: "Learn HTMX", Complete: false},
			{Id: 3, Title: "Learn HTML/CSS", Complete: true},
		},
	}

	t := template.Must(template.ParseGlob("views/*.html"))

	renderer := &TemplateRenderer{
		templates: t,
	}
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		data := handlers.HandleLoadTodos(&db)
		return c.Render(http.StatusOK, "index", data)
	})

	e.GET("/todos", func(c echo.Context) error {
		data := handlers.HandleLoadTodos(&db)
		return c.Render(http.StatusOK, "todo_item", data)
	})

	e.POST("/add-todo", func(c echo.Context) error {
		new_todo := models.Todo{
			Title:    c.FormValue("todo-title"),
			Complete: false,
		}
		handlers.HandleAddTodo(&db, &new_todo)
		return c.Render(http.StatusOK, "todo_item", new_todo)
	})

	e.POST("/toggle-todo", func(c echo.Context) error {
		current_status := utils.ToBool(c.FormValue("currentStatus"))
		id := utils.ToInt(c.FormValue("id"))

		handlers.HandleToggleTodo(&db, id, current_status)
		return c.Render(http.StatusOK, "todo_item", nil)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
