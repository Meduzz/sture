package gotmpl_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Meduzz/sture"
	"github.com/Meduzz/sture/engines/gotmpl"
	"github.com/Meduzz/sture/model"
)

func TestGotempl(t *testing.T) {
	t.Run("Unhappy cases", func(t *testing.T) {
		t.Run("bad template", func(t *testing.T) {
			target := bytes.NewBufferString("")
			err := sture.Render(tpl(badTemplate), "world", target)

			if err == nil {
				t.Error("expected an error from bad template")
			}

			if target.String() != "" {
				t.Errorf("did not expct any result, but was: %s", target.String())
			}
		})

		t.Run("missing func", func(t *testing.T) {
			target := bytes.NewBufferString("")
			err := sture.Render(tpl(funcTemplate), "world", target)

			if err == nil {
				t.Error("expected an error from missing func")
			}

			if target.String() != "" {
				t.Errorf("did not expect any result, but was: %s", target.String())
			}
		})
	})

	t.Run("Happy cases", func(t *testing.T) {
		t.Run("good template", func(t *testing.T) {
			target := bytes.NewBufferString("")
			err := sture.Render(tpl(goodTemplate), "world", target)

			if err != nil {
				t.Errorf("render threw error: %v", err)
			}

			if target.String() != "Hello world!" {
				t.Errorf("result was wrong, was: %s", target.String())
			}
		})

		t.Run("good func", func(t *testing.T) {
			gotmpl.SetFunc("upper", strings.ToUpper)
			target := bytes.NewBufferString("")
			err := sture.Render(tpl(funcTemplate), "world", target)

			if err != nil {
				t.Errorf("render threw error: %v", err)
			}

			if target.String() != "Hello WORLD!" {
				t.Errorf("result was wrong, was: %s", target.String())
			}
		})
	})
}

func tpl(data string) *model.Template {
	return &model.Template{
		Kind:    model.Gotmpl,
		Content: data,
	}
}

var (
	badTemplate  = `Hello {{.}!`
	goodTemplate = `Hello {{.}}!`
	funcTemplate = `Hello {{upper .}}!`
)
