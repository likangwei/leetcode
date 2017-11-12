package main
import "fmt"

func getMedianFromLeft(lst[]int , left_lst[]int ,left_idx int, sum_len int) float64{
	// fmt.Println(lst, left_lst, left_idx, sum_len)
	var l, r, lnum, rnum int = 0, 0, 0, 0

	r = sum_len / 2
	if sum_len % 2 == 0{
		l = r - 1 
	}else{
		l = r
	}
	// fmt.Println("l, r", l, r)
	if l < len(lst){
		lnum = lst[l]
	}else{
		lnum = left_lst[left_idx + l - len(lst) ]
	}
	if r < len(lst){
		rnum = lst[r]
	}else{
		rnum = left_lst[left_idx + r - len(lst)]
	}
	return (float64(lnum) + float64(rnum)) / 2
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j int = 0, 0
	len1 := len(nums1)
	len2 := len(nums2)
	var lst []int
	var left_lst *[]int
	var left_idx int
	for ;i < len1 && j < len2; {
		if nums1[i] < nums2[j]{
			lst = append(lst, nums1[i])
			i = i + 1
		}else if nums1[i] > nums2[j]{
			lst = append(lst, nums2[j])
			j = j + 1
		}else {
			lst = append(lst, nums1[i])
			i = i + 1
			lst = append(lst, nums2[j])
			j = j + 1
		}
	}
	if i < len1{
		left_lst = & nums1
		left_idx = i
	}else{
		left_lst = & nums2
		left_idx = j
	}
	return getMedianFromLeft(lst, *left_lst, left_idx, len1+len2)
}


func main() {
	to_test := []*[]int{&[]int{1, 2}, &[]int{2}, &[]int{1, 2}, &[]int{3, 4}}
	for i := 0; i < len(to_test); i=i+2{
		var nums1, nums2 []int = *(to_test[i]), *(to_test[i+1])
		fmt.Println(findMedianSortedArrays(nums1, nums2))
	}
	
}