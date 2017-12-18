
package main
import "fmt"
/*

https://leetcode.com/problems/longest-increasing-path-in-a-matrix/description/

description:

Given an interger matrix. find the length of the longest
increasing path.

From each cell. you can either move to four directions:
left ,right ,up or down. you may not move diagonally
or move outside of the boundary(i.e. wrap-around is not allowed)


i.e. 即
increasing 正在增加的
diagonally 斜向的
boundary 边界
wrap-around 环绕式的

# 速度类

Q: 此次的时间消耗是？ vs 专家？
A: 第一次88ms, 看了export answer后58ms, 专家答案用了70+ms

Q: 此题目的算法复杂度是？
A: o(n)

# 可读性
Q: 此次的算法精炼度是怎么样的？
A: Great

Q: 此次的命名、可读性怎么样？
A: ok

# 思路类
Q: 此次是否有你递归，专家没递归的情况？
A: no

Q: 为什么你没有想到专家思路？
A: 因为开始想存edge的个数，有可能是0，但是后来不存edge了，存node个数了

# 后续改进
Q: 此次做的差的地方?
A: 对于存储时，用了[][][]int, 实际用[][]int就可以

Q: 此次发挥好的地方?
A: 一次submit ac。 代码干净利落

Q: 之前的总结是否生效？这次没用上的原因是？以后如何保证能用上？
A: 还ok

Q: 后续如何改进?
A: 要想清楚到底是存node好还是edge好，然后对比出其diff

Q: 此问题是否可以有实际应用场景？
A: 贪吃蛇？

Q: 此题目对应的TAG是topological-sort, depth-first-search, memoization, 此TAG解决了哪类问题？
A: topological-sort   拓扑排序
   depth-first-search 深度优先搜索
   memoization        记忆化

Q: 你觉得专家答案有什么值得你学习的地方？ 后续如何保证能用到，或者提高用到的概率？
A: 

Q: 对于日常生活中，专家的思维方式给我带来了哪些启发？
A:
最终总结，并加入复习列表:

这次发挥还是可以的，一次提交就submit ac
看到expert answer之后，又进行了优化，达到了top5的速度
不完美的地方在于，我用[][][]int存 node个数，实际上用[][]int就可以，
因为我一开始想存edge个数，但后来存了node个数，中间的diff没有理的特别清楚
后续要记住，用最简单的数据结构来表达你的思想


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