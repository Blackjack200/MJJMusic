package util

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"strings"
)

func Identifier(str string) string {
	h := sha256.New()
	_, err := io.WriteString(h, str)
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(base64.StdEncoding.EncodeToString(h.Sum(nil)), "/", "_")
}
