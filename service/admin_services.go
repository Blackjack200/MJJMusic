package service

import (
	"net/http"
	"runtime"
	"strings"

	"github.com/blackjack200/mjjmusic/util"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
	DefaultService
	Account  string
	Password string
}

func (i *AuthService) Register(e *gin.Engine) {
	e.POST("/auth/req", func(c *gin.Context) {
		ac := c.PostForm("account")
		pd := c.PostForm("password")
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
		tokenString := newToken(c)
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "login success",
			"token":   tokenString,
		})
	})
	e.POST("/auth/test", func(c *gin.Context) {
		tk := c.PostForm("token")
		stat := "failed"
		if tokenValid(c, tk) {
			stat = "ok"
		}
		c.JSON(http.StatusOK, gin.H{
			"status": stat,
		})
	})
}

type RuntimeInfo struct {
	GoRoutineNum int
	AllocMem     uint64
	VirtualMem   uint64
	StackMem     uint64
	HeapMem      uint64
	GCCycles     uint32
}

func getMemStats() runtime.MemStats {
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)
	return m2
}

func newRuntimeInfo() RuntimeInfo {
	info := RuntimeInfo{}
	stat := getMemStats()
	info.GoRoutineNum = runtime.NumGoroutine()
	info.AllocMem = stat.Sys / 1024 / 1024
	info.VirtualMem = stat.HeapSys / 1024 / 1024
	info.StackMem = stat.StackSys / 1024 / 1024
	info.HeapMem = (stat.Mallocs - stat.Frees) / 1024 / 1024
	info.GCCycles = stat.NumGC
	return info
}

type AdminService struct {
	DefaultService
	Entrance string
}

func (i *AdminService) Register(e *gin.Engine) {
	e.GET("/"+i.Entrance, func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	e.POST("/panel", func(c *gin.Context) {
		requestAuthorized(c, func() {
			c.HTML(http.StatusOK, "panel.tmpl", nil)
		})
	})
}

type ManipulateService struct {
	DefaultService
	Entrance string
}

func (i *ManipulateService) Register(e *gin.Engine) {
	e.POST("/manipulate/gc", func(c *gin.Context) {
		requestAuthorized(c, func() {
			runtime.GC()
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		})
	})
	e.POST("/manipulate/mem", func(c *gin.Context) {
		requestAuthorized(c, func() {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"info":   newRuntimeInfo(),
			})
		})
	})
}

func requestAuthorized(c *gin.Context, then func()) {
	tk := c.PostForm("token")
	if tokenValid(c, tk) {
		then()
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "unauthorized"})
	}
}
