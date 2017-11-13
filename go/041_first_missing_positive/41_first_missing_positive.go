/*
https://leetcode.com/problems/first-missing-positive/description/
*/

func firstMissingPositive(nums []int) int {
    if len(nums) == 0{
        return 1
    }
    sort.Ints(nums)
    start := sort.SearchInts(nums, 1)
    j := 1
    for i := start ; i<len(nums); {
        if j != nums[i]{
            return j
        }
        for i<len(nums) && j == nums[i]{
            i ++
        }
        j += 1
    }
    return j
}