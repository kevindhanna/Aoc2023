package helpers

func Sum(values []int) int {
	return Reduce(values, func(t int, v int, _ int) int {
		return t + v
	}, 0)
}
