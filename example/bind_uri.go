package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindUri(c *gin.Context) {

	var p Person
	if err := c.ShouldBindUri(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": "Error",
			"data":   err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"status": "Success",
			"data":   p,
		})
	}
}
