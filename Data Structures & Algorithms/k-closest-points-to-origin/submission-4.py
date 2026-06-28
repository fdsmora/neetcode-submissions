import heapq
from math import sqrt

class Point:
    def __init__(self, distance: float, index: int):
        self.distance = distance
        self.index = index

    def __lt__(self, other: 'Node') -> bool:
        return self.distance < other.distance


class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        min_heap = []
        for i, p in enumerate(points):
            dist = sqrt(p[0]**2 + p[1]**2)
            min_heap.append(Point(dist, i))

        heapq.heapify(min_heap)
        res = []
        for i in range(k):
            p = heapq.heappop(min_heap)
            res.append(points[p.index])
        return res