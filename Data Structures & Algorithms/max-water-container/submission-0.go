func maxArea(heights []int) int {
    l, r := 0, len(heights)-1
    maxCap :=0
    shortest := 0
    for l < r {
        base := r - l
        if heights[l]<heights[r]{
            shortest = heights[l]
            l++
        }else{
            shortest = heights[r]
            r--
        }
        cap := base*shortest
        if cap > maxCap {
            maxCap = cap
        }
    }
    return maxCap
}
