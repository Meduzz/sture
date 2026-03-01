package mustache

import (
	"io"

	"github.com/Meduzz/sture"
	"github.com/Meduzz/sture/model"
	. "github.com/cbroglie/mustache"
)

type (
	engine struct{}
)

var (
	_        model.Engine = &engine{}
	provider PartialProvider
)

func init() {
	sture.Register(model.Mustache, &engine{})
}

func (e *engine) Render(tpl *model.Template, data any, to io.Writer) error {
	var musch *Template
	var err error

	if provider != nil {
		musch, err = ParseStringPartials(tpl.Content, provider)
	} else {
		musch, err = ParseString(tpl.Content)
	}

	if err != nil {
		return err
	}

	return musch.FRender(to, data)
}

func SetPartialProvider(it PartialProvider) {
	provider = it
}
