package triangle

import "math"

type Kind string

const (
	Equ = "equilateral"
	Iso = "isosceles"
	Sca = "scalene"
	NaT = "not a triangle"
)

func KindFromSides(a, b, c float64) Kind {
	if a == 0 || b == 0 || c == 0 ||
		(a > 0 && b*c < 0) || (a < 0 && b*c > 0) ||
		a+b <= c || b+c <= a || c+a <= b ||
		math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	} else if a == b || b == c || c == a {
		return Iso
	} else {
		return Sca
	}
}
