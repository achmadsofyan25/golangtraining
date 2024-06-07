package router

import (
	"traininggolang/Pertemuan1/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", handler.GetHandler)

	r.POST("/post", handler.PostHandler)
}
