package main

import (
	"fmt"
	"os"
	"my-ls-1/lib/fs"
)

func main() {
	args := os.Args[1:]

	fsys := os.DirFS(args[0])

	fs.List(fsys)
	fmt.Println()
}
