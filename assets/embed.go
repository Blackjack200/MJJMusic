package assets

import (
	"embed"
	_ "embed"
	"io/fs"

	"github.com/blackjack200/mjjmusic/util"
)

//go:embed html
var html embed.FS
var avaFs fs.FS

//go:embed favicon.ico
var favicon []byte

//go:embed config_default.json
var defaultConfig []byte

func init() {
	var err error
	avaFs, err = fs.Sub(html, "html")
	util.Must(err)
}

func DefaultConfig() []byte {
	return defaultConfig
}

func Favicon() []byte {
	return favicon
}
