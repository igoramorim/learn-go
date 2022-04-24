package main

import (
	"fmt"
	"time"
)

func main() {
	// For
	k := 1
	for k < 3 {
		fmt.Println(k)
		k++
	}

	fmt.Println()

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	for { // Runs until a break occurs
		fmt.Println("loop")
		break
	}

	fmt.Println()

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	fmt.Println()

	// If
	if k < 5 {
		fmt.Println("k is < 5")
	}

	if n := 10; n < 5 {
		fmt.Println("n is < 5")
	} else {
		fmt.Println("n is NOT < 5")
	}

	fmt.Println()

	// Swtich
	// Go's switch doesn't need 'break'. It only runs the selected case
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	fmt.Println()

	t := time.Now()
	// It also accepts no condition
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}
