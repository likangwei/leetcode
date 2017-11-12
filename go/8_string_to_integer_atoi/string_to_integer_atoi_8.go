
package main
import "fmt"

func myAtoi(str string) int {
	validNumMap := map[byte]int{byte('0'): 0,byte('1'): 1,byte('2'): 2,byte('3'): 3,byte('4'): 4,byte('5'): 5,byte('6'): 6,byte('7'): 7,byte('8'): 8,byte('9'): 9}
	validJJMap := map[byte]bool{byte('+'): true, byte('-'): false}
	spaceCharMap := map[byte]int{byte(' '): 0, byte('\n'): 0, byte('\t'): 0}

	var j int = 0
	var rst int = 0
	bgt_zero := true
	for ; j<len(str); j++{
		c := str[j]
		_, found_space := spaceCharMap[c]
		if found_space{
			continue
		}else{
			tmp, found_valid_char := validJJMap[c]
			if found_valid_char{
				bgt_zero = tmp
				j += 1
			}
			loop_count := 0
			for ; loop_count <= 10 && j<len(str);j++{
				loop_count = loop_count + 1
				n, found_num := validNumMap[str[j]]
				if found_num{
					rst = rst * 10 + n
					continue
				}
				break
			}
			break
		}
	}
	if !bgt_zero{
		rst = 0 - rst
	}
	if rst > 2147483647{
		return 2147483647
	} else if rst < -2147483648{
		return -2147483648
	}	
	return rst
}

func main() {
	to_test := []string{"  1123", "   -asdf", "-123aa1234"}
	for i := 0; i < len(to_test); i=i+1{
		fmt.Println(to_test[i], "rst:", myAtoi(to_test[i]))
	}
	
}
