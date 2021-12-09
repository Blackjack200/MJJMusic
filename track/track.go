package track

import (
	"encoding/json"
	"github.com/blackjack200/mjjmusic/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
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
var publicRecords []PublicRecord
var songs = make(map[string]InternalRecord)

/*func Write(path string) ([]byte, error) {
	mux.Lock()
	defer mux.Unlock()
	return json.Marshal(songs)
}*/

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
				r.Path = filepath.Join(path, r.Path)
				if _, err := os.Stat(r.Path); os.IsNotExist(err) {
					panic(err)
				}
				songs[util.Identifier(r.Name)] = r
			}
		}
	}
	rcd := make([]PublicRecord, 0, len(songs))
	for _, v := range songs {
		rcd = append(rcd, toPublic(v))
	}
	publicRecords = sortPublic(rcd)
	return nil
}

func keys(elements map[string]PublicRecord) []string {
	i, ks := 0, make([]string, len(elements))
	for key := range elements {
		ks[i] = key
		i++
	}
	return ks
}

func sortPublic(rcd []PublicRecord) []PublicRecord {
	nameMap := make(map[string]PublicRecord)
	for _, v := range rcd {
		nameMap[v.Name] = v
	}
	k := keys(nameMap)
	sort.Strings(k)
	newMap := make([]PublicRecord, 0, len(rcd))
	for _, s := range k {
		newMap = append(newMap, nameMap[s])
	}
	return newMap
}

func GetAll() []PublicRecord {
	return publicRecords
}

func Get(hash string) (InternalRecord, bool) {
	mux.Lock()
	defer mux.Unlock()
	r, ok := songs[hash]
	return r, ok
}

/*func Register(basename string, file string, record PublicRecord) {
	mux.Lock()
	songs[util.Identifier(basename)] = toInternal(file, record)
	mux.Unlock()
}*/
