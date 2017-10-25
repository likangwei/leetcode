/*
https://leetcode.com/problems/4sum/description/

高手答案
func fourSum(num sort.IntSlice, target int) [][]int {
    var ans [][]int
	length := len(num)
	if length < 4 {
		return ans
	}
	sort.Sort(num)
	for i := 0; i < length-3; i++ {
		if num[i]+num[i+1]+num[i+2]+num[i+3] > target {
			break
		}
		if num[i]+num[length-1]+num[length-2]+num[length-3] < target {
			continue
		}
		if i > 0 && num[i] == num[i-1] { // 去掉重复的
			continue
		}

		for j := i + 1; j < length-2; j++ {
			if num[i]+num[j]+num[j+1]+num[j+2] > target {
				break
			}
			// 每次比较前两个和后两个
			if num[i]+num[j]+num[length-1]+num[length-2] < target {
				continue
			}
			if j > i+1 && num[j] == num[j-1] {
				continue
			}
			low, high := j+1, length-1
			for low < high {
				sum := num[i] + num[j] + num[low] + num[high]
				if sum == target {
					ans = append(ans, []int{num[i], num[j], num[low], num[high]})
					for low < high && num[low] == num[low+1] {
						low++
					}
					for low < high && num[high] == num[high-1] {
						high--
					}
					low++
					high--
				} else if sum < target {
					low++
				} else {
					high--
				}
			}
		}
	}
	return ans
}


总结：
* 排序后的数组=算盘，index=算盘珠
* 排序后的数组，可以用一些技巧快速跳过index, 比如 nums[n] == nums[n+1], 或者 sum(nums[~]) < 最小值
* 如果一个list里面有两个index来控制，可以右面的index拖到最右
*/

package main
import "fmt"
import "sort"

func set(nums [][]int) [][]int {
	var rst [][]int = make([][]int, 0, 1)
	m := make(map[string][]int)
	for _ , lst := range nums{
		s := fmt.Sprint(lst)
		m[s] = lst
	}
	for _, lst := range m{
		rst = append(rst, lst)
	}
	return rst
}

func fourSum(nums []int, target int) [][]int {
	var rst [][]int = make([][]int, 0, 10)
	sort.Ints(nums)
	for i:=0; i <= len(nums) - 4; i++{
		if nums[i] + nums[i+1] + nums[i+2] + nums[i+3] > target{
			break
		}
		if i+4 < len(nums) && nums[i] == nums[i+4]{
			continue
		}
		if nums[i] + nums[len(nums)-3] + nums[len(nums)-2] + nums[len(nums)-1] < target{
			continue
		}

		for j:=i+1; j <= len(nums) - 3; j++{
			if nums[j] + nums[j+1] + nums[j+2] > target - nums[i]{
				break
			}
			if j+3 < len(nums) && nums[j] == nums[j+3]{
				continue
			}
			if nums[j] + nums[len(nums)-2] + nums[len(nums)-1] < target - nums[i]{
				continue
			}

			k, l := j+1, len(nums) - 1
			for k < l{
				a, b, c, d := nums[i], nums[j], nums[k], nums[l]
				sum := a + b + c +d
				if sum > target{
					l -= 1
					for l > k{
						if nums[l] == nums[l+1]{
							l--
							continue
						}
						break
					}
					
				}else if sum < target{
					k ++
					for k < l{
						if nums[k] == nums[k-1]{
							k ++ 
						}
						break
					}
				}else{
					l -= 1
					rst = append(rst, []int{a, b, c, d})
				}
				
			}
		}
	}
	return set(rst)
}

func main() {
	to_test := [][]int{
		// []int{1, 0, -1, 0, -2, 2},
		// []int{0, 0, 0, 0},
		// []int{-1,2,2,-5,0,-1,4},
		// []int{-3,-2,-1,0,0,1,2,3},
		// []int{-4,-3,-2,-1,0,0,1,2,3,4},
		[]int{-1,0,-5,-2,-2,-4,0,1,-2},
	}
	to_test2 := []int{
		// 0, 
		// 1, 
		-9, 
		// 0,
	}

	for i := 0; i < len(to_test); i=i+1{
		rst := fourSum(to_test[i], to_test2[i])
		fmt.Println(">>>", to_test[i], "rst:", rst, len(rst))
	}
	
}
