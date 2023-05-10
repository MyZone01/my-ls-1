package fs

import (
	"io/fs"
	"my-ls-1/lib/utils"

	"strings"
)

func List(fsys fs.FS, f func(files []string) []string) []string {
	res, _ := fs.Glob(fsys, "*")
	utils.OrderFiles(res, strings.Compare)
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
