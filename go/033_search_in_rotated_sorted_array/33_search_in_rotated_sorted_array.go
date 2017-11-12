package main
import "fmt"
/*
https://leetcode.com/problems/search-in-rotated-sorted-array/description/

高手答案: 
func search(nums []int, target int) int {
    lo, hi := 0, len(nums) - 1
    
    for lo < hi {
        mid := (lo + hi) / 2
        if nums[mid] > nums[hi] {lo = mid + 1} else {hi = mid}
    }
    
    idx := -1
    
    if lo > 0 && target >= nums[0] && target <= nums[lo - 1] {
        idx = sort.SearchInts(nums[0:lo], target)
    } else if len(nums) > 0 && target >= nums[lo] && target <= nums[len(nums) - 1] {
        idx = sort.SearchInts(nums[lo:], target) + lo
    }
    
    if idx >= 0 && idx < len(nums) && nums[idx] == target {
        return idx
    }
    
    return -1
}

高手比我强的地方：
  命名： 高手用 lo, hi来表示高低点，我用s, e来表示
  行数： 22 < 43
  速度： 9ms == 9ms
  思路：高手使用了更清晰思路。先找到中间的点，再分别进行二分查找，而我是之间二分查找和找中间的点同时
       进行的，写起来复杂度更高，更不容易理解
  比我多懂的： sort.SearchInts 是二分查找法， 这是我之前不知道的

自己跟自己比：
  1. 我在第一时间也是想到了跟高手一样的思路，但是没有执行，触动了我的一个老犯的毛病： 过度提前优化

当前水平改进：
  1. 解题时，多用思路清晰的代码，清晰是第一位的，这样写出来的代码更稳定
  2. 自己在库的使用上，比如sort不熟，或者没有想到google golang的二分查找法出来，所以自
     己会造很重要，但是如果有现成的，还是现成的快
  3. 

此题目特点：
 用最简单的算法，也可以达到12ms的效果, 更加确定了提前优化是万恶之源
# 12ms
func search(nums []int, target int) int {
    for idx, n := range nums{
        if n == target{
            return idx
        }
        
    }
    return -1
}

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
