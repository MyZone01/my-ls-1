package utils

import (
	"fmt"
	"my-ls-1/lib/sys"
	"strings"
)

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
		count += len(files[i])
	}
	// println(count + (len(files)-1)*2)
	return count + (len(files)-1)*2
}

func PrintFiles(files []string) {
	width := sys.GetTerminalWidth()
	size := GetOutputLength(files)
	var curColAt = 1
	var curLinAt = 1
	var nCol, temp int

	if width-size >= 0 {
		for _, file := range files {
			fmt.Printf("\033[%dG%v", curColAt, file)
			curColAt += len(file) + 2

		}
	} else {
		maxCols := maxCols(width, files)

		l := (len(files) + maxCols) / maxCols
		lin := ""
		var lastf string
		for i, file := range files {
			if strings.Contains(file, " ") {
				file = fmt.Sprintf("'%v'", file)
			}
			if i < l {
				fmt.Printf("%v", file)
				if i < l-1 {
					fmt.Println()
				}
			} else {
				if i%l == 0 {
					lin = fmt.Sprintf("\033[%vA", l-1)
					curColAt = temp + nCol
					curLinAt = 1
					nCol += temp
				} else {
					lin = "\033[1B"
				}
				fmt.Printf("%s\033[%dG%v", lin, curColAt+1, file)
			}
			if len(file) > len(lastf) {
				temp = len(file) + 2

			}
			lastf = file
			curLinAt++
		}
		fmt.Printf("\033[%vB", l-curLinAt+1)
	}
}

func maxCols(width int, files []string) int {
	var lw int
	for _, v := range files {
		if len(v) > lw {
			lw = len(v)
		}
	}
	return width / lw
}
