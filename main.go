package main

import (
	_ "embed"
	"github.com/blackjack200/mjjmusic/track"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

//go:embed html/index.html
var index []byte

//go:embed html/list.html
var list []byte

//go:embed html/about.html
var about []byte

func main() {
	if wd, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		path := filepath.Join(wd, "music")
		if err := os.MkdirAll(path, 0777); err != nil {
			panic(err)
		}
		if err := track.Load(path); err != nil {
			panic(err)
		}
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		_, _ = c.Writer.Write(index)
	})
	r.GET("/list", func(c *gin.Context) {
		_, _ = c.Writer.Write(list)
	})
	r.GET("/about", func(c *gin.Context) {
		_, _ = c.Writer.Write(about)
	})
	r.GET("/details/:index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "implement me",
		})
	})
	r.GET("/obtain_list", func(c *gin.Context) {
		c.JSON(http.StatusOK, track.GetAll())
	})
	r.GET("/download/:index", func(c *gin.Context) {
		record, found := track.Get(c.Param("index"))
		if found {
			c.Header("Content-Description", "File Transfer")
			c.Header("Content-Transfer-Encoding", "binary")
			c.Header("Content-Disposition", "attachment; filename="+record.Name+filepath.Ext(record.Path))
			c.Header("Content-Type", "application/octet-stream")
			c.File(record.Path)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		}
	})
	r.GET("/play/:index", func(c *gin.Context) {
		record, found := track.Get(c.Param("index"))
		if found {
			c.File(record.Path)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		}
	})
	err := r.Run(":80")
	if err != nil {
		return
	}
}
