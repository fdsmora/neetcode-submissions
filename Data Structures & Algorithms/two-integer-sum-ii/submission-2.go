func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	for i, v := range numbers {
		want := target - v
		if want == v {
			continue
		}
		if j, found := bs(numbers, want, i+1, n-1); found {
			return []int{i+1,j+1}
		}
	}
	return nil
}

func bs(nums []int, want, s, e int) (int, bool){
	for s <= e {
		m := (s+e)/2
		if want == nums[m]{
			return m, true
		}
		if want < nums[m] {
			e = m - 1
		} else {
			s = m + 1
		}
	}
	return 0, false
}
