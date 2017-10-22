/*
https://leetcode.com/problems/palindrome-number/description/

高手的答案不知道是不是比我快，但是对于循环控制思路上很好

public class Solution {
    public bool IsPalindrome(int x) {
        // Special cases:
        // As discussed above, when x < 0, x is not a palindrome.
        // Also if the last digit of the number is 0, in order to be a palindrome, 
        // the first digit of the number also needs to be 0.
        // Only 0 satisfy this property.
        if(x < 0 || (x % 10 == 0 && x != 0)) {
            return false;
        }
        int revertedNumber = 0;
        while(x > revertedNumber) {
            revertedNumber = revertedNumber * 10 + x % 10;
            x /= 10;
        }
        // When the length is an odd number, we can get rid of the middle digit by revertedNumber/10
        // For example when the input is 12321, at the end of the while loop we get x = 12, revertedNumber = 123, 
        // since the middle digit doesn't matter in palidrome(it will always equal to itself), we can simply get rid of it.
        return x == revertedNumber || x == revertedNumber/10;
    }
}
*/

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
