// Author: Alangi Derick
// Description: Boiling prints the boiling point of water.

package main

import "fmt"

// constant
const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c) // the ° on Mac is done by holding "Option + Shift + 8"

	// Output:
	// boiling point = 212°F or 100°C
}
