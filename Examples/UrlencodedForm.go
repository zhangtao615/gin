package Examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UrlencodedForm(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(http.StatusOK, gin.H{"status": "posted", "message": message, "nick": nick})
}
