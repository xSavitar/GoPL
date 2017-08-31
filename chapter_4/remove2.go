// Author: Alangi Derick N
// Remove an element / item from a slice but not preserving the order of the slice

package main

import "fmt"

// define the remove function
func remove(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

// main function definition
func main() {
	s := []int{5, 6, 7, 8, 9}

	// remove the element specified and print the new slice
	fmt.Println("The updated slice is: ", remove(s, 2))
}
