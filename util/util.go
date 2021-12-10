package util

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"os/exec"
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

func Must(params ...interface{}) {
	for _, param := range params {
		if v, ok := param.(error); ok && v != nil {
			panic(v.Error())
		}
	}
}

func Error(params ...interface{}) error {
	for _, param := range params {
		if v, ok := param.(error); ok && v != nil {
			return v
		}
	}
	return nil
}

func FileInfo(file string) (string, error) {
	cmd := exec.Command("file", "-b", file)
	cmd.Env = os.Environ()

	if output, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("failed to read fileinfo: %v", err)
	} else {
		return string(output), nil
	}
}
