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
        totalCost = [None]*(n+1)
        def minCost(i):
            if totalCost[i] is not None:
                return totalCost[i]
            if i<0:
                return 0
            curCost = 0
            if i<n:
                curCost=cost[i]
            totalCost[i] = curCost + min(minCost(i-1), minCost(i-2))
            return totalCost[i]
        return minCost(n)