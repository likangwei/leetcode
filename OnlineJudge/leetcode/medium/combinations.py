__author__ = 'likangwei'
from OnlineJudge.leetcode.decorator import how_quick


class Solution(object):

    def combine(self, n, k, l=None):
        """
        :type n: int
        :type k: int
        :rtype: List[List[int]]
        """
        if l is None:
            l = [x+1 for x in range(n)]

        if k == 1:
            return [[x] for x in l]
        rst = []
        idx = 0
        for num in l[:-1]:
            for t in self.combine(n, k-1, l=l[idx+1:]):
                add_tmp = [num]
                add_tmp.extend(t)
                rst.append(add_tmp)
            idx += 1
        return rst


s = Solution()
rst = s.combine(10, 7)
print rst