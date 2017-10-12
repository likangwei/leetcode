__author__ = 'lxzMac'
# https://leetcode.com/problems/excel-sheet-column-number/
# Given a column title as appear in an Excel sheet, return its corresponding column number.
#
# For example:
#
#     A -> 1
#     B -> 2
#     C -> 3
#     ...
#     Z -> 26
#     AA -> 27
#     AB -> 28
# Credits:
# Special thanks to @ts for adding this problem and creating all test cases.
class Solution:
    # @param s, a string
    # @return an integer
    def titleToNumber(self, s):

        result = 0
        for i in range(len(s)):
            result = result * 26 + (ord(s[i])-64)

        return result

print Solution().titleToNumber("AAA")