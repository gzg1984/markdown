package markdown_test

import (
	"io/ioutil"
	"testing"

	"github.com/gzg1984/markdown"
)

func TestCode(t *testing.T) {
	code, _ := ioutil.ReadFile("markdown_test.go")
	book := markdown.NewMarkDown()
	book.WriteMultiCode(string(code), "go")
	err := book.Export("code.md")
	if err != nil {
		t.Log(err)
	}

}
