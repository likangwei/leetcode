
package main
import "fmt"
import "sort"
/*
https://leetcode.com/problems/combination-sum-ii/description/

我的解决过程：
  有了昨天的经验，今天做起来还是很顺的

高手答案:  
func combinationDfs2(input []int, idx int, t int, buf *[]int, out *[][]int) {
	fmt.Println(input, idx, t, buf, out)
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
            
            (*buf) = append((*buf), d)      
            combinationDfs2(input, j + 1, t - d, buf, out)
            (*buf) = (*buf)[:len(*buf) - 1] 
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


当前水平改进：
  命名： 要带着的数据，可以叫 buf，而不是 bring
  
  关于组合不能重复的判定，不如 1，1，1，1，1，1  =》组合为3 ，因为第一个为1的会把所有的情况报含进去，所以
  用高手的写法是Ok的


*/

func combinationSum3(candidates []int, target int, start int, bring []int, rst *([][]int)){
	// fmt.Println(candidates, target, start, bring)
	if target == 0{
		*rst = append(*rst, bring)
		return
	}
	for i:=start; i<len(candidates) && candidates[i] <= target; {
		n := candidates[i]
		end := i+1
		for end <= len(candidates)-1 && candidates[end] == n{
			end++
		}

		for j:=i; j<end; j++{
			bring2 := bring
			next_target := target
			for k:=j; k<end; k++{
				bring2 = append(bring2, n)
				next_target -= n
			}
			combinationSum3(candidates, next_target, end, bring2, rst)
		}
		i = end
	}

}


func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	rst := [][]int{}
	combinationSum3(candidates, target, 0, []int{}, &rst)
	return rst
}


func main() {
	to_test := [][]int{
		// []int{10, 1, 2, 7, 6, 1, 5},
		// []int{1},
		[]int{1, 1, 1, 2},
	}

	to_test2 := []int{
		// 8,
		// 1,
		3,
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
		rst := combinationSum2(p1, to_test2[i])
		fmt.Println(p1, rst)
	}

}
