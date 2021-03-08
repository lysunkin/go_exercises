package diffsquares

func SquareOfSum(s int) int {
	res := 0
	for i := 1; i <= s; i++ {
		res += i
	}

	return res*res
}

func SumOfSquares(s int) int {
	res := 0
	for i := 1; i <= s; i++ {
		res += i*i
	}

	return res
}

func Difference(s int) int {
	return SquareOfSum(s)-SumOfSquares(s)
}
