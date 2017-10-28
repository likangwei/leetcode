package main
import "fmt"

func prduct(bytes_lst [][]byte)[] string{
	l = len(bytes_lst)
	c := 1
	for i:=0; i<len(bytes_lst); i++{
		c *= len(bytes_lst[i])
	}
	lst := make([][]byte, c, c)
	for i:=0; i<c; i++{
		lst[i] = make([]byte, l, l)
	}
	for i:=0; i<len(bytes_lst); i++{
		bytes := bytes_lst[i]
		skip := c/len(bytes)
		for j:=0; j<c; j+=skip{
			ke := j+skip
			for k:=j; k<ke; k++{
				lst[k][i] = bytes[i]
			}
		}
	}
	rst := []string{}
	for i=0; i<c; i++{
		rst[i] = string(lst[i])
	}
	return rst
}

func letterCombinations(digits string) []string {
	m := map[byte][]byte{
		'1': []byte{''},
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k','l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
		'0': []byte{'+'},
		'#': []byte{' '},
		'*': []byte{''},
	}
}


func main() {
	to_test := []string{"23"}
	for i:=0; i < len(to_test); i ++{
		s := to_test[i]
		rst := letterCombinations(s)
		fmt.Println(rst)
	}
}
