package utils

import (
	"fmt"
	"io/fs"
	"my-ls-1/lib/fsys"
	"my-ls-1/lib/helper"
	"my-ls-1/lib/sys"
	"os"
	"strings"
)

func PrintFiles(files []string) {
	fmt.Print("")
	width := sys.GetTerminalWidth()
	size := helper.GetOutputLength(files)
	var curColAt = 1
	var curLinAt = 1
	var nCol, temp int

	if width-size >= 0 {
		for _, file := range files {
			fmt.Printf("\033[%dG%v", curColAt, file)
			curColAt += len(file) + 2
		}
	} else {
		maxCols := helper.GetMaxCols(width, files)

		l := (len(files) + maxCols) / maxCols
		nline := ""
		var lastf string
		for i, file := range files {
			if strings.Contains(file, " ") {
				file = fmt.Sprintf("'%v'", file)
			}
			if i < l {
				fmt.Printf("%v", file)
				if i < l-1 {
					fmt.Println()
				}
			} else {
				if i%l == 0 {
					nline = fmt.Sprintf("\033[%vA", l-1)
					curColAt = temp + nCol
					nCol += temp
					temp = len(file) + 2
					lastf = ""
					curLinAt = 1
				} else {
					nline = "\033[1B"
				}
				fmt.Printf("%s\033[%dG%v", nline, curColAt+1, file)
			}
			if len(file)+2 > temp && len(file) > len(lastf) {
				temp = len(file) + 2

			}
			lastf = file
			curLinAt++
		}
		fmt.Printf("\033[%vB", l-curLinAt+1)
	}
}

func RecursivePrint(dirPath string) {
	f := os.DirFS(dirPath)
	ndir := 0
	var fileNumInCurrentDir int
	var upCursor = 2
	fs.WalkDir(f, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			subFS, _ := fs.Sub(f, path)
			files := fsys.List(subFS, fsys.NormalOutput)
			fileNumInCurrentDir = len(files)
			if d.Name() == "." {
				path =""
			}else {
				path="/"+path
			}
			fmt.Printf("\033[0G%v%v:\n", dirPath,path)
			PrintFiles(files)
			fmt.Println()
			fmt.Println()
			ndir++
		}
		return err
	})
	if fileNumInCurrentDir ==0 {
		upCursor++
	}
	if ndir != 0 {
		fmt.Printf("\033[%vA",upCursor)
	}

}
