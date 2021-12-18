package assets

import (
	"html/template"

	"github.com/gin-gonic/gin/render"
)

type HTMLRender struct {
	templates map[string]*template.Template
}

func NewRender() *HTMLRender {
	r := &HTMLRender{
		templates: make(map[string]*template.Template),
	}
	register := func(f string, d string) {
		r.templates[f] = template.Must(template.New(f).Parse(d))
	}
	register("index.html", index)
	register("list.html", list)
	register("details.tmpl", details)
	register("about.html", about)
	register("login.html", login)
	register("panel.tmpl", panel)
	return r
}

func (r *HTMLRender) Instance(s string, i interface{}) render.Render {
	return render.HTML{
		Template: r.templates[s],
		Data:     i,
	}
}
