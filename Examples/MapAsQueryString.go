package Examples

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MapAsQueryString(c *gin.Context) {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")
	c.JSON(http.StatusOK, gin.H{
		"ids":   ids,
		"names": names,
	})
	fmt.Printf("ids: %v; names: %v", ids, names)
}
