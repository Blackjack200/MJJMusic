package main

import (
	"bufio"
	"encoding/json"
	"github.com/blackjack200/mjjmusic/track"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func readLine() string {
	if !scanner.Scan() {
		panic("Scan failed")
	} else {
		return strings.TrimSpace(scanner.Text())
	}
}

func main() {
	if wd, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		path := filepath.Join(wd, "music")
		if err := os.MkdirAll(path, 0777); err != nil {
			panic(err)
		}
		if err := track.Load(path); err != nil {
			panic(err)
		}
		for {
			r := input()
			f := filepath.Join(path, r.Name+".json")
			if len(r.Path) == 0 {
				continue
			}
			if len(r.Name) == 0 {
				continue
			}
			b, _ := json.Marshal(r)
			if err := ioutil.WriteFile(f, b, 0777); err != nil {
				panic(err)
			}
			println("stored to : " + f)
		}
	}
}

func input() track.InternalRecord {
	r := track.InternalRecord{}
	print("track name: ")
	r.Name = readLine()
	print("track description: ")
	r.Desc = readLine()
	print("track year: ")
	suc := false
	for !suc {
		if y, err := strconv.Atoi(readLine()); err != nil {
			println("invalid year")
		} else {
			r.Year = y
			suc = true
		}
	}
	print("music file: ")
	r.Path = readLine()
	return r
}
