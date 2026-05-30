func maxArea(heights []int) int {
    if len(heights)==2 {
        return min(heights[0], heights[1])
    }
    l, r := 0, len(heights)-1
    maxCap := min(heights[l], heights[r])*(r-l)
    for l < r {
        if heights[l]<heights[r]{
            l++
        }else{
            r--
        }
    /*    if heights[l]<heights[r]{
            l++
           if heights[l+1]>heights[r-1]{
                l++
            }else{
                r--
            }
        } else if heights[r-1]>heights[l+1]{
            r--
        }else{
            l++
            r--
        } 
        */
        currCap := min(heights[l], heights[r])*(r-l)
        if currCap > maxCap{
            maxCap = currCap
        }
    }
    return maxCap
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}