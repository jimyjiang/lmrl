package router

import (
	"lmrl/api"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.LoadHTMLGlob("templates/*")
	r.GET("/", api.LMRL)
}
