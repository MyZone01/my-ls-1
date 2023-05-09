package utils

import "strings"

func OrderFiles(a []string, f func(a, b string) int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if f(strings.ToLower(a[i]), strings.ToLower(a[j])) == 1 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func GetOutputLength(files []string) int {
	var count int

	for i := 0; i < len(files)-1; i++ {
		count+=len(files[i])
	}
	println(count + (len(files)-1)*2)
	return count + (len(files)-1)*2
}
