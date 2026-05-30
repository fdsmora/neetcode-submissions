func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	var l, r = 0, n-1
	for l < r {
		want := target - numbers[l]
		for numbers[r]>=want{
			if want==numbers[r]{
				return []int{l+1,r+1}
			}
			r--
		}
		l++
	}
	return nil
}

