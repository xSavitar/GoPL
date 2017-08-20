// Author: Alangi Derick
// basename removes directory components and a .suffix. e.g., a => a, 
// a.go => a, a/b/c.go => c, a/b.c.go => b.c

package main

import (
	"fmt"
)

// definition of the basename() function
func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

// main function
func main() {
	fmt.Println(basename("go/src/basename1.go"))
}