func searchMatrix(matrix [][]int, target int) bool {
	R, C := len(matrix), len(matrix[0])
	n := R*C
	s, e := 0, n-1
	for s <= e {
		m := (s + e) / 2
		i, j := m / C, m % C
		fmt.Printf("s=%d,e=%d,m=%d,matrix[i][j]=%d\n", s, e, m, matrix[i][j])
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
