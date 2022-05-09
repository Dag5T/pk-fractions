package fractions

import "fmt"

type Fraction struct {
	numerator, denominator int
}

func NewDefault() Fraction {
	fraction := Fraction{1, 1}
	return fraction
}

func New(numerator, denominator int) Fraction {
	fraction := Fraction{numerator, denominator}
	return fraction
}

func (fraction Fraction) Display() string {
	return fmt.Sprintf("%v/%v", fraction.numerator, fraction.denominator)
}
