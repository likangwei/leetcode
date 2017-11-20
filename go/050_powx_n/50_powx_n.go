/*
https://leetcode.com/problems/powx-n/description/
func myPow(x float64, n int) float64 {
    var r float64 = 1.0
    if n < 0 {
        n = -n
        x = 1/x
    }
    for ;n>0;n>>=1 {
        if n&1 != 0 {
            r *= x
        }
        x *= x
    }
    return r
}
高手思路：  2**7 = (2 ** 4) * (2 ** 2) ** (2 ** 1)   7 = 0x0111
我的思路:   2 ** 7 = (2 ** 3) * (2 ** 3) * 2

1. vs高手
    * 时间 0ms < 3ms
    * 空间、变量个数 
    * 行数 14<16
    * 命名、可读性 
    * 技巧 
    * 是否有递归 高手没有，我有
2. 此题感悟
    关于次方，位移这方面，还是不够敏感，对数学的理解也是不够深刻
    关于负数次方，可以是 1/x
    临摹时，发现数比对时，用 !=0 比 == 1 时更快
*/
func myPow(x float64, n int) float64 {
    rst := x
    if n == 1{
        return x
    }else if n == 0{
        return 1
    }else if n < 0{
        rst := 1.0
        rst = rst/ myPow(x, 0-n)
        return rst
    }else{
        c := myPow(x, n/2)   
        rst = c * c * myPow(x, n%2)
        return rst
    }
}