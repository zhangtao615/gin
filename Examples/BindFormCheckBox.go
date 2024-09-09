package Examples

import "github.com/gin-gonic/gin"

type myForm struct {
	Colors []string `form:"colors[]"`
}

func FormHandler(c *gin.Context) {
	var form myForm
	err := c.ShouldBind(&form)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{"color": form.Colors})
}
