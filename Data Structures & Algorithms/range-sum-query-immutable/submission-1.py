class NumArray:

    def __init__(self, nums: List[int]):
        self.prefix = nums
        for i in range(1, len(nums)):
            self.prefix[i]=self.prefix[i-1]+nums[i]

    def sumRange(self, left: int, right: int) -> int:
        leftSum = self.prefix[left-1] if left>0 else 0
        return self.prefix[right]-leftSum


# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# param_1 = obj.sumRange(left,right)