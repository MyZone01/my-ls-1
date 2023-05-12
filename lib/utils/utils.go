package utils

import (
	"fmt"
	"io/fs"
	"my-ls-1/lib/helper"
	"my-ls-1/lib/fsys"
	"my-ls-1/lib/sys"
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
		lin := ""
		// var lastf string
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
					lin = fmt.Sprintf("\033[%vA", l-1)
					curColAt = temp + nCol
					nCol += temp
					temp = len(file)+2
					curLinAt = 1
				} else {
					lin = "\033[1B"
				}
				fmt.Printf("%s\033[%dG%v", lin, curColAt+1, file)
			}
			if len(file) > temp {
				temp = len(file) + 2

			}
			// lastf = file
			curLinAt++
		}
		fmt.Printf("\033[%vB", l-curLinAt+1)
	}
}

func RecursivePrint(f fs.FS) {
	ndir := 0
	fs.WalkDir(f, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			subFS, _ := fs.Sub(f, path)
			files := fsys.List(subFS,fsys.NormalOutput)
			fmt.Printf("\033[0G%v:\n", d.Name())
			PrintFiles(files)
			fmt.Print("\033[2B")
		}

		return err
	})
	if ndir != 0 {
		fmt.Print("\033[2A")	
	}

}
