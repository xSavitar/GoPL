// Author: Alangi Derick
// Description: Echo prints its command-line arguments.

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	// This is the only loop statement in Go
	// Why?
	for i := 1; i < len(os.Args); i++ { // Opening braces must be here and not below (syntax / convention)
		s += sep + os.Args[i]
		sep = " "
	}
	// Looking at the out of this input: 1   2   3   4   5
	// Output is: 1 2 3 4 5 meaning it doesn't consider the 
	// extra spaces. Why?
	fmt.Println(s)
}