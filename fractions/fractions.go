package fractions

import "fmt"

type Fraction struct {
	numerator, denominator int
}

func Construct(numerator, denominator int) *Fraction {
	fraction := new(Fraction)
	fraction.numerator = numerator
	fraction.denominator = denominator
	return fraction
}

func (fraction Fraction) Display() {
	fmt.Printf("%v/%v\n", fraction.numerator, fraction.denominator)
}
