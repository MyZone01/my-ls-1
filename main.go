package main

import (
	"fmt"
	"io/fs"
	my_ls "my_ls/lib"
	"os"
	"strings"
)

func main() {
	// Parse command line flags
	showAll, longFormat, recursive, reverse, sortByTime, dirPath := my_ls.GetFlags()

	// Perform the directory listing
	if recursive {
		f := os.DirFS(dirPath)
		fs.WalkDir(f, ".", func(path string, entry fs.DirEntry, err error) error {
			if entry.IsDir() {
				if !strings.HasPrefix(entry.Name(), ".") || (strings.HasPrefix(entry.Name(), ".") && (showAll || len(entry.Name()) == 1)) {
					if entry.Name() != "." {
						fmt.Println("./" + dirPath + "/" + path + ":")
					} else {
						fmt.Println(dirPath + ":")
					}
					my_ls.ListFiles(dirPath + "/" + path, showAll, longFormat, recursive, reverse, sortByTime)
					fmt.Println()
				} else {
					err = fs.SkipDir
				}
			}
			return err
		})
		fmt.Print("\033[1A")
	} else {
		my_ls.ListFiles(dirPath, showAll, longFormat, recursive, reverse, sortByTime)
	}
}
