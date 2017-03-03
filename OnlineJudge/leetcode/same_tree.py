__author__ = 'lxzMac'
# https://leetcode.com/problems/same-tree/
# Given two binary trees, write a function to check if they are equal or not.
#
# Two binary trees are considered equal if they are structurally identical and the nodes have the same value.


class Solution:
    # @param p, a tree node
    # @param q, a tree node
    # @return a boolean
    def isSameTree(self, p, q):
        left_container = [p]
        right_container = [q]
        while left_container:

            left_root = left_container[0]
            right_root = right_container[0]
            if left_root!=None and right_root!=None :
                if left_root.val==right_root.val:
                    del left_container[0]
                    del right_container[0]
                    left_container.append(left_root.left)
                    left_container.append(left_root.right)
                    right_container.append(right_root.left)
                    right_container.append(right_root.right)
                else:
                    return False
            elif left_root!=right_root and (left_root==None or right_root==None):
                return False
            else:
                del left_container[0]
                del right_container[0]

        return True