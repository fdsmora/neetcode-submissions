class Solution:
    def combinationSum2(self, candidates: list[int], target: int) -> list[list[int]]:
        res = []
        total = 0
        temp = []
        candidates.sort()
        def comb(temp: list[int], j: int, total: int):
            n = len(candidates)
            if total == target:
                res.append(list(temp))
                return None
            
            if j >= n or total > target:
                return None

            for i in range(j, n):
                if i > j and candidates[i] == candidates[i-1]:
                    continue
                temp.append(candidates[i])

                comb(temp, i+1, total+candidates[i])
                temp.pop()
        comb(temp, 0, total)
        return res
