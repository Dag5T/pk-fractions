package test

import (
	"fmt"
	"fractions/fractions"
)

func UnitTest() {
	test1 := fractions.NewDefault()
	fmt.Println(test1.Display())
}
