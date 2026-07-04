class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        # totalCost = [0]*len(cost)
        # def minCost(i: int)->int:
        #     if i>=len(cost):
        #         return 0
        #     curCost = 0
        #     if i<0:
        #         curCost = 0
        #     else:
        #         curCost = cost[i]
        #         totalCost[i]=
        #     return curCost + min(minCost(i+1), minCost(i+2))

        # return minCost(-1)
        n = len(cost)
        totalCost = [None]*n
        def minCost(i):
            if i<0:
                return 0
            curCost = 0
            if i<n:
                if totalCost[i] is not None:
                    return totalCost[i]
                curCost=cost[i]
            res = curCost + min(minCost(i-1), minCost(i-2))
            if i<n:
                totalCost[i]= res
            return res
        return minCost(n)