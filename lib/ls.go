package my_ls

import (
	"fmt"
	"io/fs"
	"math"
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

func ListFiles(dirPath string, flags Flag) {
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
	if len(entries) == 0 {
		return
	}
	OrderFiles(entries, strings.Compare)
	if flags.ShowAll {
		_dotEntries := []os.FileInfo{}
		_entries := []os.FileInfo{}
		currentFolderInfos, err := os.Lstat(currentFolderName)
		if err != nil {
			fmt.Println("Error getting file info:", err)
		}
		_dotEntries = append(_dotEntries, currentFolderInfos)

		parentFolderInfos, err := os.Lstat(parentFolderName)
		if err != nil {
			fmt.Println("Error getting file info:", err)
		}
		_dotEntries = append(_dotEntries, parentFolderInfos)

		for i := 0; i < len(entries); i++ {
			if strings.HasPrefix(entries[i].Name(), ".") {
				_dotEntries = append(_dotEntries, entries[i])
			} else {
				_entries = append(_entries, entries[i])
			}
		}
		entries = append(_dotEntries, _entries...)
	}

	// Sort the entries based on the specified options
	if flags.SortByTime {
		sortByModificationTime(entries)
	}

	// Reverse the entries
	if flags.Reverse {
		entries = reverseEntries(entries)
	}

	if flags.LongFormat {
		var TotalSize int
		for _, v := range entries {
			if !flags.ShowAll && strings.HasPrefix(v.Name(), ".") || (v.Name() == "." || v.Name() == "..") {
				continue
			}
			if v.Size() < 4096 {
				continue
			}

			TotalSize += int((v.Size() + 4096 - 1) / 4096 * (4096 / 1024))
		}
		fmt.Printf("total %v\n", TotalSize)
		for _, entry := range entries {
			fileName := entry.Name()

			if !flags.ShowAll && strings.HasPrefix(fileName, ".") {
				continue
			}
			OrderFiles(entries, strings.Compare)
			printLongFormat(entry)
		}
	} else {
		entries := SortByFileName(entries)
		width := GetTerminalWidth()
		numberOfColumn, maxWordColumn := GetColNumber(width, entries)
		numberOfLine := int(math.Ceil(float64(len(entries)) / float64(numberOfColumn)))
		printShortFormat(numberOfLine, numberOfColumn, entries, maxWordColumn)
	}
}

func printShortFormat(numberOfLine int, numberOfColumn int, entries []fs.FileInfo, maxWordColumn []int) {
	for i := 0; i < numberOfLine; i++ {
		for j := 0; j < numberOfColumn; j++ {
			index := (numberOfLine * j) + i

			if index > len(entries)-1 {
				break
			}

			fmt.Print(entries[index].Name())
			if j < numberOfColumn-1 {
				rest := maxWordColumn[j] - len(entries[index].Name())
				for i := 0; i < rest; i++ {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}

func reverseEntries(entries []os.FileInfo) []os.FileInfo {
	length := len(entries)
	reversed := make([]os.FileInfo, length)

	for i, entry := range entries {
		reversed[length-i-1] = entry
	}

	return reversed
}

func sortByModificationTime(entries []os.FileInfo) {
	n := len(entries)
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if entries[i].ModTime().After(entries[i-1].ModTime()) {
				entries[i], entries[i-1] = entries[i-1], entries[i]
				swapped = true
			}
		}
		n--
	}
}

func printLongFormat(entry os.FileInfo) {
	//Get the file name
	name := entry.Name()

	permissions := entry.Mode().String()

	if permissions[0] == 'L' {
		dst, _ := os.Readlink(name)
		dstInfo, err := os.Stat(dst)
		if err != nil {
			fmt.Println("Error getting file info:", err)
		}
		name = name + " -> " + dstInfo.Name()
		permissions = string(append([]rune{'l'}, []rune(permissions[1:])...))
	}

	numHardLinks := entry.Sys().(*syscall.Stat_t).Nlink

	size := entry.Size()

	modTime := entry.ModTime().Format("Jan _2 15:04")

	owner, err := user.LookupGroupId(strconv.Itoa(int(entry.Sys().(*syscall.Stat_t).Uid)))
	if err != nil {
		fmt.Printf("Error retrieving owner information for %s: %s\n", name, err)
		return
	}

	group, err := user.LookupGroupId(strconv.Itoa(int(entry.Sys().(*syscall.Stat_t).Gid)))
	if err != nil {
		fmt.Printf("Error retrieving group information for %s: %s\n", name, err)
		return
	}

	fmt.Printf("%s %2d %s %s %5d %s %s\n", permissions, numHardLinks, owner.Name, group.Name, size, modTime, name)
}
