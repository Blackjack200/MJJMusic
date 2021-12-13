package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/blackjack200/mjjmusic/assets"
	"github.com/blackjack200/mjjmusic/service"
	"github.com/blackjack200/mjjmusic/track"
	"github.com/blackjack200/mjjmusic/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

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
		if err := util.WriteFile(path, assets.DefaultConfig()); err != nil {
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
		set, err := track.Load(path)
		if err != nil {
			log.Fatal(err)
		}
		gin.ForceConsoleColor()
		r := gin.New()
		r.HTMLRender = assets.NewRender()
		writer := gin.LoggerWithWriter(&util.LogrusInfoWriter{Logger: log})
		r.Use(writer, gin.Recovery(), util.NewFavicon(assets.Favicon()))
		if err := initServices(r, cfg, set); err != nil {
			log.Fatalf("error init service: %v", err)
		}
		log.Infof("hash: %s", util.Hash256(cfg.AdminAccount))
		log.Infof("hash: %s", util.Hash256(cfg.AdminPassword))
		log.Infof("Running on %v", cfg.Bind)
		util.Must(r.Run(cfg.Bind))
	}
}

func initServices(r *gin.Engine, cfg *Config, set *track.Set) error {
	reg := func(s service.Service) {
		s.SetLogger(log)
		s.Register(r)
	}
	reg(&service.IndexService{})
	reg(&service.AboutService{})
	reg(&service.ListService{Tracks: set})
	reg(&service.AudioService{Tracks: set})
	reg(&service.DetailsService{Tracks: set})
	reg(&service.AdminService{
		Entrance: cfg.AdminEntrance,
	})
	reg(&service.AuthService{
		Account:  cfg.AdminAccount,
		Password: cfg.AdminPassword,
	})
	return nil
}
