
package main
import "fmt"
/*

https://leetcode.com/problems/count-of-smaller-numbers-after-self/description/

description:


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
A: O(n)

# 可读性
Q: 此次的算法精炼度是怎么样的？
A: 我的第一次还可以，虽然慢，但是也算是精确翻译了

Q: 此次的命名、可读性怎么样？
A: 还可以

# 思路类
Q: 此次是否有你递归，专家没递归的情况？
A: 没有

Q: 为什么你没有想到专家思路？
A: 我当时脑子里也有一个一闪而过的思路， 但是

# 后续改进
Q: 此次做的差的地方?
A: 1） 我意识到了，这种比对的问题应该是快速过滤，比如 比6小的，就是比5小的+5的个数，但是
   图快，先写了链表版的出来，效率感觉也还行，但离BST还稍微差点
   2） 思路可以写在纸上，或者记在手机上

Q: 此次发挥好的地方?
A: 链表版发挥还是不错的，起码解决了问题


Q: 之前的总结是否生效？这次没用上的原因是？以后如何保证能用上？
A: 总结了这种比对问题要通过快速筛选，类似于快排这种方法，没用上的原因是这个实现成本稍微高点，另外
   想到了链表解决方式，就先用的链表。后续如果时间充裕的话，多想想解决方案，这种属于创新类的，对
   自己的提升应该挺有帮助的

Q: 后续如何改进?
A: 对于常见的数据结构要进行统计，然后背一背

Q: 此问题是否可以有实际应用场景？
A: 暂时没想到

Q: 此题目对应的TAG是segment-tree, divide-and-conquer,
                  binary-indexed-tree, binary-search-tree, 此TAG解决了哪类问题？
A: 
 
Q: 你觉得专家答案有什么值得你学习的地方？ 后续如何保证能用到，或者提高用到的概率？

Q: 对于日常生活中，专家的思维方式给我带来了哪些启发？

最终总结，并加入复习列表:
 url: https://leetcode.com/submissions/detail/130996108/
 leetcode 315. 这道题目要根据nums, 返回counts.
 count[i] = the number of smaller elements to the right of nums[i].
 比如 [5, 2, 6, 1] => [2, 1, 1, 0]
 这道题目，我使用了链表来解决，用时1116 ms，专家答案用时 442ms，差一半吧
 这次我意识到了，这种比对的问题应该是快速过滤，比如 比6小的，就是比5小的+5的个数
 也和专家的思想有点类似，但是实现这个思路脑子里没有现成的方案，就先放弃了，又随即
 想到了链表的解决方案，虽然链表方式也非常好，但是不如BST来得快
 后续改进： 1. 将自己脑中模糊的想法或者方向，写在纸上，防止忘记
          2. 好的解决方案基于你当前的认知，比如我想到了要快速比对，但是脑中没有好的数据结构
             来表示和计算，所以后续
                1）提高自己的元能力，为什么我就不能创造一个BST出来呢，对吧
                2）还是要多吸收知识，脑中的数据结构多了，应该创造数据结构的能力就更强了

3. TODO:
  * 统计数据结构，并复习

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