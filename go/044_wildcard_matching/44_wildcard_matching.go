
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
	方法 struct 方法 带不带 *

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
	IsValid bool
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

func (r *ReBlock) Zoom(){
	beforeScope := []int{r.Scope[0], r.Scope[1]}
	before_lock := []bool{r.IsLockLeft, r.IsLockRight}
	fmt.Println("start zoom: ", r.Str())
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
	if !r.Valid(){
		fmt.Println("invalid", r.Str())
		return 
	}
	if r.Pre != nil && !r.Pre.IsLockRight{
		if (!before_lock[0] && r.IsLockLeft) || (beforeScope[1] > r.Scope[1]){
			fmt.Printf("invke others:%v\nbefore:\n beforlock %v beforescope %v \nafter:\n %v\n\n\n", r.Pre.Str(), before_lock, beforeScope, r.Str())
			r.Pre.Zoom()
		}
	} 
	if r.Next != nil && !r.Next.IsLockLeft{
		if (beforeScope[0] < r.Scope[0]) || (!before_lock[1] && r.IsLockRight){
			fmt.Printf("invke others:%v\nbefore:\n beforlock %v beforescope %v \nafter:\n %v\n\n\n", r.Next.Str(), before_lock, beforeScope, r.Str())
			r.Next.Zoom()
		}
	}
	fmt.Println("end zoom: ", r.Str())
}

func (r *ReBlock) Valid() bool{
	if (!r.IsValid){
		return false
	}
	if r.Scope[1] - r.Scope[0] < r.MinLen{
		r.IsValid = false
		return false
	}
	if r.IsLockLeft && r.IsLockRight{
		s := r.P[r.Scope[0]:r.Scope[1]]
		r.IsValid = isMatch3_3(s, r)
		return r.IsValid
	}
	return true
}

func (r *ReBlock) Str() string {
	return fmt.Sprintf(" %s: HasStar: %v, minLen:%d, Scope: %v, isStr: %v lock:[%v, %v]. Valid:%v", r.S, r.HasStar,r.MinLen,  r.Scope, r.IsStrMode, r.IsLockLeft, r.IsLockRight, r.IsValid)
}

func (r *ReBlock) Link() string {
	idx := 1
	rst := ""
	for r != nil{
		rst += fmt.Sprintf("%d. %v\n", idx, r.Str())
		r = r.Next
		idx++
	}
	return rst
}

func (r *ReBlock) SetLeft(n int) bool{
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

func (r *ReBlock) SetRight(n int) bool{
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
		block.IsValid = true
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


func isMatch3_2(s string, r *ReBlock, fromIdx int) bool {
	if r == nil{
		return fromIdx == len(s)
	}
	fmt.Println("\n\nmatch 3.2  ", s, r.Str(), fromIdx)

	r.Zoom()
	if fromIdx < r.Scope[0] || fromIdx > r.Scope[1]{
		return false
	}
	if r.IsLockLeft && fromIdx != r.Scope[0]{
		return false
	}
	fmt.Println(1)
	if r.Pre == nil && r.Next == nil{
		return isMatch3_3(s, r)
	}
	fmt.Println(2)
	if !r.Valid(){
		return false
	}
	fmt.Println(3)
	if r.IsLockRight{
		fmt.Println(4)
		return isMatch3_2(s, r.Next, r.Scope[1])

	}
	fmt.Println(5)
	if r.HasStar{
		fmt.Println("has star")
		for i:=r.Scope[1]; i>=r.Scope[0]+r.MinLen; i--{
			r.Zoom()
			fmt.Println("star...", fromIdx, i)
			if isMatch3_3(s[fromIdx:i], r){
				if isMatch3_2(s, r.Next, i){
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
				if isMatch3_2(s, r.Next, i){
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
	phead.Zoom()
	// return false
	return isMatch3_2(s, phead, 0)
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
	to_test = to_test
	for i:=0; i < len(to_test); i++{
		p1, p2:= to_test[i][0], to_test[i][1]
		rst := isMatch3(p1, p2)
		fmt.Println("rstttt", p1, p2, rst)
	}
}
