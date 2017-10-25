

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
	return rst
}

func threeSumClosest(nums []int, target int) int {

	var cj int = 9223372036854775807
	var rst int = 0
	combs := combinations(nums, 3, 0)
	for i:=0; i<len(combs); i++{
		cur_total := 0
		cur_lst := combs[i]
		for j:=0; j<len(cur_lst);j++{
			cur_total += cur_lst[j]
		}
		if cur_total == target{
			return cur_total
		}else if cur_total > target{
			if cur_total-target < cj{
				rst = cur_total
				cj = cur_total - target
			}
			
		}else if (target-cur_total) < cj{
			rst = cur_total
			cj = target - cur_total
		}
	}
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


func main() {
	to_test := [][]int{
		[]int{-1, 2, 1, -4},
		[]int{1, 1, 1, 1},
		[]int{0, 2, 1, -3},
		[]int{-11,-2,17,-16,1,-5,-5,-5,-20,7,10,-2,3,-7,-17,-13,-19,-15,-8,-7,6,-6,-8,-4,12,-12,9,-17,-13,4,-5,-15,-9,-18,-17,1,-15,-8,14,8,20,-3,-11,17,-18,10,-16,5,-9,-18,2,-3,4,-18,2,20,0,-6,18,-12,0,-17,3,-19,-20,15,12,-17,-7,8,16,7,-5,5,-13,16,-18,-7,-9,-8,-17,6,-18,0,-15,10,-13,7,9,20,7,-13,3,0,0,19,8,0,-5,-9,6,8,16,14,3,-4,5,9,-12,-19,16,6},
		// []int{0, 0, 0, 0},
		// []int{-3,-2,-1,0,0,1,2,3},
		// []int{-4,-3,-2,-1,0,0,1,2,3,4},
	}
	to_test2 := []int{
		1, 
		100, 
		1,
		-48, 
		// 0,
	}
	for i := 0; i < len(to_test); i=i+1{
		rst := threeSumClosest(to_test[i], to_test2[i])
		fmt.Println(">>>", to_test[i], "rst:", rst)
	}
	
}
