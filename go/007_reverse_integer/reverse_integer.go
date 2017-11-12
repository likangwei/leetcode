
package main
import "fmt"

func reverse(x int) int {
	var MinInt, MaxInt int64 = -2147483648, 2147483647
	var tmp int64 = 0
	bgt_zero := x > 0
	if !bgt_zero{
		x = 0 - x
	}
	for ;x != 0;{
		tmp = tmp * 10 + (int64(x)%10)
		x = x/10
	}
	if bgt_zero{
		if tmp > MaxInt{
			return 0
		}else{
			return int(tmp)
		}
	}else{
		tmp = 0 - tmp
		if tmp < MinInt{
			return 0
		}else{
			return int(tmp)
		}
	}
}

func main() {
	to_test := []int{123, -312}
	for i := 0; i < len(to_test); i=i+1{
		fmt.Println(to_test[i], "rst:", reverse(to_test[i]))
	}
	
}
