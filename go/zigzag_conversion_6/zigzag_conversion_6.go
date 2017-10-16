package main
import "fmt"
import "strconv"

func max(a int, b int) int {  
    if a > b {  
        return a  
    }  
    return b  
}  

func convert(s string, numRows int) string {
	if len(s) < 2 || numRows == 1{
		return s
	}
	var numColumns int
	completeNum := len(s) / (numRows*2-2)
	remainderNum := len(s) % (numRows*2-2)
	remainderColNum := 0
	if remainderNum > 0{
		remainderColNum = 1 + max(0, remainderNum-numRows)
	} else{
		remainderColNum = 0
	}
	numColumns = completeNum *  (numRows-1) + remainderColNum
	matrix := make([][]rune, numRows, numRows)
	i, j := 0, 0
	isDown := true
	rst := make([]rune, len(s), len(s))
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
	c_idx := 0
	for i := range matrix{
		for j:= range matrix[i]{
			c := matrix[i][j]
			if c != 0{
				rst[c_idx] = c
				c_idx = c_idx + 1
			}
		}
	}

	return string(rst)
}

func main() {
	to_test := []string{"PAYPALISHIRING", "3", "AB", "1", "AB", "3"}
	for i := 0; i<len(to_test); i=i+2{
		s := to_test[i]
		numRows, _ := strconv.Atoi(to_test[i+1])
		fmt.Println(convert(s, numRows))
	}
}
