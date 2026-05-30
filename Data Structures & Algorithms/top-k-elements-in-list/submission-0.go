func topKFrequent(nums []int, k int) []int {
    count := map[int]int{}
    for _, v := range nums {
        count[v]++
    }
    ordered := [][]int{}
    for k, v := range count {
        ordered = append(ordered, []int{k,v})
    }
    sort.Slice(ordered, func(i, j int) bool {
        return ordered[i][1] > ordered[j][1]
    })
    res := []int{}
    for _, v := range ordered[0:k] {
        res = append(res, v[0])
    }
    return res
}
