
package main
import "fmt"
import "sort"
/*
https://leetcode.com/problems/valid-sudoku/description/

我的解决过程：


高手答案:  6ms < 9ms
func combinationSum1(candidates []int, target int, results *([][]int), combination []int, start int) {
	if target == 0 {
		// deep copy combination
		combination1 := make([]int, len(combination))
		copy(combination1, combination)

		*results = append(*results, combination1)
		return
	}
	for i := start; i < len(candidates) && target >= candidates[i]; i++ {
		combination = append(combination, candidates[i])
		combinationSum1(candidates, target-candidates[i], results, combination, i)
		combination = combination[:len(combination)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	results := [][]int {}
	combinationSum1(candidates, target, &results, []int {}, 0)

	return results
}


1. vs高手
    a) 命名 combination 组合
    b) 行数 24 < 44
    c) 思路 高手的思路是直译了到底递归应该怎么样递归，而我是利用数学层面来写的到底怎么算才能把结果弄出来
    d) 技巧： 递归时，带着一个框子把结果装起来
    e) 此题感悟：
        1）高手翻译更精确一些，后续不应绕过，哪怕能通过数学来解决此事
        2）递归时，带框子去装rst
        3) 关于需要一个unique组合的数据，题目 原始数据有误重复数据 是非常关键的因素
        4）扩展：组合分为几种 
           a) 算盘式组合: 各自不重合 == itertools.combinations, itertools.product
           b) 重合式组合: 自己与自己重合 == itertools.combinations_with_replacement
           c) 原始数据有重复数据的组合，输出set: 第一个数可以包含后续所有情况 
           d) 原始数据无重复数据的组合, 输出set: idx可以随意后移


2. 此题感悟
    在直译方面，还需加强。用最直接的方案解决问题

知识点：
func change(x []int){
	x[0] = 2
}

func change3(x[] int){
	x = []int{3, 3, 3}
}

func change2(x *[]int){
    *x = ([]int{1, 1, 1})
}

func main() {
	a := []int{1}
	fmt.Println(a)  // [1]
	change(a)
	fmt.Println(a)  // [2]
	change2(&a)
	fmt.Println(a)  // [1 1 1]
	change3(a)
	fmt.Println(a)  // [1 1 1]
}
*/

func combinationSum2(candidates []int, target int, fidx int) [][]int {
	// fmt.Println(candidates, target)
	var rst [][]int
	if target == 0{
		return rst
	}

	if len(candidates) == 0 || candidates[0] > target{
		return rst
	}

	for k:=fidx; k<len(candidates); k++{
		totalFirst := target / candidates[k]
		for i:=0; i<totalFirst; i++{
			count := i + 1
			targetNext := target - (candidates[k] * count)
			if targetNext == 0{
					curLst := make([]int, count, count)
					for j:=0; j<count; j++{
						curLst[j] = candidates[k]
					}
					rst = append(rst, curLst)
			}else{
				combsNext := combinationSum2(candidates, targetNext, k+1)
				for _, nextComb := range combsNext{
					curLst := make([]int, count, count)
					for j:=0; j<count; j++{
						curLst[j] = candidates[k]
					}
					curLst = append(curLst, nextComb...)
					rst = append(rst, curLst)
				}
			}
			
		}
	}
	return rst
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combinationSum2(candidates, target, 0)
}
func combinationDfs2(input []int, idx int, t int, buf *[]int, out *[][]int) {
    if t == 0 {
        found := make([]int, len(*buf))
        copy(found, *buf)
        *out = append(*out, found)
    } else {
        for j := idx; j < len(input); j++ {        
            if j > idx && input[j] == input[j - 1] {
                continue
            }
            
            d := input[j]
            
            if (t < d) {
                break
            }
            
            (*buf) = append((*buf), d)      /* push d */
            combinationDfs2(input, j + 1, t - d, buf, out)
            (*buf) = (*buf)[:len(*buf) - 1] /* pod */
        }         
    }
}

func combinationSum4(candidates []int, target int) [][]int {
    var buf []int
    var out [][]int

    sort.Ints(candidates)
    combinationDfs2(candidates, 0, target, &buf, &out)

    return out;
}

func main() {
	to_test := [][]int{
		[]int{2, 3, 6, 7},
		[]int{1},
		[]int{8,7,4,3},
	}

	to_test2 := []int{
		7,
		1,
		11,
	}

	// _ = [][]string{
	// 	[]string{"foo", "bar"},
	// 	[]string{"dhvf","sind","ffsl","yekr","zwzq","kpeo","cila","tfty","modg","ztjg","ybty","heqg","cpwo","gdcj","lnle","sefg","vimw","bxcb"},
	// }

	// to_test := [][][]byte{
	// 	[][]byte{
	// 		[]byte{'.','.','9','7','4','8','.','.','.'},
	// 		[]byte{'7','.','.','.','.','.','.','.','.'},
	// 		[]byte{'.','2','.','1','.','9','.','.','.'},
	// 		[]byte{'.','.','7','.','.','.','2','4','.'},
	// 		[]byte{'.','6','4','.','1','.','5','9','.'},
	// 		[]byte{'.','9','8','.','.','.','3','.','.'},
	// 		[]byte{'.','.','.','8','.','3','.','2','.'},
	// 		[]byte{'.','.','.','.','.','.','.','.','6'},
	// 		[]byte{'.','.','.','2','7','5','9','.','.'},
	// 	},
	// }
	for i:=0; i < len(to_test); i++{
		p1 := to_test[i]
		rst := combinationSum4(p1, to_test2[i])
		fmt.Println(p1, rst)
	}

}
