# __author__ = 'likang'
# https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
# Given a binary tree
#
#     struct TreeLinkNode {
#       TreeLinkNode *left;
#       TreeLinkNode *right;
#       TreeLinkNode *next;
#     }
# Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.
#
# Initially, all next pointers are set to NULL.
#
# Note:
#
# You may only use constant extra space.
# You may assume that it is a perfect binary tree (ie, all leaves are at the same level, and every parent has two children).
# For example,
# Given the following perfect binary tree,
#          1
#        /  \
#       2    3
#      / \  / \
#     4  5  6  7
# After calling your function, the tree should look like:
#          1 -> NULL
#        /  \
#       2 -> 3 -> NULL
#      / \  / \
#     4->5->6->7 -> NULL

# Definition for binary tree with next pointer.
# class TreeLinkNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
#         self.next = None

class Solution:
    # @param root, a tree link node
    # @return nothing
    def connect(self, root):

        cur_center_list = [root]

        while cur_center_list:
            cur_trees_len = len(cur_center_list)
            next_center_trees = []
            for idx in range(cur_trees_len):
                cur_tree = cur_center_list[idx]

                if cur_tree != None:
                    if idx != cur_trees_len-1:
                        cur_tree.next = cur_center_list[idx+1]

                    next_center_trees.append(cur_tree.left)
                    next_center_trees.append(cur_tree.right)
            cur_center_list = next_center_trees
from OnlineJudge.leetcode import tree_node
str = """{0,-9,7,-1,-5,-8,9,4,3,8,-6,-6,4,-5,2,5,5,2,6,-6,2,7,-2,8,4,3,2,0,1,0,5,9,-4,5,4,4,2,-8,7,7,-8,-3,-6,2,-1,3,7,5,-2,1,3,0,6,9,6,-2,-9,7,1,1,1,-5,5,9,-4,-3,-2,-5,3,5,0,3,-2,-3,-1,-1,-3,-7,5,-5,5,8,5,4,-1,-4,-8,-5,5,-8,1,-7,-2,-4,-3,7,-8,5,5,3,9,-9,6,6,3,-1,4,8,2,-8,8,5,6,9,-5,-3,5,-7,-3,-8,-4,0,2,-6,-4,1,-8,-1,-3,3,-2,1,-1,9,-9,-6,0,-9,-1,6,-7,8,8,9,0,5,-9,3,1,-3,-5,5,-7,-1,-4,8,-4,-5,-1,2,0,-8,-3,-8,9,7,2,-5,4,-7,-4,4,0,9,1,7,5,8,-4,-4,1,-7,3,-8,6,-6,-1,-1,9,-7,3,1,9,2,-3,-4,3,9,6,-7,8,-7,-4,5,-5,-3,-6,-9,4,4,9,-3,1,-9,-7,2,-7,-3,-7,-9,-2,-1,-2,9,6,0,-6,6,-5,-1,-9,0,-7,1,1,-8,-4,-9,-8,3,8,-5,-1,7,2,2,-5,-2,-7,9,-5,-6,2,-6,-2,5,3,-5,-5,0,-3,-6,0,6,3,2,-2,3,8,8,-1,-3,7,1,8,-2,6,-1,-4,0,-2,8,2,-7,-5,-9,3,-1,-7,-3,2,-2,6,4,-3,6,-4,-2,3,7,-3,-2,6,0,-1,7,-9,-1,-7,9,-4,3,-1,-9,-8,5,8,9,5,6,8,-7,-3,-1,4,-4,6,-3,6,8,-9,-1,1,0,-1,-7,2,-8,6,-5,9,1,2,-2,8,-4,5,-6,3,1,-8,-8,5,7,-6,9,7,1,-8,-3,2,-7,3,-7,8,-5,8,4,7,-2,7,6,-4,-2,-6,-8,7,-6,-2,0,-7,1,-9,-9,5,-6,1,5,-5,-8,0,-2,9,0,0,7,2,8,5,1,-5,4,5,-2,-9,-4,0,7,-9,-2,-3,-6,-1,-7,-9,4,1,-3,0,-3,9,3,-1,3,0,9,-6,2,4,-1,-4,0,0,6,6,-6,-2,6,7,-1,-1,-1,2,-9,0,-1,5,0,-2,-8,5,4,-1,0,-2,-5,-5,-7,0,4,0,0,5,-6,4,-5,-4,7,-5,-3,-1,5,-3,5,3,-3,3,6,-2,-4,-8,-2,6,-9,7,-1,-5,5,2,-2,5,-1,3,6,-8,0,-5,4,7,8,-7,-6,-3,0,-7,5,8,-3,1,9,3}"""
root = tree_node.getTreeNodeRoot(str)
print 0
print Solution().connect(root)