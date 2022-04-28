package main

import (
	"fmt"
	"fractions/fractions"
)

func main() {
	testFraction := fractions.Construct(2, 4)
	testFraction.Display()
	fmt.Println("OK")
}
