package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/sumitroajiprabowo/learn-gin-gonic/example"
)

// Create test example.go

func TestBasicExample(t *testing.T) {

	// Create a request to send to the server
	router := gin.Default()

	// Add the route.
	router.GET("/api/basic", example.BasicExample)

	//request to the router
	request, _ := http.NewRequest("GET", "/api/basic", nil)

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
