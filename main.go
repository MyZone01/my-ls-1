package main

import (
	my_ls "my_ls/lib"
)

func main() {
	// Parse command line flags
	showAll, longFormat, recursive, reverse, sortByTime, dirPath := my_ls.GetFlags()

	// Perform the directory listing
	my_ls.ListFiles(dirPath, showAll, longFormat, recursive, reverse, sortByTime)
}
