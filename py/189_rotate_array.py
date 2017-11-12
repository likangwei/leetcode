# https://leetcode.com/problems/rotate-array/
# Rotate an array of n elements to the right by k steps.
#
# For example, with n = 7 and k = 3, the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4].
#
# Note:
# Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
#
# [show hint]
#
# Hint:
# Could you do it in-place with O(1) extra space?
# Related problem: Reverse Words in a String II
#
# Credits:
# Special thanks to @Freezen for adding this problem and creating all test cases.

class Solution:
    # @param nums, a list of integer
    # @param k, num of steps
    # @return nothing, please modify the nums list in-place.


    def rotate(self, nums, k):

        nubs_length = len(nums)
        k = k % nubs_length
        if k == 0:
            return

        removeCount = 0;
        jumpPoint = 0;

        tmp = None

        hasAnybodyInhere = [False] * nubs_length
        jump2point = None

        jumpPoint = 0
        jump2point = jumpPoint + k

        while removeCount < nubs_length:
            # print 'jump point %d jump2point %d' % (jumpPoint, jump2point)
            if tmp == None:
                tmp = nums[jump2point]
                nums[jump2point] = nums[jumpPoint]
            else:
                tmp2 = nums[jump2point]
                nums[jump2point] = tmp
                tmp = tmp2

            hasAnybodyInhere[jump2point] = True
            removeCount = removeCount + 1



            jumpPoint = jump2point
            jump2point = jumpPoint + k

            if jump2point > nubs_length-1:
                jump2point = jump2point - nubs_length

            if hasAnybodyInhere[jump2point]:
                tmp = None
                jumpPoint = jumpPoint + 1
                jump2point = jump2point + 1

            if jump2point > nubs_length-1:
                jump2point = jump2point - nubs_length




s = Solution()
nums = [1,2,3,4,5,6]
k = 3
s.rotate(nums,k)
print nums