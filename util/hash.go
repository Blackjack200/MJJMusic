package util

import (
	"strconv"
	"sync"
	"sync/atomic"
)

var mu = sync.Mutex{}
var cnt int64 = 0
var m = make(map[string]string)

func Identifier(str string) string {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := m[str]; !ok {
		m[str] = "song_" + strconv.FormatInt(atomic.AddInt64(&cnt, 1), 16)
	}
	return m[str]
}
