func notExistFreeRoomBelow(matrix [][]int, currentRow, currentCol int) bool {
	for x := 0; x < currentRow; x++ {
		if matrix[x][currentCol] == 0 {
			return false
		}
	}
	return true
}

func solution(matrix [][]int) int {
	sum := 0

	rows := len(matrix)
	cols := len(matrix[0])

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			value := matrix[x][y]

			if value != 0 && notExistFreeRoomBelow(matrix, x, y) {
				sum += value
			}
		}
	}

	return sum
}