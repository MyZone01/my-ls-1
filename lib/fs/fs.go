package fs

import (
	"fmt"
	"io/fs"
	"my-ls-1/lib/utils"

	"strings"
)

func List(fsys fs.FS) {
	res, _ := fs.Glob(fsys, "*")
	utils.OrderFiles(res, strings.Compare)

	PrintFiles(res)
}

func PrintFiles(files []string) {
	// width := sys.GetTerminalWidth()
	// size := utils.GetOutputLength(files)

	// if width-size > 0 {
	for i, file := range files {
		if file[0] != '.' {
			fmt.Print(file)
			if i < len(files)-1 {
				fmt.Print("  ")
			}
		}
	}
	// }
}
