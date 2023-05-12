package fsys

import (
	"io/fs"
	"my-ls-1/lib/helper"

	"strings"
)

func List(fsys fs.FS, f func(files []string) []string) []string {
	res, _ := fs.Glob(fsys, "*")
	helper.OrderFiles(res, strings.Compare)
	return f(res)
}

func NormalOutput(files []string) []string {
	var temp []string
	for _, v := range files {
		if v[0] != '.' {
			temp = append(temp, v)
		}
	}
	return temp
}
