
package main
import "fmt"
/*

https://leetcode.com/problems/spiral-matrix/description/

descrition: Give a metric of  m*n elements (m rows, n columns), return all elements of the metric in spiral sorder.

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

func solution(matrix [][]int)[]int{
	if len(matrix) == 0 || len(matrix[0]) == 0{
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	rst := make([]int, m*n)
	i, j, rLen := 0, 0, 0
	length, hight := n, m
	for length > 1 && hight > 1{
		for k:=0; k<length-1; k++{
			rst[rLen] = matrix[i][j]
			j++
			rLen ++
		}
		for k:=0; k<hight-1; k++{
			rst[rLen] = matrix[i][j]
			i++
			rLen++
		}
		for k:=0; k<length-1; k++{
			rst[rLen] = matrix[i][j]
			j--
			rLen ++
		}
		for k:=0; k<hight-1; k++{
			rst[rLen] = matrix[i][j]
			rLen ++
			i--
		}
		i, j = i+1, j+1
		length, hight = length-2, hight-2
	}
	if hight == 1 && length >= 1{
		for k:=0; k<length; k++{
			rst[rLen] = matrix[i][j]
			rLen++
			j++
		}
	}else if length == 1 && hight >= 1{
		for k:=0; k<hight; k++{
			rst[rLen] = matrix[i][j]
			rLen++
			i++
		}
	}
	return rst
}

func main() {
	to_test := [][][]int{
		[][]int{
			[]int{1, 2, 3, 4},
			[]int{5, 6, 7, 8},
			[]int{9, 10, 11, 12},
		},
	}

	for i:=0; i < len(to_test); i++{
		p1 := to_test[i]
		rst := solution(p1)
		fmt.Println(p1, rst)
	}

}