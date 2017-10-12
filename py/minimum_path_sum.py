#/bin/python
__author__ = 'likangwei'
class Solution(object):
    def minPathSum(self, grid, d=None, x=0, y=0):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        if d is None:
            d = {}
        m, n = len(grid), len(grid[0])
        k = x, y

        if k in d:
            return d[k]

        if x == m - 1:
            rst = sum(grid[x][y:])
            d[k] = rst
            return rst

        if y == n - 1:
            rst = sum(map(lambda t: t[y], grid[x:]))
            d[k] = rst
            return rst


        cur_num = grid[x][y]

        rr = self.minPathSum(grid, d=d, x=x, y=y+1)
        dr = self.minPathSum(grid, d=d, x=x+1, y=y)

        rst = cur_num + min(rr, dr)
        d[k] = rst
        return rst

x = [[7,1,3,5,8,9,9,2,1,9,0,8,3,1,6,6,9,5],[9,5,9,4,0,4,8,8,9,5,7,3,6,6,6,9,1,6],[8,2,9,1,3,1,9,7,2,5,3,1,2,4,8,2,8,8],[6,7,9,8,4,8,3,0,4,0,9,6,6,0,0,5,1,4],[7,1,3,1,8,8,3,1,2,1,5,0,2,1,9,1,1,4],[9,5,4,3,5,6,1,3,6,4,9,7,0,8,0,3,9,9],[1,4,2,5,8,7,7,0,0,7,1,2,1,2,7,7,7,4],[3,9,7,9,5,8,9,5,6,9,8,8,0,1,4,2,8,2],[1,5,2,2,2,5,6,3,9,3,1,7,9,6,8,6,8,3],[5,7,8,3,8,8,3,9,9,8,1,9,2,5,4,7,7,7],[2,3,2,4,8,5,1,7,2,9,5,2,4,2,9,2,8,7],[0,1,6,1,1,0,0,6,5,4,3,4,3,7,9,6,1,9]]
print len(x), len(x[0])
import time
now = time.time()

print 'solution', Solution().minPathSum(x)
print time.time() - now