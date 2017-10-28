package main
import "fmt"

func product(bytes_lst [][]byte)[] string{
	// fmt.Println(bytes_lst)
	w := len(bytes_lst)
	h := 1
	for i:=0; i<len(bytes_lst); i++{
		h *= len(bytes_lst[i])
	}
	lst := make([][]byte, h, h)
	for i:=0; i<h; i++{
		lst[i] = make([]byte, w, w)
	}
	preh := h
	for i:=0; i<w; i++{
		bytes := bytes_lst[i]
		cur_skip := preh / len(bytes)
		for j:=0; j<h; {
			for _, b := range bytes{
				for k:=0; k<cur_skip; k++{
					lst[j][i] = b
					j++
				}
			}
		}
		preh = cur_skip
	}
	rst := make([]string, h, h)
	for i:=0; i<h; i++{
		rst[i] = string(lst[i])
	}
	return rst
}

func letterCombinations(digits string) []string {
	if digits == ""{
		return []string{}
	}
	m := map[byte][]byte{
		'1': []byte{},
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k','l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
		'0': []byte{' '},
		'#': []byte{' '},
		'*': []byte{'+'},
	}
	bytes_lst := make([][]byte, 0, 0)
	for i:=0; i<len(digits); i++{
		bts, ok := m[digits[i]]
		if ok{
			bytes_lst = append(bytes_lst, bts)
		}else{
			fmt.Println("error")
		}
	}
	return product(bytes_lst)
}


func main() {
	to_test := []string{"2345", "22"}
	for i:=0; i < len(to_test); i ++{
		s := to_test[i]
		rst := letterCombinations(s)
		fmt.Println(rst)
	}
}
