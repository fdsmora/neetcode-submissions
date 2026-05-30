func maxProfit(prices []int) int {
    maxProfit := 0
    for i:=0; i<len(prices); i++ {
        for j:=i; j<len(prices); j++ {
            maxProfit = max(maxProfit, prices[j]-prices[i])
        }
    }
    return maxProfit
}

func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}
