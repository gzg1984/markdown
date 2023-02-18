package markdown_test

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gzg1984/markdown"
)

type vidioCollation struct {
	name     string
	year     string
	episodes map[string] /*season*/ string /*episode */
}

var vcm = make(map[string]vidioCollation)

// -args selpg -s1 -e1 input.txt
func onlyRecordJap(path string, info os.FileInfo, err error) error {
	if err != nil {
		return nil /*ignore error*/
	}

	if info.IsDir() == true {
		current_deepth := len(strings.Split(path, "/"))
		if current_deepth == original_deepth+1 {
			vcm[info.Name()] = vidioCollation{
				name: info.Name(),
			}
		}
		return nil
	} else {
		/* ignore  files and  go on*/
		return nil
	}

}

var global_analyze_path = "/Volumes/2T Vedio/日本动画"
var original_deepth = len(strings.Split(global_analyze_path, "/"))

func TestListFile(t *testing.T) {
	t.Logf("Args are: %v", os.Args)
	_ = filepath.Walk(global_analyze_path, onlyRecordJap)

	book := markdown.NewMarkDown()
	book.WriteTitle("2T Vedio/日本动画", markdown.LevelTitle).
		WriteLines(2)

	vidiaTable := markdown.NewTable(len(vcm), 2)
	i := 0
	vidiaTable.SetTitle(0, "名字")
	for name := range vcm {
		vidiaTable.SetContent(i, 0, name)
		//t.Logf("contest are: %v", c)
		i++
	}
	book.WriteTable(vidiaTable)
	err := book.Export("video.md")
	if err != nil {
		log.Fatal(err)
	}

	s, _ := json.Marshal(vcm)
	/*TODO: 名字和初始路径关联*/
	os.WriteFile("2.json", s, os.ModePerm)

}

/*
	Run It by

go test -v -run TestMain
*/
func TestMain(t *testing.T) {
	example()
}

func example() {
	book := markdown.NewMarkDown()
	book.WriteTitle("Go-MarkDownDoc-Generator", markdown.LevelTitle).
		WriteLines(2)

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
