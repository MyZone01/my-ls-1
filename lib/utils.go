package my_ls

import (
	"fmt"
	"math"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

type Flag struct {
	ShowAll    bool
	LongFormat bool
	Recursive  bool
	Reverse    bool
	SortByTime bool
}

func GetFlags() (Flag, []string) {
	args := os.Args[1:]
	flags := Flag{}
	dirPath := []string{}
	for _, flag := range args {
		if len(flag) == 2 && flag[0] == '-' {
			switch flag {
			case "-a":
				flags.ShowAll = true
			case "-l":
				flags.LongFormat = true
			case "-R":
				flags.Recursive = true
			case "-r":
				flags.Reverse = true
			case "-t":
				flags.SortByTime = true
			default:
				dirPath = append(dirPath, flag)
			}
		} else {
			dirPath = append(dirPath, flag)
		}
	}
	if len(dirPath) == 0 {
		dirPath = append(dirPath, ".")
	}
	return flags, dirPath
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

func GetColNumber(width int, files []os.FileInfo) (int, []int) {
	var biggestFileName int
	filesLength := []int{}
	for _, file := range files {
		if len(file.Name()) > biggestFileName {
			biggestFileName = len(file.Name())
		}
		filesLength = append(filesLength, len(file.Name())+2)
	}
	numberOfColumn := int(math.Ceil(float64(width) / float64(biggestFileName)))
	numberOfLine := int(math.Ceil(float64(len(files)) / float64(numberOfColumn)))
	longestLine, tempMaxWordColumn := GetLongestLine(numberOfLine, numberOfColumn, filesLength)
	maxWordColumn := []int{}
	for longestLine-2 < width {
		maxWordColumn = tempMaxWordColumn
		numberOfColumn++
		numberOfLine = int(math.Ceil(float64(len(files)) / float64(numberOfColumn)))
		longestLine, tempMaxWordColumn = GetLongestLine(numberOfLine, numberOfColumn, filesLength)
	}
	numberOfColumn--
	return numberOfColumn, maxWordColumn
}

func PrintFiles(files []os.FileInfo) {
	fmt.Println()
	for _, file := range files {
		fmt.Print(file.Name() + " ")
	}
	fmt.Println()
}

func GetLongestLine(numberOfLine int, numberOfColumn int, filesLength []int) (int, []int) {
	longestLine := 0
	maxWordColumn := []int{}
	for i := 0; i < numberOfColumn; i++ {
		actual := 0
		longestWordColumn := 0
		for j := 0; j < numberOfLine; j++ {
			index := (numberOfLine * i) + j
			if index > len(filesLength)-1 {
				break
			}
			actual = filesLength[index]
			if actual > longestWordColumn {
				longestWordColumn = actual
			}
		}
		maxWordColumn = append(maxWordColumn, longestWordColumn)
		longestLine += longestWordColumn
	}
	return longestLine, maxWordColumn
}

func GetOutputLength(files []os.FileInfo) int {
	var outputLength int
	for _, file := range files {
		outputLength += len(file.Name())
	}
	return outputLength
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

func SortByFileName(files []os.FileInfo) []os.FileInfo {
	n := len(files)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if files[i-1].Name() > files[i].Name() {
				files[i-1], files[i] = files[i], files[i-1]
				swapped = true
			}
		}
		n--
	}
	return files
}
