# 算法心得
---
###后续改进

####算法思路
* 不要怕破坏原始结构，例如排序、图edges
* 将大脑里的想法，慢慢转换成算法
* 正着循环的算法 看看可以根据情况倒着来吗
* 用最直接的方式来解决问题，而不是拐弯抹角，也别过多的想着优化，先解决当前问题


####算法速度
* 算法类用while比较快，因为range和xrange都有开销
* 能用数字表示或者idx来控制的话就不再生成一个list
* 用最直接的方式来实现需求，不用过早的想着优化加一些map来cache，或者用map[int][][int]的方式，lst的遍历是很耗时的

###更清晰的实现
* 完全直接的翻译脑中的算法，而不是碰到问题就绕过去
* 有一些地方用 <= 比 < 更能清晰的表达index界限
* target 与 data[index] 相比较时，最好target在前面比较清晰

## List
### 2: Add Two Link List

```
在可以完成答案的情况下，可以不用太拘谨。 

我的答案：
	var head *ListNode
	var pre_handle *ListNode
	for bababa:
		if pre_handle != nil{
			pre_handle.Next = &cur
		}else{
			head = &cur
		}
	return head

高手答案:
     var head *ListNode
     pre_handle = & head
     for bababa:
         省了我的if else
     return head.next
```
### 4. Median of Two Sorted Arrays
```
Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5

class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """

我的答案没有考虑到
nums1 = [1000], nums2 = range(1000) 的极限情况
这样我那个算法的复杂度就是 N(m+n)
3 = 1,1
4 = 1,2
5 = 2,2
6 = 2,3
假定
m, n = len(nums1), len(nums2)
l, r = (m-1)/2, m/2
所以我那个算法可以由
>>> while i < m and j < n
改进成
>>> while i < m and j < n and hand_count++ < r

再优化一下， nums1[-1] > nums2[-1] ， 所以肯定是nums2先被handle完，可以再优化
>>>large_lst = nums1 if nums1[-1] > nums2[-1] else nums2
>>> while large_i < len(large_lst) and hand_count++ < r

后续改进: 多枚举几种极限情况，根据不同的维度，比如这个题目有两层维度 lenght 和 num
```


## 树
### 95: 不同结构的平衡二叉树
 高手答案相对比我来讲, 没有生成过多的list, 而是用数字idx控制
 后续改进: 在用到过多list的时候, 看看是否可以用数字代替, 或者 idx代替

##动态规划:
### 122: 最好的股票卖出时间


---
```
题目: N*M的网格，将水倒入最上角网格，问一共流过了几个网格
grid = [[3, 5, 1], [2, 1, 5], [4,2, 1]]
n, m = len(grid), len(grid[0])
lst = [(0,0)]
rst = set()
while lst:
    element = lst.pop()
    rst.add(element)
    i, j = element
    left, right, top, bottom = (i, j-1), (i, j+1), (i-1, j), (i+1, j)
    for t in [left, right, top, bottom]:
        if t not in rst and 0 <= t[0] <= (n-1) and 0 <= t[1] <= (m-1):
            if grid[i][j] >= grid[t[0]][t[1]]:
                lst.append(t)
print len(rst)

```


---
```
一个list [n0, n1, n2,...] 求最大和子集

```

---
```
3SUM
For example, given array S = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
#高手答案
def threeSum(self, nums):
    res = []
    nums.sort()
    for i in xrange(len(nums)-2):
        if i > 0 and nums[i] == nums[i-1]:
            continue
        l, r = i+1, len(nums)-1
        while l < r:
            s = nums[i] + nums[l] + nums[r]
            if s < 0:
                l +=1 
            elif s > 0:
                r -= 1
            else:
                res.append((nums[i], nums[l], nums[r]))
                while l < r and nums[l] == nums[l+1]:
                    l += 1
                while l < r and nums[r] == nums[r-1]:
                    r -= 1
                l += 1; r -= 1
    return res

复盘：我的办法属于暴力破解法，效率低
后需改进： 不要怕破坏原始数据的数据结构, 包括： 排序，图等等
```


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


