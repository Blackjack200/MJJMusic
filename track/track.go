package track

import (
	"encoding/json"
	"github.com/blackjack200/mjjmusic/util"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
)

type AlbumPeriod string

/*const (
	OffTheWall           = AlbumPeriod("Off The Wall")
	Thriller             = AlbumPeriod("Thriller")
	Bad                  = AlbumPeriod("Bad")
	Dangerous            = AlbumPeriod("Dangerous")
	History              = AlbumPeriod("History")
	BloodOnTheDanceFloor = AlbumPeriod("Blood On The Dance Floor")
	Invincible           = AlbumPeriod("Invincible")
	ThisIsIt             = AlbumPeriod("This is it")
)*/

type InternalRecord struct {
	Name string
	Desc string
	Year int
	Path string
}

type PublicRecord struct {
	Name  string
	Desc  string
	Year  int
	Index string
}

var mux = sync.Mutex{}
var songs = make(map[string]InternalRecord)

func Write(path string) ([]byte, error) {
	mux.Lock()
	defer mux.Unlock()
	return json.Marshal(songs)
}

func Load(path string) error {
	mux.Lock()
	defer mux.Unlock()
	songs = make(map[string]InternalRecord)
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, f := range dir {
		if !f.IsDir() {
			jsonPath := f.Name()
			if !strings.Contains(jsonPath, ".json") {
				continue
			}
			jsonPath = filepath.Join(path, jsonPath)
			if b, err := ioutil.ReadFile(jsonPath); err != nil {
				return err
			} else {
				var r InternalRecord
				if err := json.Unmarshal(b, &r); err != nil {
					return err
				}
				r.Path = filepath.Join(path, r.Name)
				songs[util.Hash256(r.Name)] = r
			}
		}
	}
	return nil
}

func GetAll() []PublicRecord {
	mux.Lock()
	defer mux.Unlock()
	var r []PublicRecord
	for _, v := range songs {
		r = append(r, toPublic(v))
	}
	return r
}

func Get(hash string) (InternalRecord, bool) {
	mux.Lock()
	defer mux.Unlock()
	r, ok := songs[hash]
	return r, ok
}

func toPublic(record InternalRecord) PublicRecord {
	return PublicRecord{
		Name:  record.Name,
		Desc:  record.Desc,
		Year:  record.Year,
		Index: util.Hash256(record.Name),
	}
}

func toInternal(file string, record PublicRecord) InternalRecord {
	return InternalRecord{
		Name: record.Name,
		Desc: record.Desc,
		Year: record.Year,
		Path: file,
	}
}

func Register(basename string, file string, record PublicRecord) {
	mux.Lock()
	songs[util.Hash256(basename)] = toInternal(file, record)
	mux.Unlock()
}
