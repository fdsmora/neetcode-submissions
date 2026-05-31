class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        l = 0
        seen = {}
        for r in range(len(nums)):
            if nums[r] in seen and abs(r - seen[nums[r]]) <= k:
                return True
            seen[nums[r]] = r
        return False
