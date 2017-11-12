package main
import "fmt"

func twoSum(nums []int, target int) []int {
	var rst []int
	m := make(map[int]int)
	for i, n  := range nums{
		idx, ok := m[n]
		if ok{
			rst = append(rst, idx)
			rst = append(rst, i)
			break
		}else{
			num_found := target - n
			m[num_found] = i
		}
	}
	return rst
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	rst := twoSum(nums, target)
	fmt.Println(rst)
}