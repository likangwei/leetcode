

/*
https://leetcode.com/problems/valid-parentheses/description/
总结：

后续优化：

*/

package main
import "fmt"

func isValid(s string) bool {
	str_len := len(s)
	if str_len % 2 != 0{
		return false
	}
	if s == ""{
		return true
	}
	m := map[byte]byte{'}':'{', ')': '(', ']': '['}

	max_stack_len := str_len / 2
	left_stack := make([]int, max_stack_len)
	cur_stack_len := 0

	for i:=0; i<str_len; i++{
		b := s[i]
		if b == '{' || b == '(' || b == '['{
			if cur_stack_len == max_stack_len{
				return false
			}
			left_stack[cur_stack_len] = i
			cur_stack_len += 1
			if cur_stack_len > max_stack_len{
				return false
			}
		}else if cur_stack_len == 0{
			return false
		}else if b == '}' || b == ')' || b == ']' {
			c, _ := m[b]
			if c == s[left_stack[cur_stack_len-1]]{
				cur_stack_len -= 1
			}else{
				return false
			}
		}else{
				return false
		}
	}
	return cur_stack_len == 0

}


func main() {
	to_test := []string{
		"()",
		"{}[]{}",
		"[){",
		"([)]",
		"((",
	}
	right := []bool{
		true,
		true,
		false,
		false,
		false,
	}
	for i := 0; i < len(to_test); i=i+1{
		rst := isValid(to_test[i])
		if right[i] == rst{
				fmt.Println("right", to_test[i], "rst:", rst)
			}else{
				fmt.Println("wrong", to_test[i], "rst:", rst)
			}
		
	}
	
}
