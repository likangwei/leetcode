// https://leetcode.com/problems/palindrome-number/description/

package main
import "fmt"

func isPalindrome(x int) bool {
	if x < 0{
		return false
	}
	var num_len, left_loop, right_loop int = 0, 0, 0
	var left_num, right_num = x, 0
	for ;x != 0;{
		x = x / 10
		num_len += 1
	}
	if num_len == 1{
		return true
	}
	left_loop, right_loop = (num_len+1)/2, num_len/2
	for i:=0; i<right_loop; i++{
		right_num = right_num * 10 + (left_num % 10)
		left_num = left_num / 10
		left_loop -= 1
	}
	for j:=0; j<left_loop; j++{
		left_num = left_num / 10
	}
	return left_num == right_num
}

func main() {
	to_test := []int{1, -100, 121, 222, 343, 555, 1234, 12394}
	for i := 0; i < len(to_test); i=i+1{
		fmt.Println(to_test[i], "rst:", isPalindrome(to_test[i]))
	}
	
}
