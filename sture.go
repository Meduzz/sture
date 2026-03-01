package sture

import (
	"fmt"
	"io"

	"github.com/Meduzz/sture/model"
	"github.com/Meduzz/sture/registry"
)

func Render(template *model.Template, data any, to io.Writer) error {
	engine, ok := registry.EngineFor(template.Kind)

	if !ok {
		return fmt.Errorf("no engine registered for %s", template.Kind)
	}

	return engine.Render(template, data, to)
}

func Register(kind model.Kind, eng model.Engine) {
	registry.Register(kind, eng)
}
