
package main
import "fmt"
/*
TAG: array, greedy
https://leetcode.com/problems/jump-game-ii/description/

description:
  Give an array of  non-nevegation nums. You ar initially postioned at first

  index of the array. Every elements means your maximum jump lenght at that postion.

  please return the min jump step .

Note:
  you are assume the array can reach last postion.
  you are assume that you can always reach the last index.

错误单词

non-nevegation --> non-negative
          nums --> integers
     postioned --> positioned
     Every     --> Each
your goal is to reach the last index in the minimum number of jumps.

For example
Given array A [2,3,1,1,4]
The minimum number of jumps to reach the last index is 2.
(jump 1 step from index 0 to 1, the 3 steps to teh last index)

总结:


高手答案:  

1. vs高手
	* 时间
	* 空间、变量个数
	* 行数
	* 命名、可读性
	* 技巧
	* 是否有递归


2. 5why分析

Q: 此次做的差的地方?
A: 有模糊的地方没有摘出来

Q: 此次发挥好的地方?
A: 1）代码简洁，行数少
   2）一步步的将执行时间变短了

Q: 为什么你没有想到专家思路？
A: 1）固化了之前有类似题目的经验，但是那个是判断能不能调到最后，可以从后往前走，但这个是求最短步数
   2）专家用了BFS，BFS是解决类似问题的算法，而我是通过这个才百度到的BFS

Q: 之前的总结是否生效？这次没用上的原因是？以后如何保证能用上？
A: 1）之前总结的，返回bool的话就不用缓存，但是这次还是用了缓存
   2）没用上的原因是： 没有思路，而且刚开始的算法是递归开始的，一步步走到了这一步
   3）旧有的经验固然好，但是在解决问题时不能思维定式，再一个是对算法的时间复杂度不敏感，专家是O(n), 我是o(n方)

Q: 后续如何改进?
A: 1）解决问题时，将模糊的地方想清楚
   2）扩大算法知识面
   3）遇到问题时，不要思维定式
   4) 深入了解下时间复杂度

Q: 此问题是否可以有实际应用场景？
A: 各种地图导航类的

Q: 此题目对应的TAG是？此TAG解决了哪类问题？
A: 

3. TODO
* 将模糊的地方写在纸上
* 可以在纸上画的步骤或逻辑，在纸上画
* 如何扩大算法知识面？1） 背leetcode 算法TAG, 2) 每个算法TAG对应哪一类问题？ 3）如何润到每天任务里？
* 如何深入了解时间复杂度？背时间复杂度wiki? 润到每天任务里

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