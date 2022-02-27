package example

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func FormHandler(c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	fmt.Println(fakeForm)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}

func IndexFormHandler(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}

func CreateArticlePostForm(c *gin.Context) {
	title := c.PostForm("title")
	desc := c.PostForm("desc")

	c.JSON(200, gin.H{
		"code":   200,
		"status": "success",
		"data":   []string{title, desc},
	})
}
