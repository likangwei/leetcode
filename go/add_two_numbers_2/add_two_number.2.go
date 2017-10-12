package main
import "fmt"


// Definition for singly-linked list.
type ListNode struct {
     Val int
     Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{0, nil}
	cur := head
	for l1 != nil{
		cur.Val = l1.Val + l2.Val
		fmt.Println(cur.Val)
		prettyPrint(&head)
		l1 = l1.Next
		l2 = l2.Next
		tmp := ListNode{0, nil}
		cur.Next = &tmp
		cur = tmp

	}
	return &head
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
	prettyPrint(&l1)
	rst := addTwoNumbers(&l1, &l2)
	prettyPrint(rst)
}