
package main
import "fmt"
/*
TAG: array, greedy
https://leetcode.com/problems/jump-game-ii/description/

高手答案:  

1. vs高手
	* 时间
	* 空间、变量个数
	* 行数
	* 命名、可读性
	* 技巧
	* 是否有递归
2. 此题感悟

此次做的差的地方:

专家算法备注：BFS。BFS在求解最短路径或者最短步数上有很多的应用。


此次发挥好的地方:
* 代码行数首次超过专家答案

之前的总结是否游泳？这次没用上的原因是？以后如何保证能用上？

后续如何改进:

扩展思考：
    为什么查看能否到最后要倒序来做，而查看最短步数要正序来做？
*/

func jump(nums []int) int {
	rst := 0
	lastIdx := len(nums) - 1
	for lastIdx != 0{
		for i:=0; i<lastIdx; i++{
			if nums[i] + i >= lastIdx{
				lastIdx = i
				rst ++
				break
			}
		}
	}
	return rst
}

//Expert answer
func jump_master(nums []int) int {
   // BFS
    if len(nums) < 2 {
        return 0
    }
    currentMax, nextMax, level, i := 0, 0, 0, 0
    
    for i <= currentMax {
        level ++
        for ; i<=currentMax; i++{
            nextMax = max(nums[i]+i, nextMax)
            if nextMax >=len(nums) -1{
                return level
            }
        }
        currentMax = nextMax
    }
    return 0
}

func max(a, b int) int{
    if a < b {
        return b
    }
    return a
}


func main() {
	to_test := [][]int{
		[]int{2,3,1,1,4},
	}

	for i:=0; i < len(to_test); i++{
		p1 := to_test[i]
		rst := jump(p1)
		fmt.Println(p1, rst)
	}

}