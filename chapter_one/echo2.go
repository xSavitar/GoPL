// Author: Alangi Derick
// Description: Echo prints its command-line arguments.

package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", "" // this type of declaration can be used only within a function 
	                 // not at package-level variable declaration.
	
	for _, arg := range os.Args[1:] { // this line has a blank identifier "_"
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}