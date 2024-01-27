package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Contact struct {
	Name  string
	Email string
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func newData() Data {
	return Data{
		Contacts{
			Contact{
				Name:  "hello",
				Email: "email",
			},
			Contact{
				Name:  "hello1",
				Email: "email1",
			},
		},
	}
}

type Count struct {
	Count int
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	data := newData()
	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")

		data.Contacts = append(data.Contacts, Contact{Name: name, Email: email})

		return c.Render(200, "index", data)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
