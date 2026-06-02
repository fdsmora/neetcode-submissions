class Solution:
    def maxTurbulenceSize(self, arr: List[int]) -> int:
        maxLen = 1
        if len(arr)==1:
            return maxLen
        
        def greaterThan(a: int, b:int) -> bool:
            return a > b
        def lessThan(a: int, b:int) -> bool:
            return a < b
        
        l, r = 0, 0
        comps = [greaterThan, lessThan]

        while r +1 < len(arr):
            if comps[r%2](arr[r], arr[r+1]):
                r+=1
            elif arr[r]==arr[r+1]:
                maxLen = max(maxLen, r-l+1)
                r+=1
                l=r
            else:
                maxLen = max(maxLen, r-l+1)
                comps[0], comps[1] = comps[1], comps[0]
                l = r
        return max(maxLen, r-l+1)

                