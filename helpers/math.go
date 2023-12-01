package helpers

func Sum(values []int) int {
	total := 0
	for i := 0; i < len(values); i = i + 1 {
		total = total + values[i]
	}
	return total
}
