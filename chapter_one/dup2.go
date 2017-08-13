// Author: Alangi Derick
// Description: Dup2 prints the text of each line that appears more than once
//              in the standard input, preceded by its count. It reads from stdin 
//              or from a list of named files.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 { // read from stdin if no file(s) supplied
		countLines(os.Stdin, counts)
	} else { // read from a file by default if supplied and ignore stdin
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// countLines() function definition
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

	//Note: ignoring potentials errors from input.Err()
}