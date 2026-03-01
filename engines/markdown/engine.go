package markdown

import (
	"io"

	"github.com/Meduzz/sture"
	"github.com/Meduzz/sture/model"
	md "github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

type (
	engine struct {
		md md.Markdown
	}
)

var (
	_        model.Engine = &engine{}
	markdown              = &engine{
		md: md.New(
			md.WithExtensions(extension.GFM),
		),
	}
)

func init() {
	sture.Register(model.MD, markdown)
}

func (e *engine) Render(tpl *model.Template, data any, to io.Writer) error {
	return e.md.Convert([]byte(tpl.Content), to)
}

func Set(it md.Markdown) {
	if it == nil {
		return
	}

	markdown.md = it
}
