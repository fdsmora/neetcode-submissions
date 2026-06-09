class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        maxSum, prefixSum = nums[0], 0
        for n in nums:
            # if n+prefixSum improves n then include it in the solution, else
            # do not include it
            prefixSum = n+prefixSum if n+prefixSum >= n else n
            maxSum = max(maxSum, prefixSum)
        return maxSum