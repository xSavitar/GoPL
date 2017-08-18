// Author: Alangi Derick
// Description: Dup3 prints the text of each line that appears more than once
//              in the standard input, preceded by its count. It reads from stdin
//              or from a list of named files.

// Notes
/* If there is a new line at the end of the file, this particular program also counts
   it. So considering the 2 files I have, you will see something like: "2    ". Also,
   note that this code works only when files are concerned. No Stdin. */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename) // The question here is, does ReadFile()
		//opens the file and closes it when it's done?
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
