package Examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type People struct {
	ID   string `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func BindUri(c *gin.Context) {
	var people People
	if err := c.ShouldBindUri(&people); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"Name": people.Name, "ID": people.ID})
}
