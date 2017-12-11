package main
import "fmt"
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

