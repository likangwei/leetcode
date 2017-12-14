
package main
import "fmt"
/*

https://leetcode.com/problems/longest-increasing-path-in-a-matrix/description/

description:

# 速度类

Q: 此次的时间消耗是？ vs 专家？
A: 

Q: 此题目的算法复杂度是？
A:

# 可读性
Q: 此次的算法精炼度是怎么样的？
A:

Q: 此次的命名、可读性怎么样？
A:

# 思路类
Q: 此次是否有你递归，专家没递归的情况？
A: 

Q: 为什么你没有想到专家思路？
A:

# 后续改进
Q: 此次做的差的地方?
A: 

Q: 此次发挥好的地方?
A:

Q: 之前的总结是否生效？这次没用上的原因是？以后如何保证能用上？
A:

Q: 后续如何改进?
A:

Q: 此问题是否可以有实际应用场景？
A:

Q: 此题目对应的TAG是topological-sort, depth-first-search, memoization, 此TAG解决了哪类问题？
A: 

Q: 你觉得专家答案有什么值得你学习的地方？ 后续如何保证能用到，或者提高用到的概率？

Q: 对于日常生活中，专家的思维方式给我带来了哪些启发？

最终总结，并加入复习列表:
 

3. TODO:
	* 

*/
func max(a, b int)int{
	if a > b{
		return a
	}
	return b
}

func getPath(i, j, m, n int, matrix [][]int, cache [][][]int, minNum int)int{
	if i < 0 || i > m-1 || j < 0 || j > n-1 || matrix[i][j] < minNum{
		return 0
	}
	if cache[i][j][0] == 1{
		return cache[i][j][1]
	}
	rst := 1
	num := matrix[i][j]
	//left
	rst = max(rst, 1+getPath(i, j-1, m, n, matrix, cache, num+1))
	//right
	rst = max(rst, 1+getPath(i, j+1, m, n, matrix, cache, num+1))
	//up
	rst = max(rst, 1+getPath(i-1, j, m, n, matrix, cache, num+1))
	//down
	rst = max(rst, 1+getPath(i+1, j, m, n, matrix, cache, num+1))
	cache[i][j] = []int{1, rst}
	return rst
}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0{
		return 0 
	}
	m, n := len(matrix),len(matrix[0])
	cache := make([][][]int, m)
	for i:=0; i<m; i++{
		cache[i] = make([][]int, n)
		for j:=0; j<n; j++{
			cache[i][j] = []int{0, 0}
		}
	}
	rst := 0
	for i:=0; i<m; i++{
		for j:=0; j<n; j++{
			p := getPath(i, j, m, n, matrix, cache, matrix[i][j]-1)
			rst = max(p, rst)
		}
	}
	return rst
}
func main() {
	to_test2 := [][]int{
		[]int{9, 9, 4},
		[]int{6, 6, 8},
		[]int{2, 1, 1},
	}
	to_test := [][][]int{
		to_test2,
	}
	for i:=0; i < len(to_test); i++{
		p1 := to_test[i]
		rst := longestIncreasingPath(p1)
		fmt.Println(p1, rst)
	}

}