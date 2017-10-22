

/*
https://leetcode.com/problems/regular-expression-matching/description/

*/

package main
import "fmt"

func isMatch(s string, p string) bool {
	fmt.Printf("check isMatch(%s, %s) \n", s, p)

	if p == ".*"{
		return true
	}


	pc := p[0]
	keep_loop := false
	var si, pe int = 0, 1
	if pe < len(p) && p[pe] == '*'{
		keep_loop = true
		pe += 1
	}
	if !keep_loop && si < len(s){
		sc := s[si]
		if pc == '.' || pc == sc{
			return isMatch(s[si+1:], p[pe:])
		}else{
			return false
		} 
	}
	fmt.Println("start this")

	for ; si < len(s); si++{
		sc := s[si]
		if sc == pc || pc == '.'{
			if isMatch(s[si+1:], p[pe:]){
				return true
			}
		}else{
			break
		}
	}

	return isMatch(s, p[pe:])
}

func main() {
	to_test := []string{"aa", "a", "aa", "aa", "aaa", "aa", "aa", "a*", "aa", ".*", "ab", ".*", "aab", "c*a*b", "aaa", "a*a", "aaa", "ab*a*c*a", "a", "ab*", "ab", ".*c"}
	for i := 0; i < len(to_test); i=i+2{
		fmt.Println(">>>", to_test[i], to_test[i+1], "rst:", isMatch(to_test[i], to_test[i+1]))
	}
	
}
