package Examples

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func GoroutinesInMiddleware(c *gin.Context) {
	cCp := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("Done! in path", cCp.Request.URL.Path)

	}()
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Done! in path %s", cCp.Request.URL.Path)})
}
