package my_ls

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

var (
	currentFolderName = "."
	parentFolderName  = ".."
)

func ListFiles(dirPath string, showAll, longFormat, recursive, reverse, sortByTime bool) {
	// Read the directory entries
	_dir, err := os.Open(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	entries, err := _dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	if showAll {
		parentFolderInfos, err := os.Lstat(parentFolderName)
		if err != nil {
			fmt.Println("Error getting file info:", err)
		}
		entries = append([]os.FileInfo{parentFolderInfos}, entries...)

		currentFolderInfos, err := os.Lstat(currentFolderName)
		if err != nil {
			fmt.Println("Error getting file info:", err)
		}
		entries = append([]os.FileInfo{currentFolderInfos}, entries...)
	}

	width := GetTerminalWidth()
	size := GetOutputLength(entries)
	var curColAt = 1
	var curLinAt = 1
	var nCol, temp int

	// Sort the entries based on the specified options
	if sortByTime {
		sortByModificationTime(entries)
	}

	// Reverse the entries
	if reverse {
		entries = reverseEntries(entries)
	}

	// Print the entries
	for i, entry := range entries {
		fileName := entry.Name()

		// Skip hidden files if the showAll flag is not set
		if !showAll && strings.HasPrefix(fileName, ".") {
			continue
		}

		if longFormat {
			// Print long listing format
			printLongFormat(entry)
		} else {
			// Print the file/directory name
			fileName := entry.Name()
			if width-size >= 0 {
				fmt.Printf("\033[%dG%v", curColAt, fileName)
				curColAt += len(fileName) + 2
			} else {
				maxCols := GetColNumber(width, entries)
				l := (len(fileName) + maxCols) / maxCols
				line := ""
				if strings.Contains(fileName, " ") {
					fileName = "'" + fileName + "'"
				}
				printShortFormat(fileName, i, l, line, curColAt, temp, nCol, curLinAt)
			}
			if i == len(entries)-1 {
				fmt.Println()
			}
		}

		// Recursively list subdirectories if the recursive flag is set
		if recursive && entry.IsDir() {
			subDirPath := joinPath(dirPath, fileName)
			fmt.Println("\n" + subDirPath + ":")
			ListFiles(subDirPath, showAll, longFormat, recursive, reverse, sortByTime)
		}
	}
}

// Reverse the entries
func reverseEntries(entries []os.FileInfo) []os.FileInfo {
	length := len(entries)
	reversed := make([]os.FileInfo, length)

	for i, entry := range entries {
		reversed[length-i-1] = entry
	}

	return reversed
}

// SortByModificationTime sorts an array of fs.DirEntry objects by modification time.
func sortByModificationTime(entries []os.FileInfo) {
	n := len(entries)
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			// Compare modification time of current entry and previous entry
			if entries[i].ModTime().Before(entries[i-1].ModTime()) {
				// Swap entries if the current one is older
				entries[i], entries[i-1] = entries[i-1], entries[i]
				swapped = true
			}
		}
		n--
	}
}

func printShortFormat(fileName string, i int, l int, lin string, curColAt int, temp int, nCol int, curLinAt int) int {
	if i < l {
		fmt.Printf("%v", fileName)
		if i < l-1 {
			fmt.Println()
		}
	} else {

		if i%l == 0 {
			lin = fmt.Sprintf("\033[%vA", l-1)
			curColAt = temp + nCol
			nCol += temp
			temp = len(fileName) + 2
			curLinAt = 1
		} else {
			lin = "\033[1B"
		}
		fmt.Printf("%s\033[%dG%v", lin, curColAt+1, fileName)
	}
	if len(fileName) > temp {
		temp = len(fileName) + 2

	}

	curLinAt++
	return curLinAt
}

func printLongFormat(entry os.FileInfo) {
	//Get the file name
	name := entry.Name()

	// Get the file/directory mode and permissions
	permissions := entry.Mode().String()

	// Get the number of hard links
	numHardLinks := entry.Sys().(*syscall.Stat_t).Nlink

	// Get the file/directory size
	size := entry.Size()

	// Get the file/directory modification time
	modTime := entry.ModTime().Format("Jan _2 15:04")

	// Get the file/directory owner
	owner, err := user.LookupGroupId(strconv.Itoa(int(entry.Sys().(*syscall.Stat_t).Uid)))
	if err != nil {
		fmt.Printf("Error retrieving owner information for %s: %s\n", name, err)
		return
	}

	// Get the file/directory owner's group
	group, err := user.LookupGroupId(strconv.Itoa(int(entry.Sys().(*syscall.Stat_t).Gid)))
	if err != nil {
		fmt.Printf("Error retrieving group information for %s: %s\n", name, err)
		return
	}

	// Print the long format
	fmt.Printf("%s %2d %s %s %5d %s %s\n", permissions, numHardLinks, owner.Name, group.Name, size, modTime, name)
}
