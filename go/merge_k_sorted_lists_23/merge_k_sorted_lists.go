/*
https://leetcode.com/problems/merge-k-sorted-lists/description/

高手答案:
func mergeKLists(lists []*ListNode) *ListNode {
    amount := len(lists)
    interval := 1
    for interval < amount {
        for i:=0; i<amount-interval; i+=interval*2 {
            lists[i] = merge2Lists(lists[i], lists[i+interval])
        }
        interval *= 2
    }
    if amount > 0 {
        return lists[0]
    }
    return nil
}
func merge2Lists(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    curr := head
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            curr.Next = l1
            l1 = l1.Next
        } else {
            curr.Next = l2
            l2 = l1
            l1 = curr.Next.Next
        }
        curr = curr.Next
    }
    if l1 == nil {
        curr.Next = l2
    } else {
        curr.Next = l1
    }
    return head.Next
}

总结： 高手答案的思路，类似于快排，目的是尽量减少比对次数。 高手用两两合并的方式来进行排序，也有点类似于 晋升，如果入职时是一个
比较高的职级，那也不用跟职级下的人比较了，又或者是竞技比赛，好多人比赛 不用所有人全部比一遍，而是通过晋级的方式来做

其次，高手的这次优化主要是减少了比对次数

*/

package main
import "fmt"
import "math"

type ListNode struct {
	Val int
	Next *ListNode
}


func mergeKLists(lists []*ListNode) *ListNode {

	head := ListNode{}
	cur := & head
	var min_idx = -1
	var min_num = -1

	list_len := len(lists)

	for list_len > 1 {
		min_num = math.MaxInt64
		min_idx = -1
		for idx:=0; idx<list_len;{
			h := lists[idx]
			if h == nil{
				lists[idx] = lists[list_len-1]
				list_len--
				continue
			}
			if min_num > h.Val{
				min_idx = idx
				min_num = h.Val
			}
			idx ++
		}
		if min_idx != -1{
			cur.Next = lists[min_idx]
			cur = cur.Next
			lists[min_idx] = cur.Next
		}
	}
	if list_len == 1{
		cur.Next = lists[0]
	}
	return head.Next
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
		[]int{1},
		[]int{0},
	}
	var to_test2 []*ListNode
	for i:=0; i < len(to_test); i ++{
		lst := buildLinkList(to_test[i])
		printLinkList(lst)
		to_test2 = append(to_test2, lst)
	}

	rst := mergeKLists(to_test2)
	printLinkList(rst)

}
