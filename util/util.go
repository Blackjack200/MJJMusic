package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func MakeIndex(str string) string {
	h := sha256.New()
	Must(io.WriteString(h, str))
	return strings.ReplaceAll(base64.StdEncoding.EncodeToString(h.Sum(nil)), "/", "_")
}
func HexString(str []byte) string {
	return fmt.Sprintf("%x", str)
}

func RandomBytes(len int) []byte {
	b := make([]byte, len)
	_, err := rand.Read(b)
	Must(err)
	return b
}

func Hash256(str string) string {
	h := sha256.New()
	Must(io.WriteString(h, str))
	return HexString(h.Sum(nil))
}

func Must(params ...interface{}) {
	for _, param := range params {
		if v, ok := param.(error); ok && v != nil {
			panic(v.Error())
		}
	}
}

func MustString(params ...interface{}) string {
	for _, param := range params {
		if v, ok := param.(string); ok && v != "" {
			return v
		}
	}
	panic("Non-Empty String Result Not Found")
}

func Error(params ...interface{}) error {
	for _, param := range params {
		if v, ok := param.(error); ok && v != nil {
			return v
		}
	}
	return nil
}

func IsHiddenPath(path string) bool {
	return strings.HasPrefix(path, ".")
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

func MimeType(file string) (string, error) {
	cmd := exec.Command("file", "-b", "--mime-type", file)
	cmd.Env = os.Environ()

	if output, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("failed to read mime type: %v", err)
	} else {
		return strings.TrimSpace(string(output)), nil
	}
}

func MimeEncoding(file string) (string, error) {
	cmd := exec.Command("file", "-b", "--mime-encoding", file)
	cmd.Env = os.Environ()

	if output, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("failed to read mime encoding: %v", err)
	} else {
		return strings.TrimSpace(string(output)), nil
	}
}

func FileExists(file string) bool {
	if info, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	} else {
		return !info.IsDir()
	}
	return false
}

func FileSize(file string) (int64, error) {
	if stat, err := os.Stat(file); err != nil {
		return -1, fmt.Errorf("failed to get file stat: %v", err)
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

func WriteFile(file string, d []byte) error {
	return ioutil.WriteFile(file, d, 0644)
}

func DeleteFile(file string) error {
	return os.Remove(file)
}

func ScanDir(dir string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to scan dir: %v", err)
	}
	return files, nil
}

func TempFile(fileName string, data []byte) (string, error) {
	tmpFile, err := ioutil.TempFile("", fileName)
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %v", err)
	}
	if _, err := tmpFile.Write(data); err != nil {
		return "", fmt.Errorf("failed to write temp file: %v", err)
	}
	if err := tmpFile.Close(); err != nil {
		return "", fmt.Errorf("failed to close temp file: %v", err)
	}
	return tmpFile.Name(), nil
}
