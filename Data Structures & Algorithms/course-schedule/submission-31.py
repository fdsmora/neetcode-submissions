#from collections import OrderedDict
from collections import deque
class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        visited = set()
        aL = self.buildList(prerequisites)

        def dfs(c: int, path: list)->bool:
            path.add(c)
            visited.add(c)
            for n in aL[c]:
                if n in path:
                    return False
                if n in visited:
                    continue
                if not dfs(n, path):
                    return False
            path.remove(c)
            return True
        
        for c in aL:
            if c in visited:
                continue
            if not dfs(c, set()):
                return False
        return True
    
    def buildList(self, prereq: list[list[int]]) -> dict:
        aL = dict()
        for c, pre in prereq:
            if pre in aL:
                aL[pre].append(c)
            else:
                aL[pre] = [c]
            if c not in aL:
                aL[c]=[]
        return aL


