// Author: Alangi Derick
// Description: Dup2 prints the text of each line that appears more than once
//              in the standard input, preceded by its count. It reads from stdin 
//              or from a list of named files.

// Notes:
/* Supplying multiple files with the same content doesn't treat both separately but 
   combines the content of both and treat as a single file. So, if the word "go" is 
   in file1.txt and also in file2.txt, the output is "2     a" and not "1     a", 
   "1    a" as if they were separate files. Interesting indeed */

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
			f, err := os.Open(arg) // there is a words.txt file you can supply as os.Args[1]
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