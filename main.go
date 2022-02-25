package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sumitroajiprabowo/learn-gin-gonic/example"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/api/basic/example", example.BasicExample)
	r.GET("/api/basic/example/struct", example.BasicExampleWithStruct)
	r.GET("/api/basic/example/map", example.BasicExampleWithMap)
	r.GET("/api/basic/example/ascii", example.BasicExampleAsciiJson)
	return r
}

func main() {
	router := SetupRouter()
	router.Run(":3000")
}
