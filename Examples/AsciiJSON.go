package Examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Using AsciiJSON to Generates ASCII-only JSON with escaped non-ASCII characters.
func AsciiJSON(c *gin.Context) {
	data := map[string]interface{}{"lang": "GO语言", "tag": "</br>"}
	c.AsciiJSON(http.StatusOK, data)
}
