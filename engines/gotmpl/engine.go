package gotmpl

import (
	"html/template"
	"io"

	"github.com/Meduzz/sture"
	"github.com/Meduzz/sture/model"
)

type (
	engine struct{}
)

var (
	_     model.Engine = &engine{}
	funcs              = make(map[string]any)
)

func init() {
	sture.Register(model.Gotmpl, &engine{})
}

func (e *engine) Render(tpl *model.Template, data any, to io.Writer) error {
	temple := template.New("sture")

	if len(funcs) > 0 {
		temple = temple.Funcs(funcs)
	}

	temple, err := temple.Parse(tpl.Content)

	if err != nil {
		return err
	}

	return temple.Execute(to, data)
}

func SetFuncs(them map[string]any) {
	funcs = them
}

func SetFunc(name string, it any) {
	funcs[name] = it
}
