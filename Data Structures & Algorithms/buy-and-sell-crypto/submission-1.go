func maxProfit(prices []int) int {
    var left, maxP = 0, 0
    for right:=1; right < len(prices); right++ {
        if prices[left]<prices[right] {
            profit := prices[right]-prices[left]
            maxP = max(maxP, profit)
        }else{
            left = right
        }
    }
    return maxP
}

func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}
