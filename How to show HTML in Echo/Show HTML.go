package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

// TemplateRenderer é uma estrutura que implementa a interface echo.Renderer
// para renderizar templates HTML.
type TemplateRenderer struct {
	templates *template.Template
}

// Render renderiza o template HTML usando o pacote html/template e envia como resposta.
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// Configuração do mecanismo de renderização de templates
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	// Rota para renderizar o HTML
	e.GET("/", func(c echo.Context) error {
		data := map[string]interface{}{
			"title": "My Page",
			// Outros dados que você deseja fornecer ao template
		}
		return c.Render(http.StatusOK, "index.html", data)
	})

	e.Start(":8080")
}
