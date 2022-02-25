package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "Success",
		"data":    "Hello World",
	})

}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Married bool `json:"married"`
}

func BasicExampleWithStruct(c *gin.Context){
	person := Person{Name: "Danu Budi Raharjo", Age: 22, Married: false}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "Success",
		"data":    person,
	})
}

func BasicExampleWithMap(c *gin.Context){
	person := map[string]interface{}{
		"name": "Otot Slotter",
		"age": 24,
		"married": false,
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "Success",
		"data":    person,
	})
}

func BasicExampleAsciiJson(c *gin.Context){
	data := map[string]interface{}{
		"name": "Budi",
		"email": "budi@gmail.com",
		"age": 24,
		"balance": 10000,
	}
	c.AsciiJSON(http.StatusOK, data)
}
