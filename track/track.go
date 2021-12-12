package track

import (
	"fmt"
	"github.com/blackjack200/mjjmusic/util"
	"path/filepath"
	"strings"
	"sync"
)

type AlbumPeriod string

type Manifest struct {
	Name     string `json:"Name"`
	Desc     string `json:"Desc"`
	Year     int    `json:"Year"`
	FileName string `json:"Path"`
}

type InternalRecord struct {
	Manifest      Manifest
	FilePath      string
	FileName      string
	FileSize      string
	FileInfo      string
	InternalIndex string
}

type PublicRecord struct {
	Name  string
	Desc  string
	Year  int
	Index string
}

var mux = sync.Mutex{}
var publicRecords []PublicRecord
var songs = make(map[string]InternalRecord)

func Load(basePath string) error {
	mux.Lock()
	defer mux.Unlock()
	songs = make(map[string]InternalRecord)

	if dir, err := util.ScanDir(basePath); err != nil {
		return fmt.Errorf("error scandir: %v", err)
	} else {
		for _, f := range dir {
			manifestJsonName := f.Name()
			if f.IsDir() || !strings.EqualFold(filepath.Ext(manifestJsonName), ".json") {
				continue
			}
			manifestJsonPath := filepath.Join(basePath, manifestJsonName)
			if r, err := makeInternalRecord(basePath, manifestJsonPath); err != nil {
				return fmt.Errorf("error make internal record: %v", err)
			} else {
				songs[r.InternalIndex] = *r
			}
		}

		rcd := make([]PublicRecord, 0, len(songs))
		for _, v := range songs {
			rcd = append(rcd, toPublic(v))
		}
		publicRecords = sortPublic(rcd)

		return nil
	}
}

func GetPublic() []PublicRecord {
	return publicRecords
}

func GetInternal(hash string) (InternalRecord, bool) {
	r, ok := songs[hash]
	return r, ok
}
