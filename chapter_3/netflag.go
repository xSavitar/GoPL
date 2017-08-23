// Author: Alangi Derick
// Netflag

package main

import (
	"fmt"
)

// Flags used
type Flags uint

// declaration of constants
const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

// IsUp() -- function definition
func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

// TurnDown() -- function definition
func TurnDown(v *Flags) {
	*v &^= FlagUp
}

// SetBroadcast() -- function definition
func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

// IsCast() -- function definition
func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}

func main() {
	var v Flags = FlagMulticast | FlagUp

	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)

	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)

	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}
