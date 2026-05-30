func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))

	for i, _ := range nums {
		if i==0 {
			prefix[i]=nums[i]
		}else{
			prefix[i]=nums[i]*prefix[i-1]
		}
		if n-i-1 == n-1 {
			suffix[n-i-1]=nums[n-1]
		}else {
			suffix[n-i-1]=nums[n-i-1]*suffix[n-i]
		}
	}

	for i, _ := range nums {
		var pre, suf int
		if i == 0 {
			pre = 1
		} else {
			pre = prefix[i-1]
		} 
		if i == n-1 {
			suf = 1
		}else{
			suf = suffix[i+1]
		}

		nums[i]=pre*suf
	}
	return nums
}
