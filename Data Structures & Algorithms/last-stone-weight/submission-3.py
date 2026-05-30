import heapq

class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:
        #heap.heapify(stones)
        maxHeap = []
        for s in stones:
            heapq.heappush(maxHeap, -s)
        while len(maxHeap)>1:
            st1 = -heapq.heappop(maxHeap)
            st2 = -heapq.heappop(maxHeap)
            rem = st1-st2
            if rem > 0:
                heapq.heappush(maxHeap,-rem)
        if len(maxHeap)==0:
            return 0
        return -maxHeap[0]