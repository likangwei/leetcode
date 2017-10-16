package main
import "fmt"
import "strconv"

func convert(s string, numRows int) string {
	if len(s) < 2 || numRows == 1{
		return s
	}
	var numColumns int
	numColumns = ((len(s)/(numRows*2-2))+1)*(numRows-1)
	fmt.Println(numColumns)
	matrix := make([][]rune, numRows, numColumns)
	i, j := 0, 0
	isDown := true
	rst := ""
	for t:=0; t<numRows; t++{
		matrix[t] = make([]rune, numColumns)
	}
	for _, c := range s{
		matrix[i][j] = c
		if i == 0{
			isDown = true
		}
		if i == numRows - 1{
			isDown = false
		}
		if isDown{
			i = i + 1 
		}else{
			i = i - 1
			j = j + 1
		}
	}
	
	for i := range matrix{
		for j:= range matrix[i]{
			c := matrix[i][j]
			if c != 0{
				rst = rst + string(c)
			}
		}
	}

	return rst
}

func main() {
	to_test := []string{"PAYPALISHIRING", "3", "AB", "1", "AB", "3"}
	for i := 0; i<len(to_test); i=i+2{
		s := to_test[i]
		numRows, _ := strconv.Atoi(to_test[i+1])
		fmt.Println(convert(s, numRows))
	}
}
