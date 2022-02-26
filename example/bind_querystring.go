package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
If `GET`, only `Form` binding engine (`query`) used.
If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
*/

func BindQueryString(c *gin.Context) {

	var p Person
	if c.ShouldBindQuery(&p) == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"status": "Success",
			"data":   p,
		})
	}
}
