package main
import "fmt"


// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var pre_handle *ListNode
	extra := 0
	for l1 != nil || l2 != nil || extra != 0{
		cur := ListNode{}
		if pre_handle != nil{
			pre_handle.Next = &cur
		}else{
			head = &cur
		}
		pre_handle = &cur
		sum := extra
		if l1 != nil{
			sum = sum + l1.Val
			l1 = l1.Next

		}
		if l2 != nil{
			sum = sum + l2.Val
			l2 = l2.Next
		}

		cur.Val = sum % 10
		extra = sum / 10
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
	l3 := ListNode{5, nil}
	rst = addTwoNumbers(&l3, &l3)
	prettyPrint(rst)
}