package main
import "fmt"

func getPalindromString(s string, i int, j int) string {
	for i > 0 && j < len(s){
		if s[i-1] == s[j]{
			// fmt.Println("equal", s[i-1], s[j], i, j)
			i = i - 1
			j = j + 1
			// fmt.Println("after", i, j)
		}else{
			break
		}
	}
	return s[i:j]
}


func longestPalindrome(s string) string {
	if len(s) < 2{
		return s
	}
	longSubStr := s[0:1]
	for i:=1; i<=len(s)-1; i++{
		if (i + ((len(longSubStr)+1)/2)) > len(s){
			break
		}
		l1 := getPalindromString(s, i, i)
		l2 := getPalindromString(s, i, i+1)
		// fmt.Println(i, i, l1)
		// fmt.Println(i, i+1, l2)
		if len(l1) > len(longSubStr){
			longSubStr = l1
		}
		if len(l2) > len(longSubStr){
			longSubStr = l2
		}
	}
	return longSubStr
}

func main() {
	to_test := []string{"bb", "aaaa", "eeeeeeee"}
	for i := 0; i < len(to_test); i=i+1{
		fmt.Println(to_test[i], "rst:", longestPalindrome(to_test[i]))
	}
	
}