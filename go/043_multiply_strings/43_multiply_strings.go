
package main
import "fmt"
/*
https://leetcode.com/problems/valid-sudoku/description/

我的解决过程：
  1st: multiplyLst 100ms 太慢了。。。

高手答案:  6ms < 9ms

1. vs高手
	a) 命名
	b) 行数
	c) 思路
	d) 技巧：
	e) 此题感悟：
		

2. 此题感悟
	在直译方面，还需加强。用最直接的方案解决问题

*/

func getInteger(lst []byte)int{
	rst, prod := 0, 1
	for i:=len(lst)-1; i>=0; i--{
		rst = rst + (prod * int(lst[i]-'0'))
		prod *= 10
	}
	return rst
}

func multiplyLst(longLst []byte, shortLst []byte)[]byte{
	if len(longLst) < len(shortLst){
		longLst, shortLst = shortLst, longLst
	}
	rst := []byte{}
	if len(longLst) + len(shortLst) > 18{
		for i:=len(longLst)-1; i>=0; i--{
			bts := multiplyLst(shortLst, []byte{longLst[i]})
			for j:=0; j<(len(longLst)-i-1); j++{
				bts = append(bts, '0')
			}
			k, l := len(bts)-1, len(rst)-1
			nlen := len(bts) + 1
			if len(bts) < len(rst){ nlen = len(rst) + 1}
			newRst := make([]byte, nlen, nlen)
			rstIdx := nlen - 1
			buf := 0
			for ; k>=0 || l>=0 || buf > 0;{
				b := byte('0')
				if k >= 0{
					b = b + bts[k] - '0'
					k--
				}
				if l >= 0{
					b = b + rst[l] - '0'
					l--
				}
				if buf > 0{
					b = b + byte(buf)
					buf = 0
				}
				if b > '9'{
					b, buf = b - 10, 1
				}
				newRst[rstIdx] = b
				rstIdx --
			}
			if newRst[0] == 0{
				newRst = newRst[1:len(newRst)]
			}
			rst = newRst
		}
		return rst
	}else{
		rstNum := getInteger(longLst) * getInteger(shortLst)
		if rstNum == 0{
			return []byte{'0'}
		}
		for rstNum > 0{
			rst = append([]byte{'0' + byte(rstNum % 10)}, rst...)
			rstNum = rstNum / 10
		}
		return rst
	}
}


func multiply(num1 string, num2 string) string {
	longNum, shortNum := num1, num2
	if len(longNum) < len(shortNum){
		longNum, shortNum = num2, num1
	}
	longBts := make([]byte, len(longNum), len(longNum))
	for i, b :=  range longNum{
		longBts[i] = byte(b)
	}
	shortBts := make([]byte, len(shortNum), len(shortNum))
	for i, b := range shortNum{
		shortBts[i] = byte(b)
	}
	return string(multiplyLst(longBts, shortBts))
}

func main() {
	// to_test := [][]int{
	// 	[]int{2, 3, 6, 7},
	// 	[]int{1},
	// 	[]int{8,7,4,3},
	// }

	// to_test2 := []int{
	// 	7,
	// 	1,
	// 	11,
	// }

	to_test := [][]string{
		[]string{"0", "0"},
		[]string{"1234", "999"},
		[]string{"9223372036854775808", "9223372036854775808"},
	}

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
		p1, p2:= to_test[i][0], to_test[i][1]
		rst := multiply(p1, p2)
		fmt.Println(p1, p2, rst)
	}

}
