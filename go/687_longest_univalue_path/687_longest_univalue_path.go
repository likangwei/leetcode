
package main
import "fmt"
/*

https://leetcode.com/problems/longest-univalue-path/description/

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

Q: 此题目对应的TAG是recursion, tree, 此TAG解决了哪类问题？
A: 

Q: 你觉得专家答案有什么值得你学习的地方？ 后续如何保证能用到，或者提高用到的概率？

Q: 对于日常生活中，专家的思维方式给我带来了哪些启发？

最终总结，并加入复习列表:
 

3. TODO:
	* 

*/


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func max(a, b int)int{
    if a > b{
        return a
    }
    return b
}

func getLongestSub(node *TreeNode, buf *([]*TreeNode))int{
    if node == nil{
        return 0
    }
    rst := 0
    if node.Left != nil{
        if node.Left.Val == node.Val{
            rst = 1 + getLongestSub(node.Left, buf)
        }else{
            *buf = append(*buf, node.Left)
        }
        
    }
    if node.Right != nil{
        if node.Right.Val == node.Val{
            rst = max(rst, 1+getLongestSub(node.Right, buf))
        }else{
            *buf = append(*buf, node.Right)
        }
    }
    return rst
}

func getLongestPath(node *TreeNode, buf *([]*TreeNode))int{
    if node == nil{
        return 0
    }
    curPath := 0
    if node.Left != nil{
        if node.Left.Val == node.Val{
            curPath = 1 + getLongestSub(node.Left, buf)
        }else{
             *buf = append(*buf, node.Left)
        }
    }
    if node.Right != nil{
        if node.Right.Val == node.Val{
            curPath += 1 + getLongestSub(node.Right, buf)
        }else{
            *buf = append(*buf, node.Right)
        }
    }
    return curPath
}

func longestUnivaluePath(root *TreeNode) int {
    buf := make([]*TreeNode, 0, 10000)
    buf = append(buf, root)
    rst := 0
    for len(buf) != 0{
        node := buf[len(buf)-1]
        buf = buf[0: len(buf)-1]
        rst = max(getLongestPath(node, &buf), rst)
    }
    return rst
}

func main(){
    bottom1 := TreeNode{Val: 1}
    bottom2 := TreeNode{Val: 1}
    bottom3 := TreeNode{Val: 5}
    middle1 := TreeNode{Val: 4, Left: &bottom1, Right:&bottom2}
    middle2 := TreeNode{Val: 5, Right: &bottom3}
    head := TreeNode{Val:5, Left: &middle1, Right: &middle2}
    rst := longestUnivaluePath(&head)
    fmt.Println(rst)
}

