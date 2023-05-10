package main

import (
	"fmt"
	"my-ls-1/lib/fs"
	"my-ls-1/lib/utils"
	"os"
)

func main() {
	args := os.Args[1:]

	fsys := os.DirFS(args[0])

	files := fs.List(fsys,fs.NormalOutput)

	utils.PrintFiles(files)
	fmt.Println()
}
