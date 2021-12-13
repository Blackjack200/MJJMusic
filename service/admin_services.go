package service

import (
	"github.com/blackjack200/mjjmusic/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthService struct {
	DefaultService
	Account  string
	Password string
}

func (i *AuthService) Register(e *gin.Engine) {
	e.GET("/auth/req", func(c *gin.Context) {
		ac := c.Query("account")
		pd := c.Query("password")
		if len(ac) == 0 || len(pd) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": "username or password is empty",
				"token":   nil,
			})
			return
		}
		if !strings.EqualFold(util.Hash256(i.Account), ac) ||
			!strings.EqualFold(util.Hash256(i.Password), pd) {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": "username or password is incorrect",
				"token":   nil,
			})
			return
		}
		tokenString := newToken()
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "login success",
			"token":   tokenString,
		})
	})
	e.GET("/auth/test", func(c *gin.Context) {
		tk := c.Query("token")
		c.JSON(http.StatusOK, gin.H{
			"status": tokenValid(tk),
		})
	})
}

type AdminService struct {
	DefaultService
	Entrance string
}

func (i *AdminService) Register(e *gin.Engine) {
	e.GET("/"+i.Entrance, func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	e.GET("/panel", func(c *gin.Context) {
		tk := c.Query("token")
		if tokenValid(tk) {
			//TODO Implement Admin Panel
			c.HTML(http.StatusOK, "panel.html", nil)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		}
	})
}
