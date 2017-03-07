__author__ = 'likangwei'

# https://leetcode.com/problems/contains-duplicate/


class Solution:
    # @param {integer[]} nums
    # @return {boolean}
    def containsDuplicate(self, nums):
        m = {}
        for n in nums:
            if n in m:
                return True
            else:
                m[n] = None
        return False