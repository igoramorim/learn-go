package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	b := false
	s := "igor"
	i := 10

	var by byte // Alias for uint8
	by = 65

	var r rune // Alias for int32. Represents a Unicode code point
	r = 'a'

	f := 3.14
	z := cmplx.Sqrt(-5 + 12i)

	fmt.Printf("Type: %T Value: %v\n", b, b)
	fmt.Printf("Type: %T Value: %v\n", s, s)
	fmt.Printf("Type: %T Value: %v\n", i, i)
	fmt.Printf("Type: %T Value: %v\n", by, by)
	fmt.Printf("Type: %T Value: %v\n", r, r)
	fmt.Printf("Type: %T Value: %v\n", f, f)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Conversion
	fmt.Println("Conversions")

	ifloat := float64(i)
	fmt.Printf("Type: %T Value: %v\n", ifloat, ifloat)
}
