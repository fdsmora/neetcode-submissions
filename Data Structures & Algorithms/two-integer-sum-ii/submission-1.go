func twoSum(numbers []int, target int) []int {
	hm := map[int]int{}
	for i, v := range numbers {
		hm[v]=i
	}
	for i, v := range numbers {
		if j, ok:=hm[target-v]; ok {
			return []int{i+1,j+1}
		}
	}
	return nil
}
