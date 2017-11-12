/*
https://leetcode.com/problems/implement-strstr/description/
高手答案 xms < xms

*/

package main
import "fmt"
import "sort"
/*
https://leetcode.com/problems/substring-with-concatenation-of-all-words/description/

我的解决过程：
  1. 思路上，有思维盲点， 根据testcase查找出来了，手动debug了自己的思路算法，但是因为长度太短，所以没有发现盲点的地方


高手答案: 我的答案跟高手答案一样

当前水平改进：
   在实现代码前，要多用各种testcase来验证自己的算法，不要因为感觉简单，上来就写代码，最好是在思路上就一次弄对
   根据多个维度，来造testcase  
   
*/
func nextPermutation(nums []int)  {
	l := len(nums)
	if l <= 1{
		return
	}
	for i:=l-2; i>=0; i--{
		if nums[i] < nums[i+1]{
			for j:=l-1; j>=i+1; j--{
				if nums[j] > nums[i]{
					nums[i], nums[j] = nums[j], nums[i]
					sort.Ints(nums[i+1:])
					return
				}
			}
		}
	}
	sort.Ints(nums)
}


func main() {
	to_test := [][]int{
		[]int{1, 2, 3},
		[]int{3,2,1},
		[]int{1, 1, 5},
		[]int{1,3, 2},
	}

	_ = [][]string{
		[]string{"foo", "bar"},
		[]string{"dhvf","sind","ffsl","yekr","zwzq","kpeo","cila","tfty","modg","ztjg","ybty","heqg","cpwo","gdcj","lnle","sefg","vimw","bxcb"},
	}

	for i:=0; i < len(to_test); i+=1{
		nextPermutation(to_test[i])
		fmt.Println(to_test[i])
	}
}
