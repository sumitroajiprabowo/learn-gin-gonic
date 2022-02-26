package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BasicExample is a function to show basic example
func BasicExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": "Success",
		"data":   "Hello World",
	})

}

// Create person struct with fields name, age, and married
type Person struct {
	Name    string `json:"name" form:"name" required:"true" uri:"name" binding:"required"`
	Age     int    `json:"age" form:"age" required:"true" uri:"age" binding:"required"`
	Married bool   `json:"married" form:"married" required:"true"`
}

// BasicExampleWithStruct is a function to show basic example with struct
func BasicExampleWithStruct(c *gin.Context) {
	person := Person{Name: "Danu Budi Raharjo", Age: 22, Married: false}
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": "Success",
		"data":   person,
	})
}

// BasicExampleWithMap is a function to show basic example with map
func BasicExampleWithMap(c *gin.Context) {
	person := map[string]interface{}{
		"name":    "Otot Slotter",
		"age":     24,
		"married": false,
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": "Success",
		"data":   person,
	})
}

// BasicExampleAsciiJson is a function to show basic example with ascii json
func BasicExampleAsciiJson(c *gin.Context) {
	data := map[string]interface{}{
		"name":    "Budi",
		"email":   "budi@gmail.com",
		"age":     24,
		"balance": 10000,
	}
	c.AsciiJSON(http.StatusOK, data)
}
