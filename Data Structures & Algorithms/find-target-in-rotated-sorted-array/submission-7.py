class Solution:
    def search(self, nums: List[int], target: int) -> int:
        e = self.findGreatest(nums)
        print(f"greatest:{e}")
        s, e = self.findHalf(nums, e, target)
        print(f"s={s}, e={e}")
        return self.bs(nums, s, e, target)
    
    def findHalf(self, nums, e, target)->(int,int):
        return (0, e) if nums[0] <= target <= nums[e] else (e+1, len(nums)-1)

    
    def bs(self, nums:list[int], s: int, e: int, target: int) -> int:
        while s<=e:
            m = (s+e)//2
            if nums[m]==target:
                return m
            if nums[m]<target:
                s = m+1
            else:
                e = m-1

        return -1

    def findGreatest(self, nums: list[int]) ->int:
        if nums[-1]>nums[0]:
            return len(nums)-1
        s, e, m = 0,len(nums)-1, 0
        while s<=e:
            m = (s+e)//2
            if nums[m]==nums[s]:
                break
            if nums[m]>nums[s]:
                s = m
            else:
                e = m
        return m
