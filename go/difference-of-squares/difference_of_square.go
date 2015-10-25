package diffsquares

func SquareOfSums(n int) int {
	var out int = 0
	for i := 1; i <= n; i++ {
		out += i
	}
	return out * out
}

func SumOfSquares(n int) int {
	var out int = 0
	for i := 1; i <= n; i++ {
		out += i * i
	}
	return out
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
