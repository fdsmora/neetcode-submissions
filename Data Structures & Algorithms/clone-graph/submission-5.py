"""
# Definition for a Node.
class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []
"""
from collections import deque
class Solution:
    def cloneGraph(self, node: Optional['Node']) -> Optional['Node']:
        if not node:
            return None
        visited = dict()
        head = Node(node.val)
        Q = deque([(node, head)])
        visited[head.val]=head

        while len(Q)>0:
            curr, cpy = Q.popleft()
            for n in curr.neighbors:
                if n.val in visited:
                    cpy.neighbors.append(visited[n.val])
                else:
                    nCpy = Node(n.val)
                    cpy.neighbors.append(nCpy)
                    visited[n.val]=nCpy
                    Q.append((n, nCpy))
        return head