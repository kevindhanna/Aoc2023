package helpers

func Sum(values []int) int {
	return Reduce(values, func(t int, v int, _ int) int {
		return t + v
	}, 0)
}

func Min(values []int) int {
	min := values[0]
	for _, v := range values[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func LowestCommonMultiple(a int, b int) int {
	d := GreatestCommonFactor(a, b)
	return a * b / d
}

func GreatestCommonFactor(a int, b int) int {
	var R int
	for (a % b) > 0 {
		R = a % b
		a = b
		b = R
	}
	return b
}
