__author__ = 'likangwei'

# https://leetcode.com/problems/roman-to-integer/
# Given a roman numeral, convert it to an integer.
#
# Input is guaranteed to be within the range from 1 to 3999.

class Solution:
    # @param {string} s
    # @return {integer}
    def romanToInt(self, s):
        x = {'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
        result = 0
        s_len = len(s)
        has_handle = [False] * len(s)
        for idx, c in enumerate(s):
            if has_handle[idx]:
                continue

            cur_n = x[c]
            next_n = 0 if (idx == (s_len-1)) else x[s[idx+1]]
            if cur_n >= next_n:
                result = result + cur_n
            else:
                result = result + next_n - cur_n
                has_handle[idx+1] = True

        return result



print Solution().romanToInt("MCMXCVI")#1996
print Solution().romanToInt("DCXXI")#621