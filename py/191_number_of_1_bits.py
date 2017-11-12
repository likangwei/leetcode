__author__ = 'lxzMac'
# https://leetcode.com/problems/number-of-1-bits/
# Write a function that takes an unsigned integer and returns the number of 1 bits it has (also known as the Hamming weight).
#
# For example, the 32-bit integer 11 has binary representation 00000000000000000000000000001011, so the function should return 3.
#
# Credits:
# Special thanks to @ts for adding this problem and creating all test cases.
#

class Solution:
    # @param n, an integer
    # @return an integer

    bin_map = {1:1}

    def hammingWeight(self, n):

        need_go_back_point = 0
        for i in range(32):
            if n >= self.get_num_from_idx(i+1):
                need_go_back_point = i+1
            else:
                break

        result = 0
        while n>0:
            if n >= self.get_num_from_idx(need_go_back_point):
                result = result + 1
                n = n - self.get_num_from_idx(need_go_back_point)

            need_go_back_point = need_go_back_point -1

        return result

    def get_num_from_idx(self, n):
        if self.bin_map.has_key(n):
            return self.bin_map[n]

        result = self.get_num_from_idx(n-1) * 2
        self.bin_map[n] = result
        return result


print Solution().hammingWeight(1024*1024*1024*4-10)
