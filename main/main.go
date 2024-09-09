package main

import (
	"gin-basic/Examples"
	"gin-basic/MiddleWares"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	gin.DefaultWriter = os.Stdout
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	// logger middleware
	r.Use(MiddleWares.Logger())

	// load html template
	r.LoadHTMLGlob("web/*")
	r.GET("/", indexHandler)

	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// ascii-json
	r.GET("/someJson", Examples.AsciiJSON)

	// form-data custom struct
	r.GET("/getb", Examples.GetDataB)
	r.GET("/getc", Examples.GetDataC)
	r.GET("/getd", Examples.GetDataD)

	// Bind html checkboxes
	r.POST("/bindForm", Examples.FormHandler)

	// Bind query string or post data
	r.GET("/bindQueryString", Examples.BindQueryStringOrPostData)

	// Bind Uri
	r.GET("/bindUri/:name/:id", Examples.BindUri)

	// validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(Examples.UserStructLevelValidation, Examples.User{})
	}
	r.POST("/user", Examples.Validators)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}
