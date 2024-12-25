package utils

import (
	"fmt"
)

func splitLines(s string) []string {
    var lines []string
    start := 0
    for i, char := range s {
        if char == '\n' {
            lines = append(lines, s[start:i])
            start = i + 1
        }
    }
    lines = append(lines, s[start:])
    return lines
}


func Index[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

func Contains[S ~[]E, E comparable](s S, v E) bool {
	return Index(s, v) >= 0
}

func PrintBanner(){
	banner := `
	____  _____ ____  ____   ___        _ 
	| __ )| ____|  _ \|  _ \ / _ \      / |
	|  _ \|  _| | | | | |_) | | | |     | |
	| |_) | |___| |_| |  _ <| |_| |  _  | |
	|____/|_____|____/|_| \_\\___/  (_) |_|

		BedroConfuser
	`

    lines := splitLines(banner)
    for _, line := range lines {
        fmt.Println(line)
    }
}

