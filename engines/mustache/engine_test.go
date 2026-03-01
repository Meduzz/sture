package mustache_test

import (
	"bytes"
	"testing"

	"github.com/Meduzz/sture"
	"github.com/Meduzz/sture/engines/mustache"
	"github.com/Meduzz/sture/model"
	orig "github.com/cbroglie/mustache"
)

func TestMustache(t *testing.T) {
	t.Run("Unhappy cases", func(t *testing.T) {
		t.Run("bad template", func(t *testing.T) {
			target := bytes.NewBufferString("")
			err := sture.Render(tpl(badTemplate), data, target)

			if err == nil {
				t.Error("expected render to throw an error")
			}

			if target.String() != "" {
				t.Errorf("did not expect a result, but was: %s", target.String())
			}
		})

		t.Run("missing partial", func(t *testing.T) {
			target := bytes.NewBufferString("")
			err := sture.Render(tpl(partialTemplate), data, target)

			if err != nil {
				t.Errorf("did not expect an error but was: %v", err)
			}

			if target.String() != "Hello !" {
				t.Errorf("did not expect a result, but was: %s", target.String())
			}
		})
	})

	t.Run("Happy cases", func(t *testing.T) {
		t.Run("good template", func(t *testing.T) {
			target := bytes.NewBufferString("")
			err := sture.Render(tpl(goodTemplate), data, target)

			if err != nil {
				t.Errorf("did not expect an error but was: %v", err)
			}

			if target.String() != "Hello world!" {
				t.Errorf("result was not expected but was: %s", target.String())
			}
		})

		t.Run("present partial", func(t *testing.T) {
			target := bytes.NewBufferString("")
			mustache.SetPartialProvider(&orig.StaticProvider{Partials: partial})
			err := sture.Render(tpl(partialTemplate), data, target)

			if err != nil {
				t.Errorf("did not expect an error but was: %v", err)
			}

			if target.String() != "Hello world!" {
				t.Errorf("result was not expected but was: %s", target.String())
			}
		})
	})
}

func tpl(it string) *model.Template {
	return &model.Template{
		Kind:    model.Mustache,
		Content: it,
	}
}

var (
	badTemplate     = `Hello {{world}!`
	goodTemplate    = `Hello {{world}}!`
	partialTemplate = `Hello {{>child}}!`
	data            = map[string]any{
		"world": "world",
	}
	partial = map[string]string{
		"child": "{{world}}",
	}
)
