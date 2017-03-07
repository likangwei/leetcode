// `for` is Go's only looping construct. Here are
// three basic types of `for` loops.
// Input: x = 1, y = 4

// Output: 2
// Explanation:
// 1   (0 0 0 1)
// 4   (0 1 0 0)

package main

import "fmt"

func hammingDistance(x int, y int) int {
    rst := 0
    for i := 0; i <= 31; i++{
        n := 1 << uint(i)
        if x & n != y & n{
            rst += 1
        }
    }
    return rst
}

func main() {
    r := hammingDistance(1, 4)
    fmt.Println(r, dict())
}
