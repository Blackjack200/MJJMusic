package util

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
)

func Hash256(str string) string {
	h := sha256.New()
	_, err := io.WriteString(h, str)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
