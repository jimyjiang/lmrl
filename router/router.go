package router

import (
	"embed"
	"html/template"
	"lmrl/api"

	"github.com/gin-gonic/gin"
)

//go:embed all:templates/*
var templateFS embed.FS

func Init(r *gin.Engine) {
	templ := template.Must(template.New("").ParseFS(templateFS, "templates/*.tpl"))
	r.SetHTMLTemplate(templ)
	r.GET("/", api.LMRL)
}
