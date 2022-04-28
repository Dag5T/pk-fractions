package fractions

import "fmt"

type Fraction struct {
	numerator, denominator int
}

func (numerator, denominator int) Fraction {
	fraction := FracFraction{numerator, denominator}
	return fraction
}

func (fraction Fraction) Display() {
	fmt.Printf("%v/%v\n", fraction.numerator, fraction.denominator)
}
