package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sumitroajiprabowo/learn-gin-gonic/example"
)

// SetupRouter is a function to setup router
func SetupRouter() *gin.Engine {
	r := gin.Default() // Default gin router
	r.GET("/api/basic/example", example.BasicExample)
	r.GET("/api/basic/example/struct", example.BasicExampleWithStruct)
	r.GET("/api/basic/example/map", example.BasicExampleWithMap)
	r.GET("/api/basic/example/ascii", example.BasicExampleAsciiJson)
	return r
}

// RunServer is a function to run server
func main() {
	router := SetupRouter() // Setup router
	router.Run(":3000") // Run on port 3000
}
