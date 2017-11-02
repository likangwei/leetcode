/*
https://leetcode.com/problems/implement-strstr/description/
高手答案 xms < xms

*/

package main
import "fmt"
/*
https://leetcode.com/problems/substring-with-concatenation-of-all-words/description/

高手答案

总结：

*/

type Map struct {
	Childs *map[string]*Map
}


func buildCombinationMap(words []string) *Map{
	if len(words) == 0{
		return nil
	}
	childs := make(map[string]*Map)
	rst := Map{&childs}
	for idx, s := range words{
		left_words := []string{}
		left_words = append(left_words, words[0:idx]...)
		left_words = append(left_words, words[idx+1:len(words)]...)
		fmt.Println(words, idx, "left", left_words)
		childs[s] = buildCombinationMap(left_words)
	}
	return &rst
}

func getStr(m *Map) string {
	if m == nil{
		return "nil"
	}
	rst := make(map[string]string)
	for k, v := range *m.Childs{
		rst[k] = getStr(v)
	}
	
	return fmt.Sprint(rst)
}

func findSubstring(s string, words []string) []int {
	rst := []int{}
	if len(words) == 0 || len(words[0]) == 0{
		return []int{0}
	}
	if len(s)==0{
		return []int{}
	}
	len_total := len(words[0]) * len(words)
	skip := len(words[0])
	if len_total > len(s){
		return []int{}
	}
	m := buildCombinationMap(words)
	fmt.Println(getStr(m))
	for i:=0; i <= len(s) - len_total; i++{
		curMap := m
		found := true
		for j:=i; j<i+len_total; j+=skip{
			curStr := s[j:j+skip]
			childs := *curMap.Childs
			v, ok := childs[curStr]
			if ok{
				curMap = v
			}else{
				found = false
				break
			}
		}
		if found{
			rst = append(rst, i)
		}
	}
	return rst
}

func main() {
	to_test := []string{
		"barfoothefoobarman",
	}

	param2 := [][]string{
		[]string{"foo", "bar"},
	}

	for i:=0; i < len(to_test); i+=1{
		rst := findSubstring(to_test[i], param2[i])
		fmt.Println(to_test[i], param2[i], rst)
	}
}
