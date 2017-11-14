
package main
import "fmt"
/*
https://leetcode.com/problems/valid-sudoku/description/

我的解决过程：
  1st: multiplyLst 100ms 太慢了。。。
  2nd: 9ms， 直译改进
高手答案:  6ms < 9ms
func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
		return "0"
	}
    lenNum1, lenNum2 := len(num1), len(num2)
	retArr := make([]int, lenNum1+lenNum2-1)
	for i:=0; i<lenNum1; i++ {
		for j:=0; j<lenNum2; j++ {
			tmp := (num1[i] - '0') * (num2[j] - '0')
			index := i + j
			retArr[index] += int(tmp)
		}
	}

	var ret string
	var flag int
	for i:=lenNum2+lenNum1-2; i>=0; i-- {
		a := (retArr[i]+flag) % 10
		flag = (retArr[i]+flag) / 10
		ret = string(a+'0') + ret
	}
	if flag != 0 {
		ret = string(flag+'0') + ret
	}
	return ret
}
1. vs高手
	a) 命名: 
	b) 行数: 26 < 42
	c) 思路: 高手选择了先不进行“进位”， 而是最后进位， 我是每循环一次都要把结果加一次
	d) 技巧：1）如果有一个为"0"则返回0

2. 此题感悟
	在昨完第一轮时，感觉自己的直译还是不够精准，而且耗时太长，则进行了第二次开发，耗时9ms，复合预期
    第一次作答弯弯扭扭的拐了太多弯，比如改变成int来乘啦什么的，还加入了递归，完全没必要，还降低了速度。。。

后续改进：
    在思路上，切记复杂，切记要正，要直切要害， 另外可以通过写伪代码的方式来做，由上往下递归填充代码
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


func multiply2(num1 string, num2 string) string {

	rst := make([]int, len(num1)+len(num2), len(num1)+len(num2))
	buf := 0
	for i:=len(num1)-1; i>=0; i--{
		product := make([]int, len(num2)+1, len(num2)+1)
		for j, k:=len(num2)-1, len(product)-1; j>=0; j, k=j-1, k-1{
			n := (int(num1[i]-'0') * int(num2[j]-'0')) + buf
			buf = 0
			if n > 9{
				n, buf = n%10, n/10
			}
			product[k] = n
		}
		product[0] += buf
		// fmt.Printf("%c %v %v\n", num1[i], num2, product)
		buf = 0
		for j, k := len(rst)-len(num1)+i, len(product)-1; k >=0 || buf > 0; j, k=j-1, k-1{
			n := rst[j]
			if k>=0{
				n = n + product[k]
			}
			n += buf
			buf = 0
			if n > 9{
				n, buf = n - 10, 1
			}
			rst[j] = n
		}
		buf = 0
	}
	for len(rst) >= 2 && rst[0]==0{
		rst = rst[1:]
	}

	bts := make([]byte, len(rst), len(rst))
	for i, n := range rst{
		bts[i] = byte(n) + '0'
	}
	return string(bts)
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
		[]string{"9133", "0"},
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
		rst := multiply2(p1, p2)
		fmt.Println(p1, p2, rst)
	}

}
