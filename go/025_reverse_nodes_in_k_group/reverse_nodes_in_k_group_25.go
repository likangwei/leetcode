/*
https://leetcode.com/problems/merge-k-sorted-lists/description/

高手答案 9ms < 12ms

func reverseKGroup(head *ListNode, k int) *ListNode {
    var prev *ListNode
    //下次循环的值
    headp, node, tail, i := &head, head, head, 0
    for node != nil {
        // 此次循环的值, 赋值给下次循环的值
        node, node.Next, prev = node.Next, prev, node
        i ++

        if i == k {
            *headp, prev = prev, nil
            headp = &tail.Next
            tail = node
            i = 0
        }
    }
    if i != 0 {

        node, prev = prev, node
        for node != nil {
            node.Next, node, prev = prev, node.Next, node
        }
        *headp = prev
    }
    return head
}

感悟： 高手的指针用的666，用指针改变了另外一个指针的指针，还有for循环时用的很好，感悟这样

  这样很顺：
	for a, b, c {
		//当前循环 a, b, c 的值
		根据a, b ,c 进行某些操作

		// 为下一次循环赋值
		a, b, c = 下次的值
	}

*/

package main
import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
    var prev *ListNode
    headp, node, tail, i := &head, head, head, 0
    for node != nil {
        fmt.Println("node", node)
        node, node.Next, prev = node.Next, prev, node
        fmt.Println("prev", prev)
        i ++
        if i == k {
        	fmt.Println("in prev", *prev)
            *headp, prev = prev, nil
            fmt.Printf("1  headp=%d, head=%s \n",headp, head)
            headp = &tail.Next
            fmt.Printf("2  headp=%d, head=%s \n",headp, head)
            tail = node
            i = 0
        }
    }
    if i != 0 {
        node, prev = prev, node
        for node != nil {
            node.Next, node, prev = prev, node.Next, node
        }
        *headp = prev
    }
    return head
}


func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1{
		return head
	}
	nodes := make([]*ListNode, k, k)
	cur_len := 0
	var lastNode *ListNode = &ListNode{}
	firstNode := lastNode
	for head != nil{
		for head != nil && cur_len < k {
			nodes[cur_len] = head
			cur_len ++
			head = head.Next
		}
		if cur_len == k{
			nodes[0].Next = head
			for i:=cur_len-1; i>0; i--{
				nodes[i].Next = nodes[i-1]
			}
			lastNode.Next = nodes[cur_len-1]
			lastNode = nodes[0]
			cur_len = 0
		}else{
			lastNode.Next = nodes[0]
		}
		
	}
	return firstNode.Next
}

func buildLinkList(nums[]int) *ListNode {
	var head ListNode = ListNode{0, nil}
	var cur_node *ListNode = &head
	for _, n := range nums{
		cur_node.Next = &ListNode{n, nil}
		cur_node = cur_node.Next
	}
	return head.Next
}

func printLinkList(head *ListNode){
	curNode := head
	s := ""
	for ; curNode != nil; curNode=curNode.Next{
		s += fmt.Sprintf("%d->", curNode.Val)
	}
	fmt.Println(s)
}

func main() {
	to_test := [][]int{
		[]int{1,2,3,4,5},
		[]int{1,2,3,4,5},
	}
	to_test2 := []int{2, 3}

	for i:=0; i < len(to_test); i ++{
		lst := buildLinkList(to_test[i])
		printLinkList(lst)
		rst := reverseKGroup2(lst, to_test2[i])
		printLinkList(rst)
	}
}
