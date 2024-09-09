package Examples

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func UserStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(User)
	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		sl.ReportError(user.FirstName, "FirstName", "first_name", "fnameorlname", "")
		sl.ReportError(user.LastName, "LastName", "last_name", "fnameorlname", "")
	}
	if len(user.Email) == 0 {
		sl.ReportError(user.Email, "Email", "email", "email", "")
	}
}

func Validators(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "User validation successful."})
	}
}
