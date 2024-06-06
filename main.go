package main

import (
	"traininggolang/author"
	"traininggolang/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi router Gin
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(author.AuthMiddleware())

	router.SetupRouter(r)

	// Jalankan server pada port 8080
	r.Run(":8080")
}
