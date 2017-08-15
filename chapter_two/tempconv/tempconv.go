// Author: Alangi Derick
// Description: Package tempconv performs Celsius and Fahrenheit temperature computations.

// package name is: tempconv

package tempconv

import "fmt"

// Celcius -- UD type in this package
type Celsius float64

// Fahrenheit -- UD type in this package
type Fahrenheit float64

// constants
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
