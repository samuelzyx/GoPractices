func calculateArea(n int) int {
	if n <= 0 {
		return 0 // Handle invalid input (n should be positive)
	}
	// Base case for 1-interesting polygon
	if n == 1 {
		return 1
	}
	// Recursive case: Area(n) = Area(n-1) + 4(n-1)
	return calculateArea(n-1) + 4*(n-1)
}

func solution(n int) int {
	area := calculateArea(n)
	return area
}
  