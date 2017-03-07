__author__ = 'lxzMac'
# https://leetcode.com/problems/linked-list-cycle/
# Given a linked list, determine if it has a cycle in it.
#
# Follow up:
# Can you solve it without using extra space?

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:

    # @param head, a ListNode
    # @return a boolean
    def hasCycle(self, head):

        if head == None:
            return False

        while head.next != None:

            if head.next == head:
                return True

            next = head.next
            head.next = head
            head = next

        return False


