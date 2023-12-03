package helpers

func MakeGrid[T any](width int, height int) [][]T {
	grid := [][]T{}

	for i := 0; i < height; i = i + 1 {
		grid = append(grid, make([]T, width))
	}

	return grid
}
