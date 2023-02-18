package markdown_test

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/gzg1984/markdown"
)

/*
	Run It by

go test -v -run TestMain
*/
func TestMain(t *testing.T) {
	example()
}

func example() {
	code, _ := ioutil.ReadFile("markdown_test.go")
	book := markdown.NewMarkDown()
	book.WriteTitle("Go-MarkDownDoc-Generator", markdown.LevelTitle).
		WriteLines(2)

	book.WriteMultiCode(string(code), "go")

	book.WriteTitle("Author", markdown.LevelNormal).
		WriteCodeLine("lichun")

	book.WriteTitle("Website", markdown.LevelNormal)
	book.WriteLinkLine("lichunorz", "https://lichunorz.com")

	t := markdown.NewTable(4, 4)
	t.SetTitle(0, "Version")
	t.SetTitle(1, "Date")
	t.SetTitle(2, "Creator")
	t.SetTitle(3, "Remarks")
	t.SetContent(0, 0, "v1")
	t.SetContent(0, 1, "2019-11-21")
	t.SetContent(0, 2, "Lee")
	t.SetContent(0, 3, "无")
	book.WriteTable(t)
	err := book.Export("README.md")
	if err != nil {
		log.Fatal(err)
	}
}
