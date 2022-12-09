package render

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Template struct {
	Templates map[string]*template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates[name].ExecuteTemplate(w, name, data)
}
