package service

import (
	"github.com/blackjack200/mjjmusic/util"
	"github.com/golang-jwt/jwt"
	"time"
)

var runtimeTokenPassword = util.RandomBytes(1024)

func newToken() string {
	tk := &jwt.StandardClaims{}
	tk.ExpiresAt = time.Now().Add(time.Minute * 30).Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString(runtimeTokenPassword)
	return tokenString
}

func tokenValid(token string) bool {
	if token == "" {
		return false
	}
	tk := &jwt.StandardClaims{}
	if t, err := jwt.ParseWithClaims(token, tk, func(token *jwt.Token) (interface{}, error) {
		return runtimeTokenPassword, nil
	}); err != nil {
		return false
	} else if t.Valid {
		return true
	}
	return false
}
