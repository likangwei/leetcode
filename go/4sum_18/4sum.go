

/*
https://leetcode.com/problems/4sum/description/

总结：

后续优化：

*/

package main
import "fmt"

func combinations(nums[]int , count int, fromIdx int)[][]int{
	var rst [][]int = make([][]int, 0, 1)
	if len(nums) < (fromIdx + count){
		return rst
	}
	if count == 0 || len(nums) == (fromIdx + count){
		return [][]int{nums[fromIdx:]}
	}
	if count == 1{
		for i:=fromIdx; i<=len(nums); i++{
			rst = append(rst, []int{nums[i]})
		}
		return rst
	}

	for i:=0; i<=len(nums)-count; i++{
		cur_num := nums[i]
		combs := combinations(nums, count-1, i+1)
		for j:=0; j<len(combs); j++{
			tmp_lst := combs[j]
			rst = append(rst, append(tmp_lst, cur_num))
		}
	}
	fmt.Println(">>>", nums, count, fromIdx, rst)
	return rst
}

func fourSum(nums []int, target int) [][]int {
	var rst [][]int = make([][]int, 0, 10)
	combs := combinations(nums, 4, 0)
	for i:=0; i<len(combs); i++{
		cur_total := 0
		cur_lst := combs[i]
		for j:=0; j<len(cur_lst);j++{
			cur_total += cur_lst[j]
		}
		if cur_total == 0{
			rst = append(rst, cur_lst)
		}
	}
	return rst
}

func main() {
	to_test := [][]int{[]int{1, 0, -1, 0, -2, 2}}
	for i := 0; i < len(to_test); i=i+1{
		fmt.Println(">>>", to_test[i], "rst:", fourSum(to_test[i], 0))
	}
	
}
