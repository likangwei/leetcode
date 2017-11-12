

/*
https://leetcode.com/problems/4sum/description/

总结：

后续优化：

*/

package main
import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
 
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	lst_len := 0 
	if head == nil{
		return nil
	}
	cur_node := head
	m := make(map[int]*ListNode)
	for ; cur_node != nil; cur_node=cur_node.Next{
		lst_len += 1
		m[lst_len] = cur_node
	}

	if n == lst_len{
		return head.Next
	}else if n == 1{
		m[lst_len-1].Next = nil
		return head
	}else{
		m[lst_len-n].Next = m[lst_len-n+2]
		return head
	}
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
		s += "->" + fmt.Sprintf("%d", curNode.Val)
	}
	fmt.Println(s[2:])
}


func main() {
	to_test := [][]int{
		[]int{1,2,3,4,5},
		[]int{1,2},
	}
	to_test2 := []int{
		2,
		1,
	}
	for i := 0; i < len(to_test); i=i+1{
		head := buildLinkList(to_test[i])
		printLinkList(head)
		rst := removeNthFromEnd(head, to_test2[i])
		printLinkList(rst)
		// fmt.Println(">>>", to_test[i], "rst:", rst)
	}
	
}
