package track

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/blackjack200/mjjmusic/util"
)

func makeJsonManifest(basePath string, jsonManifestFile string) (*Manifest, error) {
	manifest := &Manifest{}
	b, err := util.ReadFile(jsonManifestFile)
	if err != nil {
		return nil, fmt.Errorf("error read file: %v", err)
	}
	if err := json.Unmarshal(b, manifest); err != nil {
		return nil, fmt.Errorf("error read json: %v", err)
	}
	if !util.FileExists(filepath.Join(basePath, manifest.FileName)) {
		return nil, fmt.Errorf("error file not exists: %v", manifest.FileName)
	}
	return manifest, nil
}

func makeInternalRecord(basePath string, jsonManifestFile string) (*InternalRecord, error) {
	if manifest, err := makeJsonManifest(basePath, jsonManifestFile); err != nil {
		return nil, err
	} else {
		audioPath := filepath.Join(basePath, manifest.FileName)
		if size, err := util.FileSize(audioPath); err != nil {
			return nil, fmt.Errorf("error get file size: %v", err)
		} else if info, err := util.FileInfo(audioPath); err != nil {
			return nil, fmt.Errorf("error file info: %v", err)
		} else {
			return &InternalRecord{
				Manifest:      *manifest,
				FilePath:      audioPath,
				FileName:      manifest.FileName,
				FileSize:      util.HumanReadableFileSize(size),
				FileInfo:      info,
				InternalIndex: util.MakeIndex(manifest.Name),
			}, nil
		}
	}
}
