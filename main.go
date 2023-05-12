package main

import (
	"fmt"
	"my-ls-1/lib/utils"
	"os"
)

func main() {
	args := os.Args[1:]

	f := os.DirFS(args[0])

	// files := fsys.List(f, fsys.NormalOutput)
	utils.RecursivePrint(f)
	// utils.PrintFiles(files)
	fmt.Println()
}
