

/*
https://leetcode.com/problems/longest-common-prefix/description/这个题目做了三天。。。

总结：

后续优化：

*/

package main
import "fmt"

func isMatch(s string, p string) bool {
	fmt.Printf("check isMatch(%s, %s) \n", s, p)

	if p == ".*"{
		return true
	}
	if len(p) == 0{
		return len(s) == 0
	}

	pc := p[0]
	atlast_show := 0
	star_count := 0
	keep_loop := false
	var pi, si int = 0, 0
	for ; pi < len(p); {
		if p[pi] == '*'{
			pi += 1
			star_count += 1
			keep_loop = true
		}else if p[pi] == pc{
			pi += 1
		}else{
			break
		}
	}
	atlast_show = pi - (2 * star_count)
	// fmt.Printf("%s need %d %c\n", p[:pi], atlast_show, pc)
	for ;atlast_show !=0 && si < len(s);{
		if pc == '.' || pc == s[si]{
			si += 1
			atlast_show -= 1
		}else{
			break
		}
	}
	if atlast_show != 0{
		return false
	}


	if !keep_loop{
		return isMatch(s[si:], p[pi:])
	}else{
		for ; si < len(s); si++{
			if isMatch(s[si:], p[pi:]){
				return true
			}
			if s[si] == pc || pc == '.'{
				// fmt.Println("i need ", s[si+1:], p[pi:])
				// if isMatch(s[si+1:], p[pi:]){
				// 	return true
				// }
			}else{
				break
			}
		}
		return isMatch(s[si:], p[pi:])
	}
}

func main() {
	to_test := []string{
						// "aa", "a",
						// "aa", "aa",
						// "aaa", "aa",
						// "aa", "a*",
						// "aa", "c*a",
						// "ab", ".*",
						// "aab", "c*a*b",
						// "aaa", "a*a",
						"aaa", "ab*a*c*a",
						// "a", "ab*",
						// "ab", ".*c"
					}
	for i := 0; i < len(to_test); i=i+2{
		fmt.Println(">>>", to_test[i], to_test[i+1], "rst:", isMatch(to_test[i], to_test[i+1]))
	}
	
}
