package fs

import (
	"fmt"
	"io/fs"
	"my-ls-1/lib/utils"
	"strings"
)

func List(fsys fs.FS) {
	res, _ := fs.Glob(fsys, "*")
	utils.OrderFiles(res,strings.Compare)
	for i, v := range res {

		hiden := v[0] == '.'
		if !hiden {
			fmt.Print(v)
			if i < len(res)-1 {
				fmt.Print("  ")
			}
		}
	}
}
