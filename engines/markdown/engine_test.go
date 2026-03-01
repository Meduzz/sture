package markdown_test

import (
	"bytes"
	"testing"

	"github.com/Meduzz/sture"
	_ "github.com/Meduzz/sture/engines/markdown"
	"github.com/Meduzz/sture/model"
)

func TestMarkdown(t *testing.T) {
	md := `# Hello world!`
	target := bytes.NewBufferString("")
	err := sture.Render(tpl(md), nil, target)

	if err != nil {
		t.Errorf("did not expect error from render but was: %v", err)
	}

	if target.String() != "<h1>Hello world!</h1>" {
		t.Errorf("target was not the expected value, but was: %s", target.String())
	}
}

func tpl(it string) *model.Template {
	return &model.Template{
		Kind:    model.MD,
		Content: it,
	}
}
