
package main
import "fmt"
// import "strings"
/*
https://leetcode.com/problems/wildcard-matching/description/

我的解决过程：
  1st: 算法太慢，但比较简洁
  2nd: 优化 *** 为 *
  3rd: 优化 *?aaaa
  4rd: 根据关键str来分割，比如 a*b*c*aaefafsf*   来优先处理出现次数少的str， 


高手答案:
func isMatch(s string, p string) bool {
    star := -1
    sIdx := 0
    last := 0
    pIdx := 0
    for sIdx < len(s) {
        if pIdx < len(p) && (p[pIdx] == s[sIdx] || p[pIdx] == '?') {
            sIdx++
            pIdx++
        } else if pIdx < len(p) && p[pIdx] == '*' {
            star = pIdx
            last = sIdx
            pIdx++
        } else if star != -1 { // reset to the star+1 position of p and last (advance by one) position of s 
            pIdx = star + 1
            last++
            sIdx = last
        } else {
            return false
        }
    }
    for pIdx < len(p) && p[pIdx] == '*' {
        pIdx++
    }
    return pIdx == len(p)
}

1. vs高手
    * 时间
       me) timelimit -> 109 ms -> 79 ms
       master) 15ms
    * 空间、变量个数
       me) []string: 1, int: 5， bool: 1, [][]int: 1， []int: 1, 
       master) int 4个
    * 行数 
       me) 112
       master) 26
    * 命名、可读性
    * 技巧
      master)
         回滚技巧高明，我用了递归，高手没用，类似于那个数独填充的题目No.37
         因为这个题的返回值是bool
         这道题的关键点：
           1: 试错题。
           2: 试错失败可回滚。 回滚技巧用了idx
         我的解法用了递归，用了[]string，递归用了切片。所以造成了很多不必要的数据出来
         递归题牵扯到几种情况：
            1）递归完，将结果封装到list返回，这种需要带着筐子（*[]int）去装
            2) 递归完，只要bool活着int的那种单个值，这个时候回滚就特别重要，高手
               的回滚技巧很6， 一维的完全可以用一个idx来控制，二维的可以用递归回滚来控制，类似No.37数独填充

2. 此题感悟
	* struct方法如果不用指针，就没法改变自身值类型属性
	* 此题可以不用递归，只用idx来回滚，因为他是一维的
	* 递归也可以精妙的回滚，类似数独填充No.37
	* 试错题都牵扯到回滚，回滚操作越轻越好

*/

func isMatch4_2(s string, plst []string)bool{


	if len(plst) == 0{ return len(s) == 0}
	if len(plst) == 1{
		p := plst[0]
		if p == "*"{return true}
		if p[0] == '?'{return len(s) == len(p)}
		return s == p
	}


	for len(plst)>0 && plst[0] != "*"{

		p := plst[0]
		if len(s) < len(p){ return false}

		if p[0] == '?' || s[:len(p)] == p{

			s = s[len(p):]
			plst = plst[1:]

			continue
		}

		return false
	}

	for len(plst)>0 && plst[len(plst)-1] != "*"{

		p := plst[len(plst)-1]
		if len(s) < len(p){return false}
		if p[0] == '?' || s[len(s)-len(p):len(s)] == p{
			s = s[0:len(s)-len(p)]
			plst = plst[0:len(plst)-1]
			continue
		}
		return false
	}

	if len(plst) <= 1{
		return isMatch4_2(s, plst)
	}

	// *str* , *, *a*a*a*aaa*
	m := make([][]int, len(plst))
	hasNoStr := true
	minCount := 0
	minLeft := 0
	for i:=0; i<len(plst); i++{
		m[i] = []int{}
		p := plst[i]
		if p[0] == '?'{
			minCount += len(p)
			minLeft += len(p)
			if minCount > len(s){
				return false
			}
			continue
		}else if p[0] == '*'{
			continue
		}
		hasNoStr = false
		for j:=minLeft; j+len(p)<=len(s); j++{
			if s[j: j+len(p)] == p{
				minLeft = j+len(p)
				m[i] = append(m[i], j)
			}
			for j=j+1; j+len(p) <=len(s); j++{
				if s[j: j+len(p)] == p{
					m[i] = append(m[i], j)
				}
			}
		}
		if len(m[i]) == 0{
			return false
		}
	}
	if hasNoStr{

		return len(s) >= minCount
	}

	minCount, minIdx := len(s), -1
	for i, lst := range m{
		p := plst[i]
		if p == "*" || p[0] == '?'{
			continue
		}
		l := len(lst)
		if l == 0{
			return false
		}
		if l <= minCount{
			minIdx = i
			minCount = l
		}
	}

	sLen := len(plst[minIdx])
	minLst := m[minIdx]

	for _, idx := range minLst{
		if isMatch4_2(s[0:idx], plst[0:minIdx]) && isMatch4_2(s[idx+sLen:], plst[minIdx+1:]){
			return true
		}
	}
	return false
}

