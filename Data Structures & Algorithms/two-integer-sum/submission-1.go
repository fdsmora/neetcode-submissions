func twoSum(nums []int, target int) []int {
    indices := make(map[int]int, len(nums))

    for i, v := range nums {
        indices[v]=i
    }

    i, j := 0,0
    var ok bool
    for ; i<len(nums); i++ {
        if j, ok = indices[target - nums[i]]; ok  && j != i {
            break
        }
    }
    return []int{i, j}
}
