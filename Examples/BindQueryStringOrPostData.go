package Examples

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Age      int       `form:"age"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func BindQueryStringOrPostData(c *gin.Context) {
	var person Person
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name, person.Age, person.Birthday)
	}
	c.String(http.StatusOK, "Success")
}
