# https://leetcode.com/problems/count-and-say

# The count-and-say sequence is the sequence of integers beginning as follows:
# 1, 11, 21, 1211, 111221, ...
#
# 1 is read off as "one 1" or 11.
# 11 is read off as "two 1s" or 21.
# 21 is read off as "one 2, then one 1" or 1211.
# Given an integer n, generate the nth sequence.
#
# Note: The sequence of integers will be represented as a string
#
#
# .
class Solution:
    # @return a string
    def countAndSay(self, n):
        if n == 1:
            return '1'

        preStr = self.countAndSay(n-1)
        preChar = None
        preCharcount = 0
        result = ''
        charLen = len(preStr)
        for i in range(charLen):
            currentChar = preStr[i]

            is_end = False;
            if i == charLen-1:
                is_end = True

            if preChar == None:
                preChar = currentChar
                preCharcount = 1

            elif preChar <> None and preChar <> currentChar:

                result = '%s%d%s'  %(result, preCharcount, preChar)

                preCharcount = 0;
                preChar = currentChar

            if preChar <> None and preChar == currentChar:
                preCharcount = preCharcount + 1

            if is_end:
                result = '%s%d%s'  %(result, preCharcount, preChar)

        return result

s = Solution()
for i in range(1,50):
    print i, s.countAndSay(i)