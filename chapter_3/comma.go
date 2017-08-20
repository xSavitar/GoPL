// Author: Alangi Derick
// comma inserts commas in a non-negative decimal integers string.
package main

import (
	"fmt"
)

// function definition on comma
func comma(s string) string {
	n := len(s)
	if n <= 2 {
		return s
	}
	return comma(s[:n-2]) + "," + s[n-2:]
}

func main() {
	fmt.Println(comma("GoLang"))
}
