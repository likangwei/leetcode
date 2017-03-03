__author__ = 'lxzMac'
# https://leetcode.com/problems/single-number/
# Given an array of integers, every element appears twice except for one. Find that single one.
from operator import contains
class Solution:
    # @param A, a list of integer
    # @return an integer

    def singleNumber(self, A): #106 ms

        dict = {}
        list_length = len(A)
        for i in range(list_length):
            currentCheck = A[i]
            if contains(dict, currentCheck):
                del dict[currentCheck]
            else:
                dict[currentCheck] = 0

        for i in dict:
            return i

    def singleNumber3(self, A): #105 ms
        result = 0;
        list_length = len(A)
        for i in range(list_length):
            result = result ^ A[i]
        return result;

    # slow....
    def singleNumber2(self, A):
        # time limit Exceeded
        currentCheckIndex = 0

        not_find = True
        while not_find:
            currentCheck = A[currentCheckIndex]
            try:
                secondIndex = A.index(currentCheck, currentCheckIndex+1)
                A[secondIndex] = A[currentCheckIndex+1]
                currentCheckIndex = currentCheckIndex + 2
            except:
                not_find = False
                return currentCheck

        return None


A = [1,2,3,1,2]

print Solution().singleNumber(A)