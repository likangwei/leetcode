
package main
import "fmt"
/*

https://leetcode.com/problems/n-queens/description/

此题发挥不错，13ms。 大于了 75%的 commiter


高手答案:  

1. vs高手
    * 时间
    * 空间、变量个数
    * 行数
    * 命名、可读性
    * 技巧
    * 是否有递归
2. 此题感悟


*/


func buildNQueue(buf [][]int , curStep int , n int, rst (*[][]string)){
	// fmt.Println(buf, curStep, n, *rst)
	if curStep == n{
		board := make([]string, n)
		for i, row := range buf{
			bts := make([]byte, n)
			for j:=0; j<n; j++{
				bts[j] = '.'
			}
			bts[row[1]] = 'Q'
			board[i] = string(bts)
		}
		*rst = append(*rst, board)
		return
	}
	solts := make([]bool, n)
	for pIdx:=0; pIdx<curStep; pIdx++{
		point := buf[pIdx]
		solts[point[1]] = true
		dis := curStep - pIdx
		left, right := point[1] - dis, point[1] + dis
		if left >= 0 && left < n{
			solts[left] = true
		}
		if right >= 0 && right < n{
			solts[right] = true
		}
	}
	for i, b := range solts{
		if !b{
			buf[curStep] = []int{curStep, i}
			buildNQueue(buf, curStep+1, n, rst)
		}
	}
}

func solveNQueens(n int) [][]string {
    buf := make([][]int, n)
    rst := &([][]string{})
    buildNQueue(buf, 0, n, rst)
    return *rst
}

func main() {
	to_test := []int{
		4, 
	}

	for i:=0; i < len(to_test); i++{
		p1 := to_test[i]
		rst := solveNQueens(p1)
		for _, queen := range rst{
			for _, row := range queen{
				fmt.Println(row)
			}
			fmt.Println("\n")
		}
		
	}
}