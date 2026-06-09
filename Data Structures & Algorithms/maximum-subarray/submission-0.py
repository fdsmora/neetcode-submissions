class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        maxSum, currSum = -1000,-1000
        for i in range(len(nums)):
            # if nums[i]+currSum improves nums[i] then include it in the solution, else
            # do not include it
            currSum = nums[i]+currSum if nums[i]+currSum >= nums[i] else nums[i]
            maxSum = max(maxSum, currSum)
        return maxSum