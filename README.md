# 算法心得
---

##感悟

* 提前优化是万恶之源 # leetcode. No(34, )
* 直译思路，遇到困难，或者难翻译的地方不要想着跳过去  # leetcode.No(40, )

####算法思路

* 不要怕破坏原始结构，例如排序、图edges            # 旷视面试
* 正着循环的算法 看看可以根据情况倒着来吗

####思路正确性

* 在手动验证算法思路时，要根据testcase的多个维度制造，比如nums []int, 要从排序，长度等维度来造几个数据

####算法速度

* 算法类用while比较快，因为range和xrange都有开销
* 能用数字表示或者idx来控制的话就不再生成一个list
* 用最直接的方式来实现需求，不用过早的想着优化加一些map来cache，或者用map[int]int的方式，lst的遍历是很耗时的
* 排序后的数组，可以多想想怎么快速跳过index
* 多用内置算法，内置的都是C写的，肯定比我的快

##### 内置算法快
  for range
  == 

###更清晰的翻译
* 直译自己的思路，每一句代码在编写时要对照自己的解题思路，特别注意边界判断，idx判断等等，有一些地方用 <= 比 < 更能清晰的表达index界限
* target 与 data[index] 相比较时，最好target在前面比较清晰
* 在做index操作时，可以用弧线来表现到底跳动了几次
* 可以用纸笔来协助记忆思考
* 直译自己的思路，有一些不太好翻译的地方，不要绕过去，而是想办法直译，因为毕竟有很多代码模式自己没见过

###小技巧
* 数组可以想象成算盘，多个index代表某一个算盘珠
* for 循环时，有很多缓存的参数，比如i为当前轮询值，走到for loop 最后，要进行i++供下次使用，其他参数同理


###数据结构总结
* list。 特点
* LinkList。 特点： 插入,删除 O(1), 可以加个map来实现快速定位
* Map。 特点： 快速定位
* BinaryTree
* RedBlackTree
* Dag
* Graph

###命名积累
  index 用 i, j, k
  len 用m, n
  position: 二维的idx
  options: 选项
  occupied: 占据的， unoccupied: 未被占据的
  hi, lo 来代表高低位
  first, last 来用于前后出现的位置


---
```
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
股票最好的时间买卖，获取最大利润

# 学习后出来的答案
class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        rst = 0
        if prices:
            min_sofar = prices[0]
    
            for n in prices:
                if min_sofar < n:
                    rst = max(rst, n - min_sofar)
                elif min_sofar > n:
                    min_sofar = n
        return rst

复盘： 我的第一种算法，在学堂在线面试的答案是暴力轮训，时间复杂度 o(n!). 而高效率算法的思路感觉我没想到的原因是思路不清晰？脑子里是可以得出来正确结果的，

日后改正：写算法的时候，仔细想想如何将脑子里的想法翻译成算法

```
---
```
https://leetcode.com/problems/search-a-2d-matrix-ii/?tab=Description
二维矩阵排序
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
Given target = 5, return true.
Given target = 20, return false.

我的答案
class Solution(object):
    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """
        if not matrix:
            return False
            
        max_x = len(matrix[0])
        
        for row in matrix:
            for i in xrange(max_x):
                if row[i] < target:
                    continue
                elif row[i] > target:
                    max_x = i
                else:
                    return True
        return False
     
#高手答案
public class Solution {
    searchMatrix(int[][] matrix, int target) {
        if(matrix == null || matrix.length < 1 || matrix[0].length <1) {
            return false;
        }
        int col = matrix[0].length-1;
        int row = 0;
        while(col >= 0 && row <= matrix.length-1) {
            if(target == matrix[row][col]) {
                return true;
            } else if(target < matrix[row][col]) {
                col--;
            } else if(target > matrix[row][col]) {
                row++;
            }
        }
        return false;
    }
}

解析： 我的答案比高手答案效率低
后期改进：
  1. 算法类用while比较快，因为range和xrange都有开销
  2. 正着循环的算法 看看可以根据情况倒着来吗
  3. 有一些地方用 <= 比 < 更能清晰的表达index界限
  4. target 与 data[index] 相比较时，最好target在前面比较清晰
```

---
海明威距离高手答案
```python
class Solution(object):
    def hammingDistance(self, x, y):
        """
        :type x: int
        :type y: int
        :rtype: int
        """
        return bin(x^y).count('1')
解析： ^ 操作是二进制某一位匹配上就是1，不匹配就是0, 然后用bin
```

---
```
https://leetcode.com/problems/house-robber/?tab=Description

小偷偷钱问题，一排房子，连着偷两家会触发报警，问怎么在不触发报警的情况下偷更多的钱

class Solution(object):
    def rob(self, nums, m=None):
        """
        :type nums: List[int]
        :rtype: int
        """
        if m is None:
            m = {}
        k = str(nums)
        if k in m:
            return m[k]
        
        nl = len(nums)
        if nl == 1:
            return nums[0]
        if nl == 0:
            return 0
        
        l = nums[0] + self.rob(nums[2:], m=m)
        r = nums[1] + self.rob(nums[3:], m=m)
        rst = max(l, r)
        m[k] = rst
        return rst

为什么没有思路？

因为我开始并不想暴力破解此问题，以为有某些精妙算法。 
下次注意：
 1. 能用暴力破解的问题，起码是有解的
 2. 要注意简化问题，把非常长的列表简化成长度为2，3的列表来想

```

---
```
问题 DAG检测有没有环的算法
我的思路：
  把整个DAG分解成链表来检查链表的环

最优解： 拓扑排序

下次注意：
  1. 不要怕破坏其数据结构，可以完全copy出一份数据来

```


