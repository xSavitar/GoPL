// Author: Alangi Derick
// Description: Package tempconv0 performs Celsius and Fahrenheit temperature computations.

// package name is: tempconv0
package tempconv

import "fmt"

// Two UD Types in this package
type Celsius float64
type Fahrenheit float64

// constants
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
