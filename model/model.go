package model

import "io"

type (
	Kind string

	Template struct {
		Kind    Kind
		Content string
	}

	Engine interface {
		Render(*Template, any, io.Writer) error
	}
)

var (
	MD       = Kind("md")
	Mustache = Kind("mustache")
	Gotmpl   = Kind("gotmpl")
)
