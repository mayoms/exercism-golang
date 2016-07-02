package triangle

import "math"

const testVersion = 2

type Kind string

const (
	NaT Kind = "NaT" // not a triangle
	Equ Kind = "Equ" // equilateral
	Iso Kind = "Iso" // isosceles
	Sca Kind = "Sca" // scalene
)

func ValidateSides(args ...float64) bool {
	for i := range args {
		if args[i] <= 0 || math.IsNaN(args[i]) || math.IsInf(args[i], 0) {
			return false
		}
	}
	return true
}

func KindFromSides(a, b, c float64) Kind {

	if !ValidateSides(a, b, c) {
		return NaT
	} else if a+b < c || a+c < b || b+c < a {
		return NaT
	} else if a == b && b == c {
		return Equ
	} else if a == b || a == c || b == c {
		return Iso
	}

	return Sca

}
