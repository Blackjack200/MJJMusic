package service

import (
	"net/http"

	"github.com/blackjack200/mjjmusic/track"
	"github.com/gin-gonic/gin"
)

type ListService struct {
	DefaultService
	Tracks *track.Set
}

func (d *ListService) Register(e *gin.Engine) {
	e.GET("/obtain_list", func(c *gin.Context) {
		c.JSON(http.StatusOK, d.Tracks.GetPublic())
	})
	e.GET("/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", nil)
	})
}

type IndexService struct {
	DefaultService
}

func (i *IndexService) Register(e *gin.Engine) {
	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}

type DetailsService struct {
	DefaultService
	Tracks *track.Set
}

func (i *DetailsService) Register(e *gin.Engine) {
	e.GET("/details/", func(c *gin.Context) {
		record, found := i.Tracks.Internal(c.Query("index"))
		if found {
			c.HTML(http.StatusOK, "details.tmpl", record)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})
}

type AboutService struct {
	DefaultService
}

func (i *AboutService) Register(e *gin.Engine) {
	e.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})
}

type AudioService struct {
	DefaultService
	Tracks *track.Set
}

func (i *AudioService) Register(e *gin.Engine) {
	e.GET("/download/:index", func(c *gin.Context) {
		record, found := i.Tracks.Internal(c.Param("index"))
		if found {
			c.Header("Content-Description", "File Transfer")
			c.Header("Content-Transfer-Encoding", "binary")
			c.Header("Content-Disposition", "attachment; filename="+record.FileName)
			c.Header("Content-Type", "application/octet-stream")
			c.File(record.FilePath)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})
	e.GET("/direct_play/:index", func(c *gin.Context) {
		record, found := i.Tracks.Internal(c.Param("index"))
		if found {
			c.File(record.FilePath)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})
}
