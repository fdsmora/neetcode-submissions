

// LC 15. 3Sum
func threeSum(nums []int) [][]int {
	res := map[[3]int]struct{}{}
    sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
        k := len(nums)-1
        j := i+1
        for j < k {            
            k_v := -nums[i]-nums[j]
            for j < k && nums[k]>k_v {
                k--
            }
            if j < k && nums[k]==k_v{
                res[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
            }
            j++
		}
	}
    result := [][]int{}
    for triplet := range res {
        result = append(result, []int{triplet[0], triplet[1], triplet[2]})
    }
	return result
} 

func search(nums []int, l, r, target int) (int, bool) {
    for l <= r {
        m := l + (r-l)/2
        if nums[m] == target {
            return m, true
        } else if nums[m] < target {
            l = m + 1
        } else {
            r = m - 1
        }
    }
    return -1, false
}


