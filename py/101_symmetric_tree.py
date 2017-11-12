# Definition for a  binary tree node
# https://leetcode.com/problems/symmetric-tree/
# #Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
#
# For example, this binary tree is symmetric:
#
#     1
#    / \
#   2   2
#  / \ / \
# 3  4 4  3
# But the following is not:
#     1
#    / \
#   2   2
#    \   \
#    3    3


class TreeNode:
    hasInitLeft = False
    hasInitRight = False
    hasInitVal = False
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    # @param root, a tree node
    # @return a boolean
    def isSymmetric(self, root):
        if root == None:
            return True

        leftKuang = []
        rightKuang = []
        leftKuang.append(root.left)
        rightKuang.append(root.right)

        while len(leftKuang) != 0 :
            leftRoot = leftKuang[0]
            rightRoot = rightKuang[0]
            if(leftRoot == None and leftRoot == rightRoot):
                del leftKuang[0]
                del rightKuang[0]
            elif leftRoot == None or rightRoot == None:
                return False
            elif leftRoot.val == rightRoot.val:
                leftKuang.remove(leftRoot)
                rightKuang.remove(rightRoot)
                leftKuang.append(leftRoot.left)
                leftKuang.append(leftRoot.right)
                rightKuang.append(rightRoot.right)
                rightKuang.append(rightRoot.left)
            else:
                return False
        return True


test = """{1,2,2,#,3,3}"""
from OnlineJudge.leetcode import tree_node
root = tree_node.getTreeNodeRoot(test)
print Solution().isSymmetric(root)