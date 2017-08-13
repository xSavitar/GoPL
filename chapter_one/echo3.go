// Author: Alangi Derick
// Description: Echo prints its command-line arguments.

package main

import (
	"fmt"
	"os"
	"strings" // additional package to process strings
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " ")) // calling the .Join() function over the strings
}