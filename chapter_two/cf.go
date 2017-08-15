// Author: Alangi Derick
// Cf converts its numeric argument to Celsius and Fahrenheit.

package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl/chapter_two/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t) // f is of type Fahrenheit
		c := tempconv.Celsius(t)    // c is of type Celsius

		// do the conversion
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
