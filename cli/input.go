package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/blackjack200/mjjmusic/track"
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
		if _, err := track.Load(path); err != nil {
			panic(err)
		}
		for {
			r := input()
			f := filepath.Join(path, r.Name+".json")
			if len(r.FileName) == 0 {
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

func input() track.Manifest {
	r := track.Manifest{}
	print("track name: ")
	r.Name = readLine()
	print("track description: ")
	r.Desc = readLine()
	suc := false
	for !suc {
		print("track year: ")
		if y, err := strconv.Atoi(readLine()); err != nil {
			print("\r")
		} else {
			r.Year = y
			suc = true
		}
	}
	print("music file: ")
	r.FileName = readLine()
	return r
}
