

/*
https://leetcode.com/problems/longest-common-prefix/description/这个题目做了三天。。。

总结：

后续优化：

*/

package main
import "fmt"
func longestCommonPrefix(strs []string) string {
	var rst []byte
	var loop_finish = false
	if len(strs) == 0 {
		return ""
	}
	min_str_len := len(strs[0])
	for _, s := range strs{
		if min_str_len > len(s){
			min_str_len = len(s)
		}
	}
	for i:=0; i < min_str_len && !loop_finish; i++{
		cur_char := strs[0][i]
		for _, s := range strs{
			if cur_char != s[i]{
				loop_finish = true
				break
			}
		}
		if !loop_finish{
			rst = append(rst, cur_char)
		}
	}
	return string(rst)
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
		fmt.Println(">>>", to_test[i], to_test[i+1], "rst:", longestCommonPrefix(to_test))
	}
	
}
