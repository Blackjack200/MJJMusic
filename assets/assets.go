package assets

import (
	"html/template"
	"io"

	"github.com/blackjack200/mjjmusic/util"
	"github.com/gin-gonic/gin/render"
)

type HTMLRender struct {
	templates map[string]*template.Template
}

func NewRender() *HTMLRender {
	r := &HTMLRender{
		templates: make(map[string]*template.Template),
	}
	return r
}

func (r *HTMLRender) Instance(s string, i interface{}) render.Render {
	r.lazyRegister(s)
	return render.HTML{
		Template: r.templates[s],
		Data:     i,
	}
}

func (r *HTMLRender) lazyRegister(s string) {
	if _, ok := r.templates[s]; !ok {
		register := func(name string, data string) {
			r.templates[name] = template.Must(template.New(name).Parse(data))
		}
		file, err := avaFs.Open(s)
		util.Must(err)
		data, err := io.ReadAll(file)
		util.Must(err)
		register(s, string(data))
	}
}
