
package main
import "fmt"
/*
https://leetcode.com/problems/valid-sudoku/description/

我的解决过程：


高手答案:  6ms < 9ms

1. vs高手
	a) 命名
	b) 行数
	c) 思路
	d) 技巧：
	e) 此题感悟：
		

2. 此题感悟
	在直译方面，还需加强。用最直接的方案解决问题

*/

func multiply(num1 string, num2 string) string {
	longNum, shortNum, longIdx, shortIdx := num1, num2, len(num1)-1, len(num2)-1
	if longIdx < shortIdx{
		longNum, shortNum, longIdx, shortIdx = num2, num1, shortIdx, longIdx
	}
	rst := make([]byte, len(num1)+len(num2), len(num1)+len(num2))
	rstLen := 0
	buf := 0
	for shortIdx >= 0{
		b := longNum[longIdx] + (shortNum[shortIdx] - '0') + byte(buf)
		buf = 0
		if b > '9'{
			b, buf = b - 10, 1
		}
		rst[rstLen] = b
		rstLen ++
		shortIdx -- 
		longIdx --
	}
	for longIdx >= 0{
		b := longNum[longIdx] + byte(buf)
		buf = 0
		if b > '9'{
			b, buf = b - 10, 1
		}
		rst[rstLen] = b
		rstLen ++
		longIdx --
	}
	if buf == 1{
		rst[rstLen] = '0' + byte(buf)
		rstLen ++
	}
	rst = rst[:rstLen]
	return string(rst)
}

func main() {
	// to_test := [][]int{
	// 	[]int{2, 3, 6, 7},
	// 	[]int{1},
	// 	[]int{8,7,4,3},
	// }

	// to_test2 := []int{
	// 	7,
	// 	1,
	// 	11,
	// }

	to_test := [][]string{
		[]string{"123", "999"},
	}

	// to_test := [][][]byte{
	// 	[][]byte{
	// 		[]byte{'.','.','9','7','4','8','.','.','.'},
	// 		[]byte{'7','.','.','.','.','.','.','.','.'},
	// 		[]byte{'.','2','.','1','.','9','.','.','.'},
	// 		[]byte{'.','.','7','.','.','.','2','4','.'},
	// 		[]byte{'.','6','4','.','1','.','5','9','.'},
	// 		[]byte{'.','9','8','.','.','.','3','.','.'},
	// 		[]byte{'.','.','.','8','.','3','.','2','.'},
	// 		[]byte{'.','.','.','.','.','.','.','.','6'},
	// 		[]byte{'.','.','.','2','7','5','9','.','.'},
	// 	},
	// }
	for i:=0; i < len(to_test); i++{
		p1, p2:= to_test[i][0], to_test[i][1]
		rst := multiply(p1, p2)
		fmt.Println(p1, p2, rst)
	}

}
