package main
import "fmt"


// Definition for singly-linked list.
type ListNode struct {
     Val int
     Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var just_handle *ListNode
	for l1 != nil{
		cur := ListNode{}
		if just_handle != nil{
			just_handle.Next = &cur
		}else{
			head = &cur
		}
		just_handle = &cur
		cur.Val = l1.Val + l2.Val
		fmt.Println(cur.Val)
		l1 = l1.Next
		l2 = l2.Next
	}
	return head
}

func prettyPrint(l *ListNode){
	var lst []int
	for l != nil{
		lst = append(lst, l.Val)
		l = l.Next
	}
	fmt.Println(lst)
}

func main() {

	l1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	rst := addTwoNumbers(&l1, &l2)
	prettyPrint(rst)
}