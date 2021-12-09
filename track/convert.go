package track

import "github.com/blackjack200/mjjmusic/util"

func toPublic(record InternalRecord) PublicRecord {
	return PublicRecord{
		Name:  record.Name,
		Desc:  record.Desc,
		Year:  record.Year,
		Index: util.Identifier(record.Name),
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
