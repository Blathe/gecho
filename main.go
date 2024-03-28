package main

import (
	"html/template"
	"io"

	"github.com/blathe/gecho/handlers"
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

	t := template.Must(template.ParseGlob("views/*.html"))

	renderer := &TemplateRenderer{
		templates: t,
	}
	e.Renderer = renderer

	e.GET("/", handlers.IndexHandler)
	e.GET("/todos", handlers.HandleLoadTodos)

	e.Logger.Fatal(e.Start(":8080"))
}