package service

import (
	"github.com/blackjack200/mjjmusic/assets"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Service interface {
	Register(e *gin.Engine)
	SetLogger(l *logrus.Logger)
	Logger() *logrus.Logger
}

type DefaultService struct {
	log *logrus.Logger
}

func (d *DefaultService) Register(*gin.Engine, assets.HTMLRender) {

}

func (d *DefaultService) SetLogger(l *logrus.Logger) {
	d.log = l
}

func (d *DefaultService) Logger() *logrus.Logger {
	return d.log
}
