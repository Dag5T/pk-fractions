package main

import (
	"fmt"
	"fractions/fractions"
)

func main() {
	testFraction := fractions.NewDefault()
	fmt.Println(testFraction.Display())
	fmt.Println("OK")
}
