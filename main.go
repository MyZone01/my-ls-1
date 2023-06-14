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
	flags, allPaths := my_ls.GetFlags()

	// Perform the directory listing
	for i, dirPath := range allPaths {
		if len(allPaths) > 1 {
			if i != 0 {
				fmt.Println()
			}
			fmt.Println(dirPath + ":")
		}
		if flags.Recursive {
			f := os.DirFS(dirPath)
			fs.WalkDir(f, ".", func(path string, entry fs.DirEntry, err error) error {
				if entry.IsDir() {
					if !strings.HasPrefix(entry.Name(), ".") || (strings.HasPrefix(entry.Name(), ".") && (flags.ShowAll || len(entry.Name()) == 1)) {
						if entry.Name() != "." {
							fmt.Println(dirPath + "/" + path + ":")
						} else {
							fmt.Println(dirPath + ":")
						}
						my_ls.ListFiles(dirPath+"/"+path, flags)
						fmt.Println()
					} else {
						err = fs.SkipDir
					}
				}
				return err
			})
			fmt.Print("\033[1A")
		} else {
			my_ls.ListFiles(dirPath, flags)
		}
	}
}
