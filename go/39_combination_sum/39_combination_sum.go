
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

当前水平改进：
   

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
		rst := combinationSum(p1, to_test2[i])
		fmt.Println(p1, rst)
	}

}
