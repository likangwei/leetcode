/*
https://leetcode.com/problems/implement-strstr/description/
高手答案 xms < xms

*/

package main
import "fmt"
/*
https://leetcode.com/problems/implement-strstr/description/


高手答案
func strStr(haystack string, needle string) int {
    ci := 0
    for len(haystack[ci:]) >= len(needle) {
        if haystack[ci:ci + len(needle)] == needle {
            return ci
        }
        ci++
    }
    
    return -1
}

总结：高手算法简单粗暴，有我之前的风格。另外 内置的 == 比我的for循环比对快了N倍。。。 内置的肯定比我的快，以后多用内置算法

*/
func strStr(haystack string, needle string) int {
	nlen := len(needle)
	if haystack == needle{
		return 0
	}
	if nlen == 0{
		return 0
	}
	if len(haystack) < len(needle){
		return -1
	}

	idx_lst := make([]int, nlen, nlen)
	maxValidIdx := len(haystack) - nlen
	idxs_len := 0
	i := 0
	for i=0; i<len(haystack); i++{
		b := haystack[i]
		fmt.Println(i, string(b))
		if i <= maxValidIdx{
			if needle[0] == b{
				idx_lst[idxs_len] = i
				idxs_len += 1
			}
		}else{
			if len(idx_lst) == 0{
				return -1
			}
		}
		for j:=0; j<idxs_len; {
			ci := idx_lst[j]

			if b == needle[i-ci]{
				j ++
				if i+1-ci == nlen{
					return ci
				}
			}else{
				idx_lst[j] = idx_lst[idxs_len-1]
				idxs_len--
			}
		}
	}
	return -1
}

func main() {
	to_test := []string{
		"abcdeffffffffg", "ffg",
		"abbabaaaabbbaabaabaabbbaaabaaaaaabbbabbaabbabaabbabaaaaababbabbaaaaabbbbaaabbaaabbbbabbbbaaabbaaaaababbaababbabaaabaabbbbbbbaabaabaabbbbababbbababbaaababbbabaabbaaabbbba", "bbbbbbaa",
	}

	for i:=0; i < len(to_test); i+=2{
		rst := strStr(to_test[i], to_test[i+1])
		fmt.Println(to_test[i:i+2], rst)
	}
}
