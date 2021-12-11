package util

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func MakeIndex(str string) string {
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

func FileExists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	} else {
		return true
	}
	return false
}

func FileSize(file string) (int64, error) {
	if stat, err := os.Stat(file); err != nil {
		return -1, err
	} else {
		return stat.Size(), nil
	}
}

func HumanReadableFileSize(size int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB"}
	unit := 0
	s := float32(size)
	for s > 1024 {
		s /= 1024
		unit++
	}
	return fmt.Sprintf("%.3f%s", s, units[unit])
}

func ReadFile(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

func ScanDir(dir string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func NewFavicon(data []byte) gin.HandlerFunc {
	reader := bytes.NewReader(data)

	return func(c *gin.Context) {
		if c.Request.RequestURI != "/favicon.ico" {
			return
		}
		if c.Request.Method != "GET" && c.Request.Method != "HEAD" {
			status := http.StatusOK
			if c.Request.Method != "OPTIONS" {
				status = http.StatusMethodNotAllowed
			}
			c.Header("Allow", "GET,HEAD,OPTIONS")
			c.AbortWithStatus(status)
			return
		}
		c.Header("Content-Type", "image/x-icon")
		http.ServeContent(c.Writer, c.Request, "favicon.ico", time.Now(), reader)
	}
}
