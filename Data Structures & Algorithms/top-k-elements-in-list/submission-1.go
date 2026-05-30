func topKFrequent(nums []int, k int) []int {
    freq := map[int]int{}
    buckets := make([][]int, len(nums)+1)

    for _, v := range nums {
        freq[v]++
    }

    for k, v := range freq {
        buckets[v]=append(buckets[v], k)
    }

    res := []int{}
    for i:=len(buckets)-1; i>0; i--{
        if len(buckets[i])>0{
            for _, v := range buckets[i] {
                if k > 0 {
                    res = append(res, v)
                    k--
                }else {
                    return res
                }
            }
        }
    }
    return res
}
