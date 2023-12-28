package godskeqing_test

import (
	"testing"

	godskeqing "github.com/XuanLee-HEALER/gods-keqing"
)

func TestDirTree(t *testing.T) {
	tree := godskeqing.NewDirTree()
	dir1 := godskeqing.NewDir("A")
	dir2 := godskeqing.NewDir("B")
	file1 := godskeqing.NewFile(&godskeqing.SimpleFile{FileName: "B1"})
	file2 := godskeqing.NewFile(&godskeqing.SimpleFile{FileName: "B2"})
	dir2.AddFile(file1)
	dir2.AddFile(file2)
	file5 := godskeqing.NewFile(&godskeqing.SimpleFile{FileName: "C"})
	dir4 := godskeqing.NewDir("D")
	dir5 := godskeqing.NewDir("D1")
	dir6 := godskeqing.NewDir("D2")
	file3 := godskeqing.NewFile(&godskeqing.SimpleFile{FileName: "E1"})
	file4 := godskeqing.NewFile(&godskeqing.SimpleFile{FileName: "E2"})
	dir5.AddFile(file3)
	dir5.AddFile(file4)
	dir4.AddDir(dir5)
	dir4.AddDir(dir6)
	tree.AddDir(dir1)
	tree.AddDir(dir2, "A")
	tree.AddFile(file5)
	tree.AddDir(dir4)
	tree.AddFile(godskeqing.NewFile(&godskeqing.SimpleFile{FileName: "B3"}), "A", "B")
	t.Logf("\n%s\n", tree)
}

func TestDirToTree(t *testing.T) {
	p := `F:\yzy\keen`
	tr, err := godskeqing.ReadDirTree(p)
	if err != nil {
		t.Error(err)
	}

	t.Logf("\n%s\n", tr)
}
