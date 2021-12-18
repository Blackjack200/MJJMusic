package util

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type LogrusInfoWriter struct {
	Logger *logrus.Logger
}

func (l *LogrusInfoWriter) Write(p []byte) (n int, err error) {
	l.Logger.Info(strings.TrimLeft(string(p), "[GIN] "))
	return len(p), nil
}
