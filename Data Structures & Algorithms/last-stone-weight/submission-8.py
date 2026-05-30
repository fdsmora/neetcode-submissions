import heapq

class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:
        stones = [-s for s in stones]
        heapq.heapify(stones)

        while len(stones)>1:
            st1 = -heapq.heappop(stones)
            st2 = -heapq.heappop(stones)
            if st1-st2>0:
                heapq.heappush(stones,-(st1-st2))
        return -stones[0] if stones else 0

