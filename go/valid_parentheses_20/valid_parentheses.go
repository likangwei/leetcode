

/*
https://leetcode.com/problems/4sum/description/

总结：

后续优化：

*/

package main
import "fmt"
import "sort"


func combinations(nums[]int , count int, fromIdx int)[][]int{
	var rst [][]int = make([][]int, 0, 1)
	if len(nums) < (fromIdx + count){
		return rst
	}
	if count == 0{
		return [][]int{}
	}
	if len(nums) == (fromIdx + count){
		return [][]int{nums[fromIdx:]}
	}
	for i:=fromIdx; i<len(nums); i++{
		cur_num := nums[i]
		if count - 1 >= 1{
			combs := combinations(nums, count-1, i+1)
			for j:=0; j<len(combs); j++{
				tmp_lst := combs[j]
				rst = append(rst, append(tmp_lst, cur_num))
			}
		}else {
			rst = append(rst, []int{cur_num})
		}
	}
	// if len(rst) - len(set(rst)) > 0{
	// 	fmt.Println(nums, count, fromIdx, rst, set(rst))
	// }
	rst = set(rst)
	
	return rst
}


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
	fmt.Println("nums", nums)
	combs := combinations(nums, 4, 0)
	fmt.Println(len(combs), combs)
	for i:=0; i<len(combs); i++{
		cur_total := 0
		cur_lst := combs[i]
		for j:=0; j<len(cur_lst);j++{
			cur_total += cur_lst[j]
		}
		if cur_total == target{
			rst = append(rst, cur_lst)
		}
	}
	return rst
}

func main() {
	to_test := [][]int{
		// []int{1, 0, -1, 0, -2, 2},
		// []int{0, 0, 0, 0},
		// []int{-3,-2,-1,0,0,1,2,3},
		[]int{-4,-3,-2,-1,0,0,1,2,3,4},
	}
	to_test2 := []int{
		// 0, 
		// 1, 
		// 0, 
		0,
	}
	for i := 0; i < len(to_test); i=i+1{
		rst := fourSum(to_test[i], to_test2[i])
		fmt.Println(">>>", to_test[i], "rst:", rst, len(rst))
	}
	
}
