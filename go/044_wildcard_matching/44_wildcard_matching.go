
package main
import "fmt"
// import "strings"
/*
https://leetcode.com/problems/wildcard-matching/description/

我的解决过程：
  1st: 算法太慢，但比较简洁
  2nd: 优化 *** 为 *
  3rd: 优化 *?aaaa
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

type ReBlock struct{
	S string
	MinLen int
	IsStrMode bool
	HasStar bool
	Scope []int
	IsLockLeft bool
	IsLockRight bool
	Pre *ReBlock
	Next *ReBlock
	P string
	isValid bool
}

func isMatch3_3(s string, r *ReBlock)bool{
	rst := false
	if r.IsStrMode{
		rst = s == r.S
	}else if r.HasStar{
		rst = len(s) >= r.MinLen
	}else{
		rst = len(s) == r.MinLen
	}
	// fmt.Println("isMatch33", s, r.S, rst)
	return rst
}

func (r ReBlock) Zoom(){
	before := r
	fmt.Printf("before: %v\n", before.Str())

	if !r.HasStar{
		if r.IsLockLeft && !r.IsLockRight{
			r.Scope[1] = r.Scope[0] + r.MinLen
			r.IsLockRight = true
		}
		if r.IsLockRight && !r.IsLockLeft{
			r.Scope[0] = r.Scope[1] - r.MinLen
			r.IsLockLeft = true
		}
	}
	if !r.IsLockLeft{
		if r.Pre.Scope[0] + r.Pre.MinLen > r.Scope[0]{
				r.Scope[0] = r.Pre.Scope[0] + r.Pre.MinLen
		}
		if r.Pre != nil && r.Pre.IsLockRight{
			r.IsLockLeft = true
			r.Scope[0] = r.Pre.Scope[1]
		}
	}
	if !r.IsLockRight{
		if r.Next.Scope[1] - r.Next.MinLen < r.Scope[1]{
			r.Scope[1] = r.Next.Scope[1] - r.Next.MinLen
		}
		if r.Next != nil && r.Next.IsLockLeft{
			r.IsLockRight = true
			r.Scope[1] = r.Next.Scope[0]
		}
	}
	if r.Scope[1] - r.Scope[0] == r.MinLen{
		r.IsLockLeft = true
		r.IsLockRight = true
	}
	// if r.Pre != nil{
	// 	if (!before.IsLockLeft && r.IsLockLeft) || (before.Scope[1] > r.Scope[0]){
	// 		r.Pre.Zoom()
	// 	}
	// } 
	// if r.Next != nil{
	// 	if (before.Scope[0] < r.Scope[0]) || (!before.IsLockRight && r.IsLockRight){
	// 		r.Next.Zoom()
	// 	}
	// }
	fmt.Printf("after: %v\n", r.Str())

}

func (r ReBlock) Valid() bool{
	if !r.isValid{
		return false
	}
	if r.Scope[1] - r.Scope[0] < r.MinLen{
		r.isValid = false
		return false
	}
	if r.IsLockLeft && r.IsLockRight{
		s := r.P[r.Scope[0]:r.Scope[1]]
		r.isValid = isMatch3_3(s, &r)
		return r.isValid
	}
	return true
}

func (r ReBlock) Str() string {
	return fmt.Sprintf("%s: minLen:%d, Scope: %v, isStr: %v lock:[%v, %v]", r.S, r.MinLen, r.Scope, r.IsStrMode, r.IsLockLeft, r.IsLockRight)
}

func (r ReBlock) Link() string {
	idx := 1
	rst := ""
	p := &r
	for p != nil{
		rst += fmt.Sprintf("%d. %v\n", idx, p.Str())
		p = p.Next
		idx++
	}
	return rst
}

func (r ReBlock) SetLeft(n int) bool{
	if r.IsLockLeft{
		if r.Scope[0] == n{
			return true
		}else{
			fmt.Println("i'm locked. dont set left ", n, "scope", r.Scope)
			return false
		}
	}
	if r.Scope[0] < n{
		r.Scope[0] = n
		if !r.Valid(){
			return false
		}
		if r.Next != nil{
			r.Next.Zoom()
		}
	}
	return true
}

func (r ReBlock) SetRight(n int) bool{
	if r.IsLockRight{
		if r.Scope[1] == n{
			return true
		}else{
			fmt.Println("i'm locked. dont set right ", n, r.Str())
		}
	}
	if r.Scope[1] > n{
		r.Scope[1] = n
		if r.Valid(){
			if r.Pre != nil{
				r.Pre.Zoom()
			}
		}else{
			return false
		}
	}
	return true
}

func splitP(s string, p string)*ReBlock{
	
	head := ReBlock{}
	pre := &head

	for i:=0; i<len(p); {
		c := p[i]
		block := ReBlock{}
		block.Scope = []int{0, len(s)}
		block.P = s

		if c == '*' || c == '?'{
			j := i
			for ; j<len(p); j++{
				if p[j] == '*'{
					block.HasStar = true
				}else if p[j] == '?'{
					block.MinLen = block.MinLen+1
				}else{
					break
				}
			}
			block.S = p[i:j]
			i = j
		}else{
			j := i
			for ; j<len(p) && p[j] != '*' && p[j] != '?'; j++{
				block.MinLen = block.MinLen + 1
			}
			block.S = p[i:j]
			block.IsStrMode = true
			i = j
		}
		block.Pre = pre
		pre.Next = &block
		pre = &block
	}
	pre.IsLockRight = true
	var rst *ReBlock
	rst = head.Next
	if rst != nil{
		rst.Pre = nil
		head.Next = nil
		rst.IsLockLeft = true
	}
	return rst
}


func isMatch3_2(s string, r *ReBlock, stepIdx, fromIdx int) bool {
	fmt.Println("match 3.2  ", s, r.Str(), stepIdx, fromIdx)
	if r == nil{
		return fromIdx == len(s)
	}

	r.Zoom()
	if fromIdx < r.Scope[0] || fromIdx > r.Scope[1]{
		return false
	}
	if r.IsLockLeft && fromIdx != r.Scope[0]{
		return false
	}

	if r.Pre == nil && r.Next == nil{
		return isMatch3_3(s, r)
	}

	if !r.Valid(){
		return false
	}
	if r.IsLockRight{
		return isMatch3_2(s, r.Next, stepIdx+1, r.Scope[1])
	}

	if r.HasStar{
		for i:=r.Scope[1]; i>=r.Scope[0]+r.MinLen; i--{
			r.Zoom()
			if isMatch3_3(s[fromIdx:i], r){
				if isMatch3_2(s, r.Next, stepIdx+1, i){
					return true
				}
			}

			if r.IsLockLeft{
				if !r.SetRight(i-1){
					return false
				}
			}
		
		}
	}else{
		for i:=r.Scope[1]; i>=fromIdx+r.MinLen; i--{
			r.Zoom()
			if !r.Valid(){
				return false
			}
			if isMatch3_3(s[fromIdx:i], r){
				if isMatch3_2(s, r.Next, stepIdx+1, i){
					return true
				}
			}
			if r.IsLockLeft{
				if !r.SetRight(i-1){
					return false
				}
			}
		}
	}
	return false
}

func isMatch3(s string, p string) bool {
	phead := splitP(s, p)
	fmt.Println(phead.Link())
	return isMatch3_2(s, phead, 0, 0)
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
		[]string{"abcdefg", "a*fg"},
		[]string{"", "?"},
		[]string{"ho", "ho**"},
		[]string{"a", "a"},
		[]string{"a", "aa"},
		[]string{"aa", "aa"},
		[]string{"aaaaaaaa", "aa"},
		[]string{"ab", "?*"},
		[]string{"aab", "c*a*b"},
		[]string{"aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaaba", "a*******b"},
		[]string{"aaba", "?***"},
		[]string{"abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb", "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"},
		[]string{"abefcdgiescdfimde", "ab*cd?i*de"},
		[]string{"aaaaaabbaabaaaaabababbabbaababbaabaababaaaaabaaaabaaaabababbbabbbbaabbababbbbababbaaababbbabbbaaaaaaabbaabbbbababbabbaaabababaaaabaaabaaabbbbbabaaabbbaabbbbbbbaabaaababaaaababbbbbaabaaabbabaabbaabbaaaaba", "*bbb**a*******abb*b**a**bbbbaab*b*aaba*a*b**a*abb*aa****b*bb**abbbb*b**bbbabaa*b**ba**a**ba**b*a*a**aaa"},
		[]string{"baaabbabbbaabbbbbbabbbaaabbaabbbbbaaaabbbbbabaaaaabbabbaabaaababaabaaabaaaabbabbbaabbbbbaababbbabaaabaabaaabbbaababaaabaaabaaaabbabaabbbabababbbbabbaaababbabbaabbaabbbbabaaabbababbabababbaabaabbaaabbba", "*b*ab*bb***abba*a**ab***b*aaa*a*b****a*b*bb**b**ab*ba**bb*bb*baab****bab*bbb**a*a*aab*b****b**ba**abba"},
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
	to_test = to_test[0:1]
	for i:=0; i < len(to_test); i++{
		p1, p2:= to_test[i][0], to_test[i][1]
		rst := isMatch3(p1, p2)
		fmt.Println(p1, p2, rst)
	}
}
