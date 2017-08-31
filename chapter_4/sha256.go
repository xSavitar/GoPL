// Author: Alangi Derick N
// Prints the digest / hash of 2 messages "x" and "X"

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	// hash / digest of "x"
	c1 := sha256.Sum256([]byte("x"))

	// hash / digest of "X"
	c2 := sha256.Sum256([]byte("X"))

	// print the hash and do some basic comparison
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
