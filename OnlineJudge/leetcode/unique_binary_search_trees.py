__author__ = 'lxzMac'
# https://leetcode.com/problems/unique-binary-search-trees/
# Given n, how many structurally unique BST's (binary search trees) that store values 1...n?
#
# For example,
# Given n = 3, there are a total of 5 unique BST's.
#
#    1         3     3      2      1
#     \       /     /      / \      \
#      3     2     1      1   3      2
#     /     /       \                 \
#    2     1         2                 3



class Solution:
    # @return an integer
    allStatus = {}
    allNumTrees = {}
    def __init__(self):
        self.allNumTrees[0] = 1
        self.allNumTrees[1] = 1

    def numTrees(self, n):

        if self.allNumTrees.has_key(n):
            return self.allNumTrees[n]

        result = 0
        for leftSize,rightSize in self.getAllStatus(n-1):
            result = result + self.numTrees(leftSize) * self.numTrees(rightSize)
        self.allNumTrees[n] = result
        return result

    def getAllStatus(self, n):
        if self.allStatus.has_key(n):
            return self.allStatus[n]
        result = set()
        for i in range(n):
            left = n-i;
            right = n-left
            result.add((left, right))
            result.add((right, left))
        self.allStatus[n] = result
        return result


# print Solution().getAllStatus(19)
print Solution().numTrees(8)