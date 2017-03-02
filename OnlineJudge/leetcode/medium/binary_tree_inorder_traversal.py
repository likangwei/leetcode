__author__ = 'lxzMac'
# https://leetcode.com/problems/binary-tree-inorder-traversal/
# Given a binary tree, return the inorder traversal of its nodes' values.
#
# For example:
# Given binary tree {1,#,2,3},
#    1
#     \
#      2
#     /
#    3
# return [1,3,2].

# Definition for a  binary tree node
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
from operator import concat

class Solution:
    # @param root, a tree node
    # @return a list of integers
    def inorderTraversal(self, root):

        if root == None:
            return []

        left_val = self.inorderTraversal(root.left)
        right_val = self.inorderTraversal(root.right)
        result = concat(concat(left_val, [root.val]), right_val)
        return result


from OnlineJudge.leetcode import tree_node
str = """{1,#,2,3}"""
root = tree_node.getTreeNodeRoot(str)
print Solution().inorderTraversal(root)