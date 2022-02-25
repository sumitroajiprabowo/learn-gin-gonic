package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Married bool `json:"married"`
}

func BasicExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "Success",
		"data":    "Hello World",
	})

}
