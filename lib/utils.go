package my_ls

import (
	"math"
	"os"
	"strings"
	"syscall"
	"unsafe"
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

func GetTerminalWidth() int {
	winsize := struct {
		Row    uint16
		Col    uint16
		Xpixel uint16
		Ypixel uint16
	}{}
	_, _, _ = syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&winsize)))

	return int(winsize.Col)
}

func GetColNumber(width int, files []os.FileInfo) int {
	var biggestFileName int
	for _, file := range files {
		if len(file.Name()) > biggestFileName {
			biggestFileName = len(file.Name())
		}
	}
	return int(math.Floor(float64(width) / float64(biggestFileName+2)))
}

func GetOutputLength(files []os.FileInfo) int {
	var outputLength int
	for _, file := range files {
		outputLength += len(file.Name())
	}
	return outputLength + (len(files)-1)*2
}

func OrderFiles(files []os.FileInfo, f func(a, b string) int) {
	var currentFileName, nextFileName string

	for i := 0; i < len(files); i++ {
		currentFileName = files[i].Name()
		for j := i + 1; j < len(files); j++ {
			nextFileName = files[j].Name()

			if strings.HasPrefix(currentFileName, ".") && len(currentFileName) > 1 {
				currentFileName = currentFileName[1:]
			}
			if strings.HasPrefix(nextFileName, ".") && len(nextFileName) > 1 {
				nextFileName = nextFileName[1:]
			}
			if f(strings.ToLower(currentFileName), strings.ToLower(nextFileName)) == 1 {
				files[i], files[j] = files[j], files[i]
			}
		}
	}
}
