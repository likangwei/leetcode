
package main
import "fmt"

func myAtoi(str string) int {
	validNumMap := map[rune]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}
	validJJMap := map[rune]int{'+': true, '-': false}
	spaceCharMap := map[rune]int{' ': 0, '\n': 0, '\t': 0}

	var i, j int = 0, 0
	bgt_zero := true
	for ; j<len(str); j++{
		c := rune(str[j])
		_, found_space := spaceCharMap[c]
		if found_space{
			continue
		}else{
			tmp, found_valid_char := validJJMap[c]
			if found_valid_char{
				bgt_zero = tmp
				
			}

				i = j
				for j=j+1; j < len(str); j++{
					c = rune(str[j])
					_, found_num := validNumMap[c]
					if found_num{
						continue
					}else{
						break
					}
				}
				break
			}else{
				break
			}
		}
	}
	rst, _ := int(str[i:j+1])
	return rst
}

func main() {
	to_test := []string{"  1123", "   -asdf", "-123aa1234"}
	for i := 0; i < len(to_test); i=i+1{
		fmt.Println(to_test[i], "rst:", myAtoi(to_test[i]))
	}
	
}
