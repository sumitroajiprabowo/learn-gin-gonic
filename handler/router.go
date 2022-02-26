package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sumitroajiprabowo/learn-gin-gonic/example"
)

// SetupRouter is a function to setup router
func SetupRouter() *gin.Engine {
	r := gin.Default() // Default gin router

	r.LoadHTMLGlob("templates/*") // Load html templates

	r.GET("/api/basic/example", example.BasicExample)
	r.GET("/api/basic/example/struct", example.BasicExampleWithStruct)
	r.GET("/api/basic/example/map", example.BasicExampleWithMap)
	r.GET("/api/basic/example/ascii", example.BasicExampleAsciiJson)
	r.GET("/api/basic/bindrequest/teacher", example.GetBindRequestNestedStruct)
	r.GET("/api/basic/bindrequest/student", example.GetBindRequestNestedStructPointer)
	r.GET("/api/basic/bindrequest/employee", example.GetBindRequestNestedAnonyStruct)
	r.GET("/api/basic/bindrequestformdata", example.IndexFormHandler)
	r.POST("/api/basic/bindrequestformdata", example.FormHandler)
	r.GET("/api/bindquerystring", example.BindQueryString)
	r.GET("/api/binduri/:name/:age", example.BindURI)

	return r
}