func isMatch(s string, p string) bool {
	plst := []string{}

	for i:=0; i<len(p); {
		if p[i] == '*'{
			plst = append(plst, "*")
			for i=i+1; i<len(p) && p[i]=='*'; i++{}
		}else if p[i] == '?'{
			s := i
			for i+=1; i<len(p) && p[i]=='?'; i++{}
			plst = append(plst, p[s:i])
		}else{
			s := i
			for i+=1; i<len(p) && p[i]!='?' && p[i]!='*'; i++{}
			plst = append(plst, p[s:i])
		}
	}

	return isMatch4_2(s, plst)
}


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
	// fmt.Println("start zoom: ", r.Str())
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
			// fmt.Printf("invke others:%v\nbefore:\n beforlock %v beforescope %v \nafter:\n %v\n\n\n", r.Pre.Str(), before_lock, beforeScope, r.Str())
			r.Pre.Zoom()
		}
	} 
	if r.Next != nil && !r.Next.IsLockLeft{
		if (beforeScope[0] < r.Scope[0]) || (!before_lock[1] && r.IsLockRight){
			// fmt.Printf("invke others:%v\nbefore:\n beforlock %v beforescope %v \nafter:\n %v\n\n\n", r.Next.Str(), before_lock, beforeScope, r.Str())
			r.Next.Zoom()
		}
	}
	// fmt.Println("end zoom: ", r.Str())
}

func (r *ReBlock) Valid() bool{
	if (r.IsValid){
		if r.Scope[1] - r.Scope[0] < r.MinLen{
			r.IsValid = false
		} else if r.IsLockLeft && r.IsLockRight{
			s := r.P[r.Scope[0]:r.Scope[1]]
			r.IsValid = isMatch3_3(s, r)
		}
	}
	
	if (!r.IsValid){
		pre, next := r, r
		for pre != nil{
			pre.IsValid = false
			pre = pre.Pre
		}
		for next != nil{
			next.IsValid = false
			next = next.Next
		}
	}
	return r.IsValid
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
		for i:=r.Scope[1]; r.Valid() && i>=r.Scope[0]+r.MinLen && fromIdx<=i; i--{
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
		for i:=r.Scope[1]; r.Valid() && i>=fromIdx+r.MinLen && fromIdx<=i; i--{
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
		[]string{"b", "*?*?"},
		[]string{"abbaabbbbababaababababbabbbaaaabbbbaaabbbabaabbbbbabbbbabbabbaaabaaaabbbbbbaaabbabbbbababbbaaabbabbabb", "***b**a*a*b***b*a*b*bbb**baa*bba**b**bb***b*a*aab*a**"},
		[]string{"abbbaaababbaaabaaabbbabbbbaaabbaaababaabbbbbbaababbabababbababaaabbbbbabababaababaaaaaaabbbaabaabbbaabbabaababbabaababbbabbaaabbbaaaababbaaabbaabaabbbbbaaababaabaabaaabbabaabbbabbbaabbababaabbbbbbbbaaa", "*ba***bba*b**abbaa***a*****b*a*bb*b***a*bbb***a***bba*****a****a*a*b**aaaba*aab*a*aa***a*a*b**b**a*b*"},
		[]string{"a", "*a*"},
		[]string{"mississippi", "m*iss*iss*"},
		[]string{"abefcdgiescdfimde", "ab*cd?i*de"},

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
	to_test = to_test[18:len(to_test)]
	for i:=0; i < len(to_test); i++{
		p1, p2:= to_test[i][0], to_test[i][1]
		rst := isMatch4(p1, p2)
		fmt.Println("rstttt", p1, p2, rst)
	}
}
