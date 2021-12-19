package assets

import (
	"github.com/blackjack200/mjjmusic/util"
	"html/template"
	"io/fs"
	"net/http"
)

func ParseHTMLTemplate() *template.Template {
	return template.Must(template.New("").ParseFS(embedHTML, "html/*"))
}

func StaticJavaScriptFS() http.FileSystem {
	f, err := fs.Sub(embedJs, "html/js")
	util.Must(err)
	return http.FS(f)
}
