func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	prod := 1
	zeroCount := 0
	for _, n := range nums {
		if n == 0 {
			zeroCount++
			if zeroCount > 1 {
				return res
			}
			continue
		}
		prod *= n
	}
	for i, n := range nums {
		if n == 0 {
			res[i] = prod
		} else if zeroCount == 0 {
			res[i] = prod / n
		}
	}
	return res
}
