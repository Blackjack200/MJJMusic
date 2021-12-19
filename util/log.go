package util

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type LogrusInfoWriter struct {
	Logger *logrus.Logger
}

func (l *LogrusInfoWriter) Write(p []byte) (n int, err error) {
	l.Logger.Info(strings.TrimSpace(strings.TrimLeft(string(p), "[GIN]")))
	return len(p), nil
}

type LogrusErrorWriter struct {
	Logger *logrus.Logger
}

func (l *LogrusErrorWriter) Write(p []byte) (n int, err error) {
	l.Logger.Error(strings.TrimSpace(strings.TrimLeft(string(p), "[GIN]")))
	return len(p), nil
}
