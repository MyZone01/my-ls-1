package helper

import (
	"math"
	"strings"
)

func GetMaxCols(width int, files []string) int {
	var lw int
	for _, v := range files {
		if len(v) > lw {
			lw = len(v) + 2
		}
	}
	return int(math.Floor(float64(width) / float64(lw)))
}

func GetOutputLength(files []string) int {
	var count int

	for i := 0; i < len(files)-1; i++ {
		count += len(files[i])
	}
	return count + (len(files)-1)*2
}

func OrderFiles(a []string, f func(a, b string) int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if f(strings.ToLower(a[i]), strings.ToLower(a[j])) == 1 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
