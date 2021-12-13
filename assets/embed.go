package assets

import _ "embed"

//go:embed favicon.ico
var favicon []byte

//go:embed html/index.html
var index string

//go:embed html/list.html
var list string

//go:embed html/about.html
var about string

//go:embed html/details.tmpl
var details string

//go:embed html/login.html
var login string

//go:embed html/panel.tmpl
var panel string

//go:embed config_default.json
var defaultConfig []byte

func DefaultConfig() []byte {
	return defaultConfig
}

func Favicon() []byte {
	return favicon
}
