package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Children struct {
	Name string `uri:"name" binding:"required" json:"name"`
	Age  int    `uri:"age" binding:"required" json:"age"`
}

// BindURI is a handler function that bind uri params to struct
func BindURI(c *gin.Context) {
	var ch Children //initialize struct
	if err := c.ShouldBindUri(&ch); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"status": "Success",
			"data":   ch,
		})
	}
}
