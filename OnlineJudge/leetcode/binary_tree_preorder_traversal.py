__author__ = 'lxzMac'
# https://leetcode.com/problems/binary-tree-preorder-traversal/
#
# Given a binary tree, return the preorder traversal of its nodes' values.
#
# For example:
# Given binary tree {1,#,2,3},
#    1
#     \
#      2
#     /
#    3
# return [1,2,3]



from operator import concat
class Solution:
    # @param root, a tree node
    # @return a list of integers
    def preorderTraversal(self, root):#42 ms
        if root == None:
            return []

        left_result = self.preorderTraversal(root.left)
        right_result = self.preorderTraversal(root.right)

        return concat(concat([root.val], left_result), right_result )



from OnlineJudge.leetcode import tree_node


str = """{1,4,3,2}"""
root = tree_node.getTreeNodeRoot(str)
print Solution().preorderTraversal(root)