package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sumitroajiprabowo/learn-gin-gonic/example"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/api/basic", example.BasicExample)
	return r
}

func main() {
	router := SetupRouter()
	router.Run(":3000")
}
