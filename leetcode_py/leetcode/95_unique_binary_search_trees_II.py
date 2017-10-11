# coding=utf8
# https://leetcode.com/problems/unique-binary-search-trees-ii/?tab=Description
# Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1...n.
# For example,
# Given n = 3, your program should return all 5 unique BST's shown below.
#    1         3     3      2      1
#     \       /     /      / \      \
#      3     2     1      1   3      2
#     /     /       \                 \
#    2     1         2                 3
# Definition for a binary tree node.
# public class Solution {
#     public List<TreeNode> generateTrees(int n) {
#         return genTrees(1,n);
#     }
#     public List<TreeNode> genTrees (int start, int end)
#     {
#         List<TreeNode> list = new ArrayList<TreeNode>();
#         if(start>end)
#         {
#             list.add(null);
#             return list;
#         }
#         if(start == end){
#             list.add(new TreeNode(start));
#             return list;
#         }
#         List<TreeNode> left,right;
#         for(int i=start;i<=end;i++)
#         {
#             left = genTrees(start, i-1);
#             right = genTrees(i+1,end);
#
#             for(TreeNode lnode: left)
#             {
#                 for(TreeNode rnode: right)
#                 {
#                     TreeNode root = new TreeNode(i);
#                     root.left = lnode;
#                     root.right = rnode;
#                     list.add(root);
#                 }
#             }
#
#         }
#         return list;
#     }
# }
# 高手答案相对比我来讲, 没有生成过多的list, 而是用数字idx控制
# 后续改进: 在用到过多list的时候, 看看是否可以用数字代替, 或者 idx代替
#
#
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def generateTrees(self, n, nums=None):
        """
        :type n: int
        :rtype: List[TreeNode]
        """
        if nums is None:
            nums = range(1, n+1)
        if n <= 0:
            return []
        if n == 1:
            return [TreeNode(nums.pop())]

        rst = []

        for head_idx in range(len(nums)):
            left_part, right_part = nums[:head_idx], nums[head_idx+1:]
            lts = self.generateTrees(len(left_part), nums=left_part) or [None]
            rts = self.generateTrees(len(right_part), right_part) or [None]

            for lt in lts:
                for rt in rts:
                    head = TreeNode(nums[head_idx])
                    head.left = lt
                    head.right = rt
                    rst.append(head)
        return rst


