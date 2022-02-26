package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/sumitroajiprabowo/learn-gin-gonic/v1/example"
)

// Create Test Basic Example for testing the basic example
func TestBasicExample(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)

	// Create a request to send to the server
	router := gin.Default()

	// Add the route.
	router.GET("/api/basic/example", example.BasicExample)

	//request to the router
	request, _ := http.NewRequest("GET", "/api/basic/example", nil)

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	router.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	// //check response body
	expected := `{"code":200,"data":"Hello World","status":"Success"}`

	//check status code
	assert.Equal(t, http.StatusOK, response.StatusCode)

	//check response body
	assert.Equal(t, expected, recorder.Body.String())

}

// Create TestBasicExampleWithStruct for testing the basic example with struct
func TestBasicExampleWithStruct(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)

	// Create a request to send to the server
	router := gin.Default()

	// Add the route.
	router.GET("/api/basic/example/struct", example.BasicExampleWithStruct)

	//request to the router
	request, _ := http.NewRequest("GET", "/api/basic/example/struct", nil)

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	router.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	// //check response body
	expected := `{"code":200,"data":{"name":"Danu Budi Raharjo","age":22,"married":false},"status":"Success"}`

	//check status code
	assert.Equal(t, http.StatusOK, response.StatusCode)

	//check response body
	assert.Equal(t, expected, recorder.Body.String())

}

// Create TestBasicExampleWithMap for testing the basic example with map
func TestBasicExampleWithMap(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)

	// Create a request to send to the server
	router := gin.Default()

	// Add the route.
	router.GET("/api/basic/example/map", example.BasicExampleWithMap)

	//request to the router
	request, _ := http.NewRequest("GET", "/api/basic/example/map", nil)

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	router.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	// //check response body
	expected := `{"code":200,"data":{"age":24,"married":false,"name":"Otot Slotter"},"status":"Success"}`

	//check status code
	assert.Equal(t, http.StatusOK, response.StatusCode)

	//check response body
	assert.Equal(t, expected, recorder.Body.String())

}

// Create TestBasicExampleAsciiJson for testing the basic example with ascii json
func TestBasicExampleAsciiJson(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)

	// Create a request to send to the server
	router := gin.Default()

	// Add the route.
	router.GET("/api/basic/example/ascii", example.BasicExampleAsciiJson)

	//request to the router
	request, _ := http.NewRequest("GET", "/api/basic/example/ascii", nil)

	//recorder
	recorder := httptest.NewRecorder()

	//handler
	router.ServeHTTP(recorder, request)

	//response
	response := recorder.Result()

	// //check response body
	expected := `{"age":24,"balance":10000,"email":"budi@gmail.com","name":"Budi"}`

	//check status code
	assert.Equal(t, http.StatusOK, response.StatusCode)

	//check response body
	assert.Equal(t, expected, recorder.Body.String())

}
