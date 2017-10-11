__author__ = 'likang'
# https://leetcode.com/problems/n-queens-ii/
# Follow up for N-Queens problem.
#
# Now, instead outputting board configurations, return the total number of distinct solutions.
#
# Show Tags
# Have you met this question in a real interview


class Solution:
    # @return an integer
    def totalNQueens(self, n):

        data = [[] for i in range(n)]

        for i in range(n):
            for j in range(n):
                data[i].append(1)

        return self.putZero(data, 0, n)

    def putZero(self, data, x, n):

        ifHas = False
        for i in data[x]:
            if i==1:
                ifHas = True
                break
        if not ifHas:
            return 0
        # the last line
        if x == n-1 :
            #print 'the last line', data[x]
            result = 0
            for y in data[x]:
                if y == 1:
                    # print '-'*20
                    # for tmp in data:
                    #     print tmp
                    # print '='*20
                    result = result + 1
            return result

        # current line has 1

        result = 0
        for i in range(n):
            if data[x][i] == 1:
                tmpData = self.copyOfData(data)
                tmpData = self.putZero2(tmpData, x, i, n)
                result = result + self.putZero(tmpData, x+1, n)

        return result

    def copyOfData(self, data):
        n = len(data)
        tmpData = [[] for i in range(n)]
        for i in range(n):
            for j in range(n):
                tmpData[i].append(data[i][j])

        return tmpData

    def putZero2(selfself, data, x, y, n):

        #shu
        for i in range(x,n):
            data[i][y] = 0

        # heng
        for i in range(n):
            data[x][i] = 0

        # print 'after heng shu', data

        nextLineX = x
        nextLineY = y
        while nextLineX<n and nextLineY<n:
            data[nextLineX][nextLineY] = 0
            nextLineX = nextLineX + 1
            nextLineY = nextLineY + 1
        # print 'after x1', data
        nextLineX = x
        nextLineY = y

        while nextLineX<n and nextLineY>=0:
            data[nextLineX][nextLineY] = 0
            nextLineX = nextLineX + 1
            nextLineY = nextLineY - 1
        data[x][y] = 1
        return data

n = 8

data = [[] for i in range(n)]
for i in range(n):
    for j in range(n):
        data[i].append(1)
# print data

# xy = [[0,3],[1,6],[2,2],[3,7],[4,1],[5,4],[6,0]]
# for x,y in xy:
#     data =  Solution().putZero2(data, x, y, n)
#     print data
# print Solution().putZero(data,7, n)
for i in range(2,100):
    print i, Solution().totalNQueens(i);