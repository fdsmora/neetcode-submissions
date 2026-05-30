class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        if len(nums)<=1:
            return len(nums)
        max_count, i = 0,0
        S = {}
        for v in nums:
            S[v]=False
        for v in nums:
            if S[v]==True:
                continue 
            S[v]=True
            count = 1
            l = v-1
            while l in S:
                l-=1
                count+=1
            r = v+1
            while r in S:
                r+=1
                count+=1
            max_count = max(max_count, count)
        return max_count