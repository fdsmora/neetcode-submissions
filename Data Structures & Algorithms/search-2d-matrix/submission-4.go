func searchMatrix(matrix [][]int, target int) bool {
	R, C := len(matrix), len(matrix[0])
	n := R*C
	s, e := 0, n-1
	for s <= e {
		m := (s + e) / 2
		i, j := m / C, m % C
		if matrix[i][j]==target {
			return true
		}
		if matrix[i][j] < target {
			s++
		} else {
			e--
		}
	}
	return false
}
