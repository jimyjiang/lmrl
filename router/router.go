package router

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"

	"lmrl/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed all:templates/*
var templateFS embed.FS

//go:embed all:frontend/*
var embeddedFrontend embed.FS

func Init(r *gin.Engine) {
	templ := template.Must(template.New("").ParseFS(templateFS, "templates/*.tpl"))
	r.SetHTMLTemplate(templ)

	r.GET("/lmrl/", api.LMRL)
	assetsFS, err := fs.Sub(embeddedFrontend, "frontend/assets")
	fmt.Printf("err %v \n", err)
	r.StaticFS("/lmrl/assets", http.FS(assetsFS))

	apigroup := r.Group("/lmrl/api")
	{
		apigroup.GET("/sermons", api.ListSermon)
		apigroup.GET("/search", api.Search)
	}
	r.GET("/lmrl/search", func(c *gin.Context) {
		indexHTML, err := fs.ReadFile(embeddedFrontend, "frontend/index.html")
		if err != nil {
			c.String(500, "Failed to load index.html")
			return
		}
		fmt.Printf("%s", indexHTML)
		c.Data(200, "text/html", indexHTML)
	})
}
