class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        maxSum, currSum = nums[0], 0
        for n in nums:
            # if n+currSum (prefix sum) improves n then include it in the solution, else
            # do not include it
            currSum = n+currSum if n+currSum >= n else n
            maxSum = max(maxSum, currSum)
        return maxSum