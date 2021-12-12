package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/blackjack200/mjjmusic/track"
	"github.com/blackjack200/mjjmusic/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

//go:embed favicon.ico
var favicon []byte

//go:embed html/index.html
var index []byte

//go:embed html/list.html
var list []byte

//go:embed html/about.html
var about []byte

//go:embed html/details.tmpl
var details []byte

//go:embed config_default.json
var defaultConfig []byte

type Config struct {
	Tracks        string `json:"tracks"`
	Bind          string `json:"bind"`
	AdminEntrance string `json:"admin-entrance"`
	AdminAccount  string `json:"admin-account"`
	AdminPassword string `json:"admin-password"`
}

var wd string
var log = logrus.New()

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	var err error
	if wd, err = os.Getwd(); err != nil {
		log.Fatalf("failed to get working directory: %v", err)
	}
}

func loadConfig() (*Config, error) {
	path := filepath.Join(wd, "config.json")
	if !util.FileExists(path) {
		if err := util.WriteFile(path, defaultConfig); err != nil {
			return nil, fmt.Errorf("error writing default config: %v", err)
		}
	}
	if b, err := util.ReadFile(path); err != nil {
		return nil, fmt.Errorf("error reading config: %v", err)
	} else {
		cfg := &Config{}
		if err := json.Unmarshal(b, cfg); err != nil {
			return nil, fmt.Errorf("error parsing config: %v", err)
		}
		return cfg, nil
	}
}

func main() {
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	if cfg, err := loadConfig(); err != nil {
		log.Fatalf("error load config: %v", err)
	} else {
		path := filepath.Join(wd, cfg.Tracks)
		util.Must(os.MkdirAll(path, 0777))
		util.Must(track.Load(path))
		gin.ForceConsoleColor()
		r := gin.New()
		writer := gin.LoggerWithWriter(&util.LogrusInfoWriter{Logger: log})
		r.Use(writer, gin.Recovery(), util.NewFavicon(favicon))
		if err := initServices(r, cfg); err != nil {
			log.Fatalf("error init service: %v", err)
		}
		log.Infof("Running on %v", cfg.Bind)
		util.Must(r.Run(cfg.Bind))
	}
}

func initServices(r *gin.Engine, cfg *Config) error {
	detailsTmpl, parseErr := template.New("Details").Parse(string(details))
	if parseErr != nil {
		return parseErr
	}
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
			c.AbortWithStatus(http.StatusNotFound)
		}
	})

	r.GET("/direct_play/:index", func(c *gin.Context) {
		record, found := track.GetInternal(c.Param("index"))
		if found {
			c.File(record.FilePath)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})
	r.GET("/details/:index", func(c *gin.Context) {
		record, found := track.GetInternal(c.Param("index"))
		if found {
			util.Must(detailsTmpl.Execute(c.Writer, record))
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})
	return nil
}
