
package main
import "fmt"
/*
https://leetcode.com/problems/valid-sudoku/description/

我的解决过程：
  1st: 算法太慢，但比较简洁
  2nd: 优化 *** 为 *
  3rd: 优化 *a ,在 bbba是将a与a做比较
  这次优点:
     直译能力增强

高手答案:
1. vs高手
	a) 命名: 
	b) 行数: 26 < 42
	c) 思路: 高手选择了先不进行“进位”， 而是最后进位， 我是每循环一次都要把结果加一次
	d) 技巧：1）如果有一个为"0"则返回0, make可以只传一个len

	高手主要比我快在以下几点：
	1） 少一层缓存
	   //高手
	   for i in num1:
		   for j in num2:
			   rst[i+j] += num1[i]*num[j]
	   //我
	   for i in num1:
		   cache = []
		   for j in num2:
			   cache[i+j] = num1[i]*num[j]
		   for x in rst:
			   # merge cache to rst

	2) 高手选择了先不进行“进位”， 而是最后进位， 我是每循环一次都要把结果加一次
	3）如果有0直接返回0
	4）idx从0开始算，而我是倒叙开始算，我这个算起来稍微比较复杂


2. 此题感悟
	在昨完第一轮时，感觉自己的直译还是不够精准，而且耗时太长，则进行了第二次开发，耗时9ms，符合预期，但与
	高手还是有很大差距，虽然在速度上一个是6ms，一个9ms，但是在细节上还是有很大差距，如上面我总结的4点。

	后续改进：
	  * 能不用缓存就不用，争取一次性将结果计算出来装箱。针对1
	  * 在可以极速返回结果的特殊情况，多考虑一下。针对3
	  * 在思路上，切记复杂，切记要正，要直切要害， 另外可以通过写伪代码的方式来做，由上往下递归填充代码。针对我的1st答案
	  * 临摹还是有用的，发现了之前没发现的问题。 临摹时，确实发现了diff，就是高手比我少存了一个中间层数组
	第一次作答弯弯扭扭的拐了太多弯，比如改变成int来乘啦什么的，还加入了递归，完全没必要，还降低了速度。。。
   

*/

func formatPattern(p string) []string{}{
	rst := []string{}
	
}

func isMatch3(s string, p string) bool {
	fmt.Println("isMatch3", s, p)
	if p == "*"{
		return true
	}
	if len(p) == 0{
		return len(s) == 0
	}
	pc := string(p[0])
	if pc != "?" && pc != "*"{
		i := 0
		for ;i<len(p) && p[i] != '*' && p[i] != '?'; i++{}
		pc = p[0:i]
		p = p[i:]
	}else if pc == "*"{
		i := 0
		for ;i<len(p) && p[i] == '*'; i++{}
		p = p[i:]
	}else{
		p = p[1:]
	}
	fmt.Println("left", pc, p, s)

	if pc == "*"{
		for i:=0; i<len(s); i++{
			if isMatch3(s[i:], p){
				return true
			}
		}
		return isMatch3("", p)
	}else if len(s)>0 && pc == "?"{
		return isMatch3(s[1:], p)
	}else{
		if len(pc) <= len(s) && pc == s[:len(pc)]{
			return isMatch3(s[len(pc):], p)
		}
		return false
	}
}


func isMatch1(s string, p string) bool {
	if p == "*"{
		return true
	}
	if len(p) == 0{
		return len(s) == 0
	}
	c := p[0]
	if c == '*'{
		if isMatch1(s, p[1:]){
			return true
		}
		for i:=1; i<len(s); i++{
			if isMatch1(s[i:], p[1:]){
				return true
			}
		}
		return false
	}else if len(s)>0 && (c == '?' || c == s[0]){
		return isMatch1(s[1:], p[1:])
	}else{
		return false
	}
}

func isMatch2(s string, p string) bool {
	if p == "*"{
		return true
	}
	if len(p) == 0{
		return len(s) == 0
	}
	for len(p)>=2 && p[0] == '*' && p[1] == '*'{
		p = p[1:]
	}
	c := p[0]
	if c == '*'{
		if isMatch2(s, p[1:]){
			return true
		}
		for i:=1; i<=len(s); i++{
			if isMatch2(s[i:], p[1:]){
				return true
			}
		}
		return false
	}else if len(s)>0 && (c == '?' || c == s[0]){
		return isMatch2(s[1:], p[1:])
	}else{
		return false
	}
}

func main() {
	// to_test := [][]int{
	// 	[]int{2, 3, 6, 7},
	// 	[]int{1},
	// 	[]int{8,7,4,3},
	// }

	// to_test2 := []int{
	// 	7,
	// 	1,
	// 	11,
	// }

	to_test := [][]string{
		[]string{"", "?"},
		[]string{"ho", "ho**"},
		[]string{"a", "a"},
		[]string{"a", "aa"},
		[]string{"aa", "aa"},
		[]string{"aaa", "aa"},
		[]string{"ab", "?*"},
		[]string{"aab", "c*a*b"},
		[]string{"aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaaba", "a*******b"},
		[]string{"aaba", "?***"},
		[]string{"abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb", "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"},
	}

	// to_test := [][][]byte{
	// 	[][]byte{
	// 		[]byte{'.','.','9','7','4','8','.','.','.'},
	// 		[]byte{'7','.','.','.','.','.','.','.','.'},
	// 		[]byte{'.','2','.','1','.','9','.','.','.'},
	// 		[]byte{'.','.','7','.','.','.','2','4','.'},
	// 		[]byte{'.','6','4','.','1','.','5','9','.'},
	// 		[]byte{'.','9','8','.','.','.','3','.','.'},
	// 		[]byte{'.','.','.','8','.','3','.','2','.'},
	// 		[]byte{'.','.','.','.','.','.','.','.','6'},
	// 		[]byte{'.','.','.','2','7','5','9','.','.'},
	// 	},
	// }
	for i:=0; i < len(to_test); i++{
		p1, p2:= to_test[i][0], to_test[i][1]
		rst := isMatch3(p1, p2)
		fmt.Println(p1, p2, rst)
	}
}
