# class Vertical:
#     def __init__(self):
#         R, C = len(matrix), len(matrix[0])
#         #self.matrix = [[0]*C for _ in range(R)]
#         for j in range(C):
#             accum = 0
#             for i in range(R):
#                 accum+=self.matrix[i][j]
#                 self.matrix[i][j]=accum
#     def __getitem__(self, i:int)->list[int]:
#         return self.matrix[i]


class NumMatrix:
    def __init__(self, matrix: List[List[int]]):
        self.matrix = matrix
        R, C = len(matrix), len(matrix[0])
        #self.prefix = [[0]*C for _ in range(R)]
        for i in range(R):
            accum = 0
            for j in range(C):
                accum += matrix[i][j]
                #print(f"accum:{accum}")
                self.matrix[i][j]=accum
            #print(self.prefix[i])
        self.Vertical()
       # self.print(self.prefix)
       # self.print(self.vertical.matrix)
    def Vertical(self):
        R, C = len(self.matrix), len(self.matrix[0])
        #self.matrix = [[0]*C for _ in range(R)]
        for j in range(C):
            accum = 0
            for i in range(R):
                accum+=self.matrix[i][j]
                self.matrix[i][j]=accum
    def __getitem__(self, i:int)->list[int]:
        return self.matrix[i]
    
    def print(self, M:list[list[int]]):
        #print(M)
        # R, C = len(M), len(M[0])
        # for i in range(R):
        #     print(M[i])
        for row in M:
            print(row)
        print()

    def sumRegion(self, row1: int, col1: int, row2: int, col2: int) -> int:
        compensate = self[row1-1][col1-1] if row1>0 and col1>0 else 0
        upperSum = self[row1-1][col2] if row1>0 else 0
        leftSum = self[row2][col1-1] if col1>0 else 0
        return self[row2][col2]-upperSum-leftSum+compensate



# Your NumMatrix object will be instantiated and called as such:
# obj = NumMatrix(matrix)
# param_1 = obj.sumRegion(row1,col1,row2,col2)