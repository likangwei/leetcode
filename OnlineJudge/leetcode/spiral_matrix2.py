__author__ = 'likangwei'
class Solution(object):
    def generateMatrix(self, n):
        """
        :type n: int
        :rtype: List[List[int]]
        """
        if n == 1:
            return [[1]]

        r1 = (n+1) / 2
        nums = [x+1  for x in range(n*n)]
        idx = 0
        rst = [[0] * n] * n

        for i in xrange(r1):

            zc = n - (i*2)
            if zc > 0:
                if zc == 1:
                    r2 = 1
                else:
                    r2 = (zc-1) * 4
            else:
                break

            for i2 in range(r2):
                if r2/4 == 0:
                    curIdx = 0
                else:
                    curIdx = i2 % (r2/4)

                if i2 < r2/4:
                    x, y = i, i + curIdx
                elif i2 < r2/2:
                    x, y = i + curIdx, n - 1 - i
                elif i2 < r2 * 3 / 4:
                    x, y = n - 1 - i, n - 1 - i - curIdx
                else:
                    x, y = n - 1 - i - curIdx, i
                rst[x][y] = nums[idx]
                print x,y,idx,nums[idx], rst
                idx += 1
        return rst

print Solution().generateMatrix(2)