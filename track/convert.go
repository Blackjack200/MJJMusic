package track

import (
	"github.com/blackjack200/mjjmusic/util"
	"sort"
)

func toPublic(record InternalRecord) PublicRecord {
	return PublicRecord{
		Name:  record.Manifest.Name,
		Desc:  record.Manifest.Desc,
		Year:  record.Manifest.Year,
		Index: record.InternalIndex,
	}
}

func toInternal(file string, record PublicRecord) (*InternalRecord, error) {
	info, err := util.FileInfo(file)
	if err != nil {
		return nil, err
	}
	return &InternalRecord{
		Manifest: Manifest{
			Name: record.Name,
			Desc: record.Desc,
			Year: record.Year,
		},
		InternalIndex: record.Index,
		FileName:      file,
		FileInfo:      info,
	}, nil
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
