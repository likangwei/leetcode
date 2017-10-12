__author__ = 'likangwei'
class Solution(object):
    def trailingZeroes(self, n):
        """
        :type n: int
        :rtype: int
        """

        if n == 0:
            return 0
        n = reduce(lambda x,y:x*y,[i for i in xrange(1,n+1)])
        rst = 0
        while n % 10 == 0:
            rst += 1
            n = n / 10
        return rst

    def trailingZeroes2(self, n):
        """
        :type n: int
        :rtype: int
        """

        if n == 0:
            return 0
        nums = [i for i in xrange(1, n+1)]
        rst = 0
        for idx, n in enumerate(nums):
            while n % 10 == 0:
                rst += 1
                n = n / 10
            nums[idx] = n

        n = reduce(lambda x, y: x*y, nums)

        while n % 10 == 0:
            rst += 1
            n = n / 10
        return rst

    def trailingZeroes3(self, n):
        """
        :type n: int
        :rtype: int
        """

        if n == 0:
            return 0
        n = str(reduce(lambda x, y: x*y, [i for i in xrange(1, n+1)]))
        rst = 0
        while n.endswith('0'):
            rst += 1
            n = n[:-1]
        return rst

    def trailingZeroes4(self, n):
        """
        :type n: int
        :rtype: int
        """
        return 0 if n == 0 else n / 5 + self.trailingZeroes4(n / 5)


import time
s = Solution()
for m in [s.trailingZeroes, s.trailingZeroes2, s.trailingZeroes3, s.trailingZeroes4]:
    now = time.time()
    for i in [1000, 5000, 10000]:
        print i, m(i)
    print time.time() - now, '-' * 10