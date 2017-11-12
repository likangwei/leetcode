package main
import "fmt"
/*
https://leetcode.com/problems/search-for-a-range/description/

我的解决过程：
  1. solution1 根据二分查找法，找到中间值，再左右滑动，找到目标
  2. solution2 提交后，消耗了32ms, 排名靠后，就又调整了下，耗时19ms,提升了13ms
  3. solution3 提交之后，还是19ms，也是醉了

高手答案: 

func searchRange(nums []int, target int) []int {
	first := -1
	last := -1
	for i, num := range nums {
		if num == target {
			if first == -1 {
				first = i
			}
			last = i
		}
	}
	if first != -1 && last == -1 {
		last = first
	}
	return []int{first, last}
}

当前水平改进：
   高手一个for range 就能达到19ms。。。
   1. 简单的代码最有效，最好是，自带的这种 for range， ==之类的，那是相当高效的
   2. 后续写项目的时候记住，再没发现瓶颈之前，不要过早优化，用最简单的代码实现。。。

1. vs高手
    a) 命名：高手用first, last来标记最前面的和最后面的， 而我用的是i, j不如高手精准
    b) 行数：16 < 34
    c) 思路：高手直接用for循环来达到目标，而且算法还比我的二分查找+两边扩展快，也是醉了，说明for range真的是非常快
    d) 技巧：
    e) 此题感悟： 当真是过早优化是万恶之源啊

2. 此题感悟
    
*/

func bsearch(nums[]int, target int) int{
	if len(nums) == 0{
		return -1
	}
	i, j := 0, len(nums)
	for i < j{
		mid := (i+j)/2
		if target > nums[mid]{
			i = mid + 1
		}else if target < nums[mid]{
			j = mid
		}else{
			return mid
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	rst := []int{-1, -1}
	mid := bsearch(nums, target)
	if mid != -1{
		left, right := mid, mid
		for left-1>=0 && nums[left-1] == target{
			left --
		}
		for right+1 <= len(nums)-1 && nums[right+1] == target{
			right ++
		}
		rst[0], rst[1] = left, right
	}
	return rst
}

func searchRange2(nums []int, target int) []int {
	rst := []int{-1, -1}

	for i,j:=0, len(nums); i < j;{
		mid := (i+j) / 2
		n := nums[mid]
		if target < n{
			j = mid
		}else if target == n{
			if mid == 0 || nums[mid-1] != target{
				rst[0] = mid
				break
			}
			j = mid
		}else{
			i = mid + 1
		}
	}
	for i,j:=0, len(nums); i < j;{
		mid := (i+j) / 2
		n := nums[mid]
		if target > n{
			i = mid + 1
		}else if target == n{
			if mid == len(nums)-1 || nums[mid+1] != target{
				rst[1] = mid
				break
			}
			i = mid + 1
		}else{
			j = mid
		}
	}
	return rst
}


func searchRange3(nums []int, target int) []int {
	rst := []int{-1, -1}

	rl, rr := 0, len(nums)

	for rl < rr{
		mid := (rl+rr)/2
		n := nums[mid]
		if target > n{
			rl = mid+1
		}else if target < n{
			rr = mid
		}else{
			break
		}
	}

	for i,j:=rl, rr; i < j;{
		mid := (i+j) / 2
		n := nums[mid]
		if target < n{
			j = mid
		}else if target == n{
			if mid == 0 || nums[mid-1] != target{
				rst[0] = mid
				break
			}
			j = mid
		}else{
			i = mid + 1
		}
	}
	for i,j:=rl, rr; i < j;{
		mid := (i+j) / 2
		n := nums[mid]
		if target > n{
			i = mid + 1
		}else if target == n{
			if mid == len(nums)-1 || nums[mid+1] != target{
				rst[1] = mid
				break
			}
			i = mid + 1
		}else{
			j = mid
		}
	}
	return rst
}

func main() {
	to_test := [][]int{
		[]int{5, 7, 7, 8, 8, 10},
	}

	to_test2 := []int{
		8,
	}

	_ = [][]string{
		[]string{"foo", "bar"},
		[]string{"dhvf","sind","ffsl","yekr","zwzq","kpeo","cila","tfty","modg","ztjg","ybty","heqg","cpwo","gdcj","lnle","sefg","vimw","bxcb"},
	}

	for i:=0; i < len(to_test); i+=1{
		p1, p2 := to_test[i], to_test2[i]
		rst := searchRange3(p1, p2)
		fmt.Println(p1, p2, rst)
	}
}
