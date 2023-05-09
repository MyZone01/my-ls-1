package my_ls

import (
	"os"
	"strings"
)

func GetFlags() (bool, bool, bool, bool, bool, string) {
	flags := os.Args[1:]
	var showAll, longFormat, recursive, reverse, sortByTime bool
	dirPath := "."
	for _, flag := range flags {
		if len(flag) == 2 && flag[0] == '-' {
			switch flag {
			case "-a":
				showAll = true
			case "-l":
				longFormat = true
			case "-R":
				recursive = true
			case "-r":
				reverse = true
			case "-t":
				sortByTime = true
			default:
				dirPath = flag
			}
		} else {
			dirPath = flag
		}
	}
	return showAll, longFormat, recursive, reverse, sortByTime, dirPath
}

func joinPath(elements ...string) string {
	// Join the elements using "/"
	path := strings.Join(elements, "/")

	// Replace consecutive slashes with a single slash
	path = strings.ReplaceAll(path, "//", "/")

	// Remove trailing slash if present
	if strings.HasSuffix(path, "/") && len(path) > 1 {
		path = path[:len(path)-1]
	}

	return path
}
