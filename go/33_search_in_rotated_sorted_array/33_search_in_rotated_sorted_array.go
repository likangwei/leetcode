package main
import "fmt"
/*
https://leetcode.com/problems/longest-valid-parentheses/description/

高手答案: 

我的解决过程：

当前水平改进：

*/

func lazySearch(nums[]int, target, f, e int) int {
	if f >= e{
		return -1
	}
	for i:=f; i<e; i++{
			if nums[i] == target{
				return i
			}
	}
	return -1
}

func bsearch(nums[]int , target, f, e int) int{
	if nums[f] > target && nums[e-1] < target{
		return -1
	}
	if e - f <= 10{
		return lazySearch(nums, target, f, e)
	}
	mid := f + ((e-f)/2)
	if target >= nums[mid] {
		return bsearch(nums, target, mid, e)
	}else{
		return bsearch(nums, target ,f, mid)
	}
	return -1
}

func xsearch(nums[]int , target, f, e int) int{
	if f - e <= 10 {
		return lazySearch(nums, target, f, e)
	}

	mid := f + ((e - f) / 2)
	left_is_sort := nums[mid] > nums[f]

	if left_is_sort{
		if target >= nums[f] && target <= nums[e-1]{
			return bsearch(nums, target ,f ,mid)
		}else{
			return xsearch(nums, target, mid, e)
		}
		
	}else{
		if target >= nums[mid] && target <= nums[e-1]{
			return bsearch(nums, target , mid, e)
		}else{
			return xsearch(nums, target, f, mid)
		}
	}
}

func search(nums []int, target int) int {
	if len(nums) == 0{
		return -1
	}
	return xsearch(nums, target, 0, len(nums))
}

func main() {
	to_test := [][]int{
		[]int{4, 5, 6, 7, 0, 1, 2},
	}

	// _ = [][]string{
	// 	[]string{"foo", "bar"},
	// 	[]string{"dhvf","sind","ffsl","yekr","zwzq","kpeo","cila","tfty","modg","ztjg","ybty","heqg","cpwo","gdcj","lnle","sefg","vimw","bxcb"},
	// }

	// to_test := []string{
	// 	"(())",
	// 	"())",
	// 	"(()",
	// 	")()())",
	// }
	to_test2 := []int{3}

	for i:=0; i < len(to_test); i+=1{
		rst := search(to_test[i], to_test2[i])
		fmt.Println(to_test[i], to_test2[i], rst)
	}
}
