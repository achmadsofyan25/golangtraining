package main

import (
	"traininggolang/Pertemuan2/router"

	"github.com/gin-gonic/gin"

	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Routes
	router.SetupRouter(r)

	// Run the server
	log.Println("Running server on port 8080")
	r.Run(":8080")
}
