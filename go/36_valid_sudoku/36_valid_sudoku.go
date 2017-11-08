/*
https://leetcode.com/problems/implement-strstr/description/
高手答案 xms < xms

*/

package main
import "fmt"
/*
https://leetcode.com/problems/valid-sudoku/description/
我的解决过程：
  1. 竟然理解错了题意。。。 以为这个题的题意是 数独有没有解， 没想到竟然只是验证下当前已填充的数独是不是有效。。。

高手答案: 
   秒数相同

当前水平改进：
  虽然理解错了题意，但是我写出来的这个算法还是相当牛逼的。。。将一个数独有没有解写了出来。。。

*/

func getBoardStr(board [][]byte) string{
	bts := []byte{}
	for _, row := range board{
		for _, cell := range row{
			bts = append(bts, cell)
		}
		bts = append(bts, '\n')
	}
	return string(bts)
}

func ignoreOther(vboard [][][]byte, insert byte, rowIdx, colIdx int){
	for i:=0; i < 9; i++{
		for j, b := range vboard[rowIdx][i]{
			if b == insert{
				vboard[rowIdx][i] = append(vboard[rowIdx][i][0:j], vboard[rowIdx][i][j+1: len(vboard[rowIdx][i])]...)
				break
			}
		}
	}
	for i:=0; i < 9; i++{
		for j, b := range vboard[i][colIdx]{
			if b == insert{
				vboard[i][colIdx] = append(vboard[i][colIdx][0:j], vboard[i][colIdx][j+1: len(vboard[i][colIdx])]...)
				break
			}
		}
	}
	fromRow, fromCell := rowIdx / 3 * 3, colIdx / 3 * 3
	for i:=fromRow; i<fromRow+3; i++{
		for j:=fromCell; j<fromCell+3; j++{
			for k, b := range vboard[i][j]{
				if b == insert{
					vboard[i][j] = append(vboard[i][j][0:k], vboard[i][j][k+1: len(vboard[i][j])]...)
					break
				}
			}
			
		}
	}
}

func fill(board[][]byte, vboard [][][]byte, rowIdx, cellIdx int, insert byte)bool{
	// fmt.Println("fill",  rowIdx, cellIdx, insert)
	if board[rowIdx][cellIdx] == '.'{
		for _, k := range vboard[rowIdx][cellIdx]{
			if k == insert{
				board[rowIdx][cellIdx] = insert
				vboard[rowIdx][cellIdx] = []byte{}
				ignoreOther(vboard, insert, rowIdx, cellIdx)
				return true
			}
		}
		// fmt.Println("fill not found in vboard", rowIdx, cellIdx)
		return false
	}else{
		// fmt.Println("filling, but it's not none at ", rowIdx, cellIdx)
		return false
	}
}

func copyBoard(board [][]byte) [][]byte{
	rst := make([][]byte, len(board), len(board))
	for i:=0; i < len(board); i++{
		rst[i] = []byte{}
		rst[i] = append(rst[i], board[i]...)
	}
	return rst
}


func copyVBoard(board [][][]byte) [][][]byte{
	rst := make([][][]byte, len(board), len(board))
	for i:=0; i < len(board); i++{
		rst[i] = make([][]byte, len(board[i]), len(board[i]))
		for j:=0; j < len(board[i]); j++{
			rst[i][j] = []byte{}
			rst[i][j] = append(rst[i][j], board[i][j]...)
		}
	}
	return rst
}


func isValidSudokuDetail(board[][]byte, vboard [][][]byte) bool{
	// fmt.Println("isValidSudokuDetail", board, vboard)
	isValid := true
	minLength, minRowIdx, minCellIdx := 10, -1, -1
	for vRowIdx, vRow := range vboard{
		for vCellIdx, cell := range vRow{
			if len(cell) > 0{
				isValid = false
				curLength := len(cell)
				if curLength < minLength{
					minRowIdx, minCellIdx, minLength = vRowIdx, vCellIdx, curLength
					if curLength == 1{
						break
					}
				}
			}
		}
	}
	if isValid{
		// fmt.Println(getBoardStr(board))
		return true
	}else{
		for _, maybeNum := range vboard[minRowIdx][minCellIdx]{
			cpBoard := copyBoard(board)
			cpVboard := copyVBoard(vboard)
			if fill(cpBoard, cpVboard, minRowIdx, minCellIdx, maybeNum){
				if isValidSudokuDetail(cpBoard, cpVboard){
					return true
				}
			}else{
				return false
			}

		}
		return false
	}
}

func isValidSudoku(board [][]byte) bool {
	vboard := make([][][]byte, 9, 9)
	for i:=0; i < 9; i++{
		vboard[i] = make([][]byte, 9, 9)
		for j:=0; j<9; j++{
			vboard[i][j] = make([]byte, 9, 9)
			for k:=0; k<9; k++{
				vboard[i][j][k] = byte('1'+k)
			}
		}
	}
	for rowIdx, rowData :=  range board{
		for cellIdx, cell := range rowData{
			if cell != '.'{
				fill(board, vboard, rowIdx, cellIdx, cell)
				
			}
		}
	}
	return isValidSudokuDetail(board, vboard)
}

func isValidSudoku2(board [][]byte) bool {

	for i:=0; i<9; i++{
		for j:=0; j<9; j++{
			b := board[i][j]
			if b != '.'{
				for k:=0; k<9; k++{
					if k != j && board[i][k] == b{
						return false
					}
					if k != i && board[k][j] == b{
						return false
					}
				}
				headI, headJ := 3*(i/3), 3*(j/3)
				for k:=headI; k<headI+3; k++{
					for l:=headJ; l<headJ+3; l++{
						if k==i && l==j{
							continue
						}else{
							if board[k][l] == b{
								return false
							}
						}
					}
				}
			}
		}
	}
	return true
}


func main() {
	// to_test := [][]int{
	// 	[]int{5, 7, 7, 8, 8, 10},
	// }

	// to_test2 := []int{
	// 	8,
	// }

	// _ = [][]string{
	// 	[]string{"foo", "bar"},
	// 	[]string{"dhvf","sind","ffsl","yekr","zwzq","kpeo","cila","tfty","modg","ztjg","ybty","heqg","cpwo","gdcj","lnle","sefg","vimw","bxcb"},
	// }

	to_test := [][][]byte{
		[][]byte{
			[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		},
		[][]byte{
			[]byte{'.','8','7','6','5','4','3','2','1'},
			[]byte{'2','.','.','.','.','.','.','.','.'},
			[]byte{'3','.','.','.','.','.','.','.','.'},
			[]byte{'4','.','.','.','.','.','.','.','.'},
			[]byte{'5','.','.','.','.','.','.','.','.'},
			[]byte{'6','.','.','.','.','.','.','.','.'},
			[]byte{'7','.','.','.','.','.','.','.','.'},
			[]byte{'8','.','.','.','.','.','.','.','.'},
			[]byte{'9','.','.','.','.','.','.','.','.'},
		},
	}
	for i:=0; i < len(to_test); i++{
		p1 := to_test[i]
		rst := isValidSudoku2(p1)
		fmt.Println(getBoardStr(p1), rst)

	}

}
