package helper

import (
	"strings"
)

func GetMaxCols(width int, files []string) int {
	var lw int
	for _, v := range files {
		if len(v) > lw {
			lw = len(v)
		}
	}
	return int(width / lw)
}

func GetOutputLength(files []string) int {
	var count int

	for i := 0; i < len(files)-1; i++ {
		count += len(files[i]) + 2
	}
	return count
}

// func GetColNumber(width int, files []os.FileInfo) int {
// 	var biggestFileName int
// 	for _, file := range files {
// 		if len(file.Name()) > biggestFileName {
// 			biggestFileName = len(file.Name())
// 		}
// 	}
// 	return int(width / biggestFileName)
// }

//	func GetOutputLength(files []os.FileInfo) int {
//		var outputLength int
//		for _, file := range files {
//			outputLength += len(file.Name())
//		}
//		return outputLength
//	}
func OrderFiles(a []string, f func(a, b string) int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if f(strings.ToLower(a[i]), strings.ToLower(a[j])) == 1 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
