package service

import (
	"github.com/blackjack200/mjjmusic/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

var runtimeTokenPassword = util.RandomBytes(1024)

type loginToken struct {
	ClientIP  string
	UserAgent string
	jwt.StandardClaims
}

func newToken(c *gin.Context) string {
	tk := &loginToken{}
	tk.ClientIP = c.ClientIP()
	tk.UserAgent = c.Request.UserAgent()
	tk.ExpiresAt = time.Now().Add(time.Minute * 30).Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString(runtimeTokenPassword)
	return tokenString
}

func tokenValid(c *gin.Context, token string) bool {
	if token == "" {
		return false
	}
	tk := &loginToken{}
	if t, err := jwt.ParseWithClaims(token, tk, func(token *jwt.Token) (interface{}, error) {
		return runtimeTokenPassword, nil
	}); err != nil {
		return false
	} else if t.Valid &&
		strings.EqualFold(tk.ClientIP, c.ClientIP()) &&
		strings.EqualFold(tk.UserAgent, c.Request.UserAgent()) {
		return true
	}
	return false
}
