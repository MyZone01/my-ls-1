package main

import (
	"fmt"
	"my-ls-1/lib/utils"
	"os"
)

func main() {
	args := os.Args[1:]

	// files := fsys.List(f, fsys.NormalOutput)
	utils.RecursivePrint(args[0])
	// utils.PrintFiles(files)
	fmt.Println()
}
