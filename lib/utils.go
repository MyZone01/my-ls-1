package my_ls

import (
	"os"
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

func GetColNumber(width int, files []os.FileInfo) int {
	var biggestFileName int
	for _, file := range files {
		if len(file.Name()) > biggestFileName {
			biggestFileName = len(file.Name())
		}
	}
	return int(width / biggestFileName)
}

func GetOutputLength(files []os.FileInfo) int {
	var outputLength int
	for _, file := range files {
		outputLength += len(file.Name())
	}
	return outputLength
}

func OrderFiles(files []os.FileInfo, f func(a, b string) int) {
	for i := 0; i < len(files); i++ {
		for j := i + 1; j < len(files); j++ {
			if f(files[i].Name(), files[j].Name()) == 1 {
				files[i], files[j] = files[j], files[i]
			}
		}
	}
}

func OrderFileName(fileNames []string, f func(a, b string) int) {
	for i := 0; i < len(fileNames); i++ {
		for j := i + 1; j < len(fileNames); j++ {
			if f(fileNames[i], fileNames[j]) == 1 {
				fileNames[i], fileNames[j] = fileNames[j], fileNames[i]
			}
		}
	}
}
