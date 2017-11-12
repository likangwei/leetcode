/*
https://leetcode.com/problems/3sum-closest/

高手答案:
func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    result, d := 0, math.MaxInt32
    for i := 0; i < len(nums) - 2; i++ {
        j, k := i + 1, len(nums) - 1
        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            nd := sum - target
            
            if nd > 0 {
                if nd < d {
                    result, d = sum, nd
                }
                k--
                for k > j && nums[k] == nums[k + 1] {
                    k--
                }
            } else if nd < 0{
                nd *= -1
                if nd < d {
                    result, d = sum, nd
                }
                j++
                for j < k && nums[j] == nums[j - 1] {
                    j++
                }
            } else {
                return sum
            }
        }
    }
    
    return result
}

总结： 高手用了33行，我用了50行。高手考虑到了[1,1,1,1,1,2,2,2,2,2,3,3,3,3]这种情况，而且进行了
改进

思考： 可以有快速组合个数为count的nums的组合吗？ [1,1,1,1,1,2,2,2,2,2,3,3,3,3]

后需改进：
distance用d, current_distance 用 nd
index 用 i, j, k
len 用m, n

负数取正 可以用  num *= -1

*/

package main
import "fmt"
import "sort"


func sum(nums []int) int{
	rst := 0
	for _, n := range nums{
		rst += n
	}
	return rst
}

func getDistance(x int, y int) int {
	if x > y{
		return x - y
	}else{
		return y - x
	}
}

func combinations(nums[]int , count int, fromIdx int, target int)int{
	// fmt.Println(nums, count, fromIdx, target)
	cur_min_num := sum(nums[fromIdx:fromIdx+count])
	cur_max_num := sum(nums[len(nums)-count:len(nums)])
	if cur_min_num >= target{
		return cur_min_num
	}else if cur_max_num <= target{
		return cur_max_num
	}
	var rst int = cur_min_num
	var min_dist = getDistance(rst, target)
	i, j := fromIdx, len(nums) - count



	for ; i <= j; i++{
		for ; i <=j; {
			if sum(nums[j:j+count]) > cur_max_num{
				j--
			}else{
				break
			}
		}
		cur_num := nums[i]
		cb := combinations(nums, count-1, i+1, target-cur_num)
		cur_sum := cur_num + cb
		cur_dist := getDistance(cur_sum, target)
		if cur_dist < min_dist{
			rst = cur_sum
			min_dist = cur_dist
			cur_max_num = target + min_dist
			cur_min_num = target - min_dist
		}
		for ; i <=j&&i<len(nums)-1; {
			if nums[i] == nums[i+1]{
				i++
			}else{
				break
			}
		}
		for ; i <=j&&j<len(nums)&&j>=1; {
			if nums[j] == nums[j-1]{
				j--
			}else{
				break
			}
		}
	}
	return rst
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	rst := combinations(nums, 3, 0, target)
	return rst
}

func main() {
	to_test := [][]int{
		[]int{-1, 2, 1, -4},
		[]int{1, 1, 1, 1},
		[]int{0, 2, 1, -3},
		[]int{-11,-2,17,-16,1,-5,-5,-5,-20,7,10,-2,3,-7,-17,-13,-19,-15,-8,-7,6,-6,-8,-4,12,-12,9,-17,-13,4,-5,-15,-9,-18,-17,1,-15,-8,14,8,20,-3,-11,17,-18,10,-16,5,-9,-18,2,-3,4,-18,2,20,0,-6,18,-12,0,-17,3,-19,-20,15,12,-17,-7,8,16,7,-5,5,-13,16,-18,-7,-9,-8,-17,6,-18,0,-15,10,-13,7,9,20,7,-13,3,0,0,19,8,0,-5,-9,6,8,16,14,3,-4,5,9,-12,-19,16,6},
		[]int{47,-48,-72,97,-78,50,-22,18,9,24,28,-53,44,-96,50,45,86,11,21,-44,67,83,55,-86,-33,0,-53,-94,-60,57,-72,-73,-27,13,91,80,18,-80,-29,-69,-74,-90,54,22,3,91,-47,-32,80,-55,69,-95,62,-92,4,-86,62,3,23,-30,-4,0,49,24,10,-32,79,-99,-66,-30,-83,-13,90,-27,9,-4,9,98,-70,-19,32,24,-77,83,11,-78,-94,4,41,61,20,96,-36,54,-46,-51,91,54,30,-42,82,0,9,24,-2,32,-16,-18,87,23,78,-10,-82,-67,68,-18,-61,91,-90,-53,67,-48,12,1,-71,-99,31,82,39,-56,23,-89,-58,19,-60,39,-23,-76,-85,67,-33,69,-74,-8,-99,52,-70,-71,85,-8,28,-3,-100,18,88,5,-16,17,91,-35,22,-76},
		[]int{0,5,-1,-2,4,-1,0,-3,4,-5},

	}
	to_test2 := []int{
		1, 
		100, 
		1,
		-48, 
		298,
		1,
	}
	right := []int{
					2, 
					3, 
					0, 
					-48,
					291,
					0,
				}
	for i := 0; i < len(to_test); i=i+1{
		rst := threeSumClosest(to_test[i], to_test2[i])
		if rst == right[i]{
			fmt.Println("right>>>", to_test[i], "rst:", rst)
		}else{
			fmt.Println("wrong>>>", to_test[i], "rst:", rst)
		}
		
	}
	
}
