package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sumitroajiprabowo/learn-gin-gonic/handler"
)

// RunServer is a function to run server
func main() {
	gin.SetMode(gin.ReleaseMode)
	router := handler.SetupRouter() // Setup router
	router.Run(":3000")             // Run on port 3000
}
