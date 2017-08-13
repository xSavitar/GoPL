// Author: Alangi Derick
// Description: Dup1 prints the text of each line that appears more than once
//              in the standard input, preceded by its count.

package main

import (
	"bufio"
	"fmt"
	"os"
)

// %d, %s, as all in C are called by GoLang programmers: 'verbs'

func main() {
	counts := make(map[string]int) // constant-time operation to store, retrieve or test items in the set
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	// After entering all the characters or strings,
	// press "Ctrl + D" (EOF) for the program to process
	// the input and end.
}
