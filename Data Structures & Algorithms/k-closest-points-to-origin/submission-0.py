import heapq
from math import sqrt

class Point:
    def __init__(self, index: int, distance: float):
        self.index = index
        self.distance = distance


class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        min_heap = []
        for i, p in enumerate(points):
            dist = sqrt(p[0]**2 + p[1]**2)
            min_heap.append((dist, i))

        heapq.heapify(min_heap)
        res = []
        for i in range(k):
            elem = heapq.heappop(min_heap)
            res.append(points[elem[1]])
        return res