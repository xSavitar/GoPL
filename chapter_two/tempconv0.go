// Author: Alangi Derick
// Description: Package tempconv performs Celsius and Fahrenheit temperature computations.

// package name is: tempconv
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC	  Celsius = 0
	Boiling       Celsius = 100
)

// Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32) }

// Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }