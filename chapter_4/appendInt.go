// Author: Alangi Derick N
// Append integers User Defined Function (UDF)

package main

import "fmt"

// Append integers to a slice
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// there is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

// Variadic appendInt function
func appendIntVariadic(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	// ...expand z to at least zlen...
	if zlen <= cap(x) {
		// there is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	copy(z[len(x):], y)
	return z
}

// main function definition
func main() {
	var x, y []int
	fmt.Println("Call to the appendInt() function")
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d    cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	// using the variadic append function
	y = appendIntVariadic(x, 11, 12, 13, 14, 15)
	fmt.Printf("\n\nCall to the appendIntVariadic() function\n")
	fmt.Printf("cap=%d\t%v\n", cap(y), y)
}
