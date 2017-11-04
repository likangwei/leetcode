package main
import "fmt"
/*
https://leetcode.com/problems/longest-valid-parentheses/description/

我的解决过程：
  1. 开始的时候，想到了一个比较慢的算法，但是总是感觉太慢


高手答案: 效率跟我一样

当前水平改进：
  1. 用各种数据结构 标记、打TAG， 比较初级的的就是一个 []bool ，后期可以用 LinkList, Heap, Stack, Tree, Dag
*/

func longestValidParentheses(s string) int {
	str_len := len(s)
	bool_lst := make([]bool, str_len)
	lft_lst := make([]int, str_len)
	lst_len := 0
	for i:=0; i<str_len; i++{
		b := s[i]
		if b == '('{
			lft_lst[lst_len] = i
			lst_len += 1
		}else if b == ')'{
			if lst_len > 0{
				bool_lst[i], bool_lst[lft_lst[lst_len-1]] = true, true
				lst_len --
			}else{
				lst_len = 0
			}
		}
	}
	// fmt.Println(bool_lst)
	longestNum, curLoop := 0, 0
	for i:=0; i<str_len; i++{
		if bool_lst[i]{
			curLoop++
			if curLoop > longestNum{
				longestNum = curLoop
			}
		}else{
			curLoop = 0
		}
	}
	return longestNum
}

func main() {
	// to_test := [][]int{
	// 	[]int{1, 2, 3},
	// 	[]int{3,2,1},
	// 	[]int{1, 1, 5},
	// 	[]int{1,3, 2},
	// }

	// _ = [][]string{
	// 	[]string{"foo", "bar"},
	// 	[]string{"dhvf","sind","ffsl","yekr","zwzq","kpeo","cila","tfty","modg","ztjg","ybty","heqg","cpwo","gdcj","lnle","sefg","vimw","bxcb"},
	// }

	to_test := []string{
		"(())",
		"())",
		"(()",
		")()())",
	}

	for i:=0; i < len(to_test); i+=1{
		rst :=longestValidParentheses(to_test[i])
		fmt.Println(to_test[i], rst)
	}
}
