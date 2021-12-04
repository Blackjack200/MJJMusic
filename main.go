package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
)

//go:embed html/index.html
var index []byte

func main() {
	r := gin.Default()
	r.LoadHTMLFiles()
	r.GET("/", func(c *gin.Context) {
		_, _ = c.Writer.Write(index)
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run(":80")
	if err != nil {
		return
	}
}
