
package main
import "fmt"
/*

https://leetcode.com/problems/count-of-smaller-numbers-after-self/description/

description:
You are given an integer array nums and you have to return a new counts
array. The counts array has the property where counts[i] is the number
of       smaller elements to the right of nums[i].

you are given an integer array nums and you have to return a new _count_
array. the count array has the _properties_ _that_ counts[i] is the number
of _the_ smaller elements to the right of nums[i].

words:
maintain. Every node will maintain a val sum recording the total of number on it's left bottom side

专家答案:
public class Solution {
    class Node {
        Node left, right;
        int val, sum, dup = 1;
        public Node(int v, int s) {
            val = v;
            sum = s;
        }
    }
    public List<Integer> countSmaller(int[] nums) {
        Integer[] ans = new Integer[nums.length];
        Node root = null;
        for (int i = nums.length - 1; i >= 0; i--) {
            root = insert(nums[i], root, ans, i, 0);
        }
        return Arrays.asList(ans);
    }
    private Node insert(int num, Node node, Integer[] ans, int i, int preSum) {
        if (node == null) {
            node = new Node(num, 0);
            ans[i] = preSum;
        } else if (node.val == num) {
            node.dup++;
            ans[i] = preSum + node.sum;
        } else if (node.val > num) {
            node.sum++;
            node.left = insert(num, node.left, ans, i, preSum);
        } else {
            node.right = insert(num, node.right, ans, i, preSum + node.dup + node.sum);
        }
        return node;
    }
}


# 速度类

Q: 此次的时间消耗是？ vs 专家？
A: 1st: 1116ms https://leetcode.com/submissions/detail/130996108/

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

Q: 此题目对应的TAG是segment-tree, divide-and-conquer, binary-indexed-tree, binary-search-tree, 此TAG解决了哪类问题？
A: 

新get到的知识汇总：
* 

3. TODO:

*/

//Expert answer

type BTreeNode struct{
	val int
	sum int
	dup int
	left *BTreeNode
	right *BTreeNode
}

func insertNode(n int ,i int, node *BTreeNode, preSum int, rst []int) *BTreeNode{
	if node == nil{
		node = &BTreeNode{val: n, sum: 0, dup: 1}
		rst[i] = preSum
	}else if n > node.val{
		//turn right
		node.right = insertNode(n, i, node.right, preSum+node.sum+node.dup, rst)
	}else if n == node.val{
		node.dup = node.dup + 1
		rst[i] = preSum + node.sum
	}else{
		// turn left
		node.sum = node.sum + 1
		node.left = insertNode(n, i, node.left, preSum, rst)
	}
	return node
}


func countSmaller(nums []int) []int {
	rst := make([]int, len(nums))
	var root *BTreeNode
	for i:=len(nums)-1; i>=0; i--{
		n := nums[i]
		root = insertNode(n, i, root, 0, rst)
	}
	return rst
}


//mine
type Node struct{
	val int
	count int
	next *Node
}

func (this *Node)Str()string{
	s := ""
	for this != nil{
		s += fmt.Sprintf("->%d", this.val)
		this = this.next
	}
	return s
}

func countSmaller2(nums []int) []int {
	if len(nums) == 0{
		return nums
	}
	head := &Node{val:nums[len(nums)-1], next:nil, count:1}
	rst := make([]int, len(nums))
	for i:=len(nums)-2; i>=0; i--{
		n := nums[i]
		if n < head.val{
			curNode := Node{val:n, next:head, count:1}
			head = &curNode
			continue
		}
		var pre *Node
		cur := head
		curTotal := 0
		for cur != nil && cur.val < n{
			curTotal += cur.count
			pre = cur
			cur = cur.next
		}
		rst[i] = curTotal

		if cur == nil{
			pre.next = &Node{val:n, next:nil, count:1}
		}else if cur.val == n{
			cur.count = cur.count + 1
		}else{
			pre.next = &Node{val:n, next:cur, count:1}
		}
		// fmt.Println(head.Str())
	}
	return rst
}

func main() {
	to_test := [][]int{
		[]int{5, 2, 6, 1},
		[]int{26,78,27,100,33,67,90,23,66,5,38,7,35,23,52,22,83,51,98,69,81,32,78,28,94,13,2,97,3,76,99,51,9,21,84,66,65,36,100,41},
	}
	for i:=0; i < len(to_test); i++{
		p1 := to_test[i]
		rst := countSmaller(p1)
		rst2 := countSmaller2(p1)
		fmt.Println(p1)
		fmt.Println(rst)
		fmt.Println(rst2)
	}

}