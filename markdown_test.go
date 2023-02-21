package markdown_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		fmt.Printf("onlyRecordJap meet error %v", err)
		return nil /*ignore error*/
	}

	if info.IsDir() == true {
		current_deepth := len(strings.Split(path, "/"))
		fmt.Printf("current_deepth is %d\n", current_deepth)
		fmt.Println(path) //打印当前文件或目录下的文件或目录名

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

// var original_deepth = len(strings.Split(global_analyze_path, "/"))

// 只分析到第四层目录
var original_deepth = 4

var result_file_name_prefix string

/*
自动获取外接硬盘下的目标目录的路径
*/
func getTargetPath() string {
	//获取文件或目录相关信息
	fileInfoList, err := ioutil.ReadDir("/Volumes")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range fileInfoList {
		if file.Name() == "Macintosh HD" {
			continue
		} else {
			result_file_name_prefix = file.Name()
			return "/Volumes/" + file.Name() + "/日本动画"
		}

		//fmt.Println(fileInfoList[i].Name()) //打印当前文件或目录下的文件或目录名
	}
	return "/not_exist"
}

func TestListFile(t *testing.T) {

	//t.Logf("Args are: %v", os.Args)
	_ = filepath.Walk(getTargetPath(), onlyRecordJap)

	book := markdown.NewMarkDown()
	book.WriteTitle(result_file_name_prefix+"/日本动画", markdown.LevelTitle).
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
	err := book.Export(result_file_name_prefix + ".md")
	if err != nil {
		log.Fatal(err)
	}

	s, _ := json.Marshal(vcm)
	/*TODO: 名字和初始路径关联*/
	os.WriteFile(result_file_name_prefix+".json", s, os.ModePerm)

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
