// Author: Alangi Derick N
// Nonempty is an example of an in-place slice algorithm.

package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// Doing non-empty using append()
func nonemptyAppend(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// main function definition
func main() {
	data := []string{"one", "", "three"}

	// using the define nonempty()
	fmt.Printf("Using the default nonempty() function\n")
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`

	info := []string{"one", "", "three"}
	// using the append function
	fmt.Printf("\nUsing the append() version of nonempty()\n")
	fmt.Printf("%q\n", nonemptyAppend(info))
	fmt.Printf("%q\n", info)
}
