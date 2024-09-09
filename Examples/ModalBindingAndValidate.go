package Examples

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Login struct {
	UserName string `json:"user_name" form:"user_name" xml:"user_name" binding:"required"`
	Password string `json:"password" form:"password" xml:"password" binding:"required"`
}

func UserLoginValidation(sl validator.StructLevel) {
	var login = sl.Current().Interface().(Login)
	if len(login.UserName) == 0 || login.UserName != "admin" {
		sl.ReportError(login.UserName, "UserName", "user_name", "UserName", "")
	}
	if len(login.Password) == 0 || login.Password != "123456" {
		sl.ReportError(login.Password, "Password", "password", "Password", "")
	}
}

func LoginJson(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.UserName != "admin" || json.Password != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "you are logged in"})
}

func LoginXML(c *gin.Context) {
	var xml Login
	if err := c.ShouldBindXML(&xml); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if xml.UserName != "admin" || xml.Password != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "you are logged in"})
}

func LoginForm(c *gin.Context) {
	var form Login
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if form.UserName != "admin" || form.Password != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "you are logged in"})
}
