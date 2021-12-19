package assets

import (
	"embed"
	_ "embed"
)

//go:embed favicon.ico
var favicon []byte

//go:embed config_default.json
var defaultConfig []byte

//go:embed html/js
var embedJs embed.FS

//go:embed html/*.*
var embedHTML embed.FS

func DefaultConfig() []byte {
	return defaultConfig
}

func Favicon() []byte {
	return favicon
}
