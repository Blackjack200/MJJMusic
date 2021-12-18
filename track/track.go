package track

import (
	"fmt"
	"github.com/blackjack200/mjjmusic/util"
	"path/filepath"
	"strings"
	"sync"
)

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

type Set struct {
	mux      sync.Mutex
	public   []PublicRecord
	internal map[string]InternalRecord
}

func (s *Set) GetPublic() []PublicRecord {
	return s.public
}

func (s *Set) Internal(hash string) (InternalRecord, bool) {
	r, ok := s.internal[hash]
	return r, ok
}

func (s *Set) Register(t InternalRecord) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.internal[t.InternalIndex] = t
	s.public = append(s.public, toPublic(t))
	s.public = sortPublic(s.public)
}

func Load(basePath string) (*Set, error) {
	s := &Set{
		mux:      sync.Mutex{},
		public:   make([]PublicRecord, 0),
		internal: make(map[string]InternalRecord),
	}

	if dir, err := util.ScanDir(basePath); err != nil {
		return nil, fmt.Errorf("error scandir: %v", err)
	} else {
		for _, f := range dir {
			manifestJsonName := f.Name()
			if f.IsDir() || util.IsHiddenPath(manifestJsonName) || !strings.EqualFold(filepath.Ext(manifestJsonName), ".json") {
				continue
			}
			manifestJsonPath := filepath.Join(basePath, manifestJsonName)
			if r, err := makeInternalRecord(basePath, manifestJsonPath); err != nil {
				return nil, fmt.Errorf("error make internal record: %v", err)
			} else {
				s.Register(*r)
			}
		}
		return s, nil
	}
}
