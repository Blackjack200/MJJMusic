package main

import (
	_ "embed"
	"github.com/blackjack200/mjjmusic/track"
	"github.com/blackjack200/mjjmusic/util"
	"github.com/gin-gonic/gin"
	"html/template"
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

//go:embed html/details.tmpl
var details []byte

func main() {
	detailsTmpl, parseErr := template.New("Details").Parse(string(details))
	if parseErr != nil {
		panic(parseErr)
	}
	if wd, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		path := filepath.Join(wd, "music")
		util.Must(os.MkdirAll(path, 0777))
		util.Must(track.Load(path))
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

	r.GET("/obtain_list", func(c *gin.Context) {
		c.JSON(http.StatusOK, track.GetPublic())
	})

	r.GET("/download/:index", func(c *gin.Context) {
		record, found := track.GetInternal(c.Param("index"))
		if found {
			c.Header("Content-Description", "File Transfer")
			c.Header("Content-Transfer-Encoding", "binary")
			c.Header("Content-Disposition", "attachment; filename="+record.FileName)
			c.Header("Content-Type", "application/octet-stream")
			c.File(record.FilePath)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		}
	})

	r.GET("/direct_play/:index", func(c *gin.Context) {
		record, found := track.GetInternal(c.Param("index"))
		if found {
			c.File(record.FilePath)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		}
	})
	r.GET("/details/:index", func(c *gin.Context) {
		record, found := track.GetInternal(c.Param("index"))
		if found {
			util.Must(detailsTmpl.Execute(c.Writer, record))
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		}
	})

	util.Must(r.Run(":80"))
}
