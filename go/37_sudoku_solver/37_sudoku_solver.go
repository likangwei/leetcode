
package main
import "fmt"
/*
https://leetcode.com/problems/valid-sudoku/description/

我的解决过程：
  根据昨天的答案做出来一版

高手答案: 9ms < 129 ms

func solveSudoku(board [][]byte)  {
    emptyPositions := getEmptyPositions(board)
    fillBoard(board, emptyPositions)
}

func getEmptyPositions(board [][]byte) [][]int {
    var positions [][]int
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if board[i][j] == '.' {
                positions = append(positions, []int{i, j})
            }
        }
    }
    return positions
}

func fillBoard(board [][]byte, positions [][]int) (finished bool) {
    if len(positions) == 0 {
        finished = true
        return
    }
    i, j := positions[0][0], positions[0][1]
    options := getOptions(board, i, j)
    for _, c := range options {
        board[i][j] = c
        if fillBoard(board, positions[1:]) {
            finished = true
            break
        }
	    board[i][j] = '.'
    }
    return 
}

func getOptions(board [][]byte, i, j int) []byte {
    occupied := make([]bool, 9)
    for k := 0; k < len(board[i]); k++ {
        if board[i][k] != '.' {
		occupied[int(board[i][k])-int('1')] = true
        }
    }
    for k := 0; k < len(board); k++ {
        if board[k][j] != '.' {
            occupied[int(board[k][j])-int('1')] = true
        }
    }
    for k := 0; k < 3; k++ {
        for h := 0; h < 3; h++ {
            if board[i/3*3+k][j/3*3+h] != '.' {
                occupied[int(board[i/3*3+k][j/3*3+h])-int('1')] = true
            }
        }
    }
    var unoccupied []byte
    for k := range occupied {
        if !occupied[k] {
            unoccupied = append(unoccupied, byte(k+int('1')))
        }
    }
    return unoccupied
}

高手

当前水平改进：
   命名：position: 二维的idx
        options: 选项
        occupied: 占据的， unoccupied: 未被占据的
   技巧： for idx := range lst ; 
  list数组的删减操作真的是非常慢，还有copy数组，copy二维数组
  高手特别精妙的完成了用1个二维数组board完成了填充和回滚，nb
  1）少用数组，更要少copy
  2) 对于递归要精妙的完成，不要用各种list增删查改
  3）能不用缓存就不用缓存，特别是list，二维，多维的大数据量，map可以多用
  4）精妙的利用递归可以节省很对内存空间
  5) 慎用缓存！！！慎用缓存！！！慎用缓存！！！重要的事情说三遍，用计算来替代
  6) 预制一个[]bool真的是可以超级快的做操作， 为什么？ 因为我直接可以达到O(1)的操作。
     举例： 将具有30个数的双色球进行筛选， 慢的方法是用list
     slow:
        lst = range(1, 31, 1)
        for n in random:
            for idx, e in enumerate(lst):
                if e == n:
                    lst = lst[0:i] + lst[i+1:] 或者 del lst[idx]
        return lst
     fast:
        lst = [True] * 30
        for n in random:
            lst[n-1] = False
        rst = []
        for idx, b in enumerate(lst):
            rst.append(idx+1)

*/

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

func getVBoardStr(board [][][]byte) string{
	bts := []byte{}
	for _, row := range board{
		for _, cell := range row{
			bts = append(bts, '[')
			bts = append(bts, cell...)
			bts = append(bts, ']')
		}
		bts = append(bts, '\n')
	}
	return string(bts)
}

func ignoreOther(vboard [][][]byte, insert byte, rowIdx, colIdx int)bool{
	rmIdx := []int{}
	for i:=0; i < 9; i++{
		rmIdx = append(rmIdx, rowIdx, i)
		rmIdx = append(rmIdx, i, colIdx)
	}
	fromRow, fromCell := rowIdx / 3 * 3, colIdx / 3 * 3
	for i:=fromRow; i<fromRow+3; i++{
		for j:=fromCell; j<fromCell+3; j++{
			rmIdx = append(rmIdx, i, j)
		}
	}
	for l:=0; l < len(rmIdx); l+=2{
		i, j := rmIdx[l], rmIdx[l+1]
		if len(vboard[i][j])==1 && vboard[i][j][0] == insert && i != rowIdx && colIdx != j{
			// fmt.Println("exclude fail..", i, j, vboard[i][j])
			return false
		}
	}
	for l:=0; l < len(rmIdx); l+=2{
		i, j := rmIdx[l], rmIdx[l+1]
		for k, b := range vboard[i][j]{
			if b == insert{
				vboard[i][j] = append(vboard[i][j][0:k], vboard[i][j][k+1: len(vboard[i][j])]...)
				break
			}
		}
	}
	return true
}

func fill(board[][]byte, vboard [][][]byte, rowIdx, cellIdx int, insert byte) bool {
	// fmt.Printf("fill %d %d, %c  %v\n", rowIdx, cellIdx, insert, vboard[rowIdx][cellIdx])
	if ignoreOther(vboard, insert, rowIdx, cellIdx){
		vboard[rowIdx][cellIdx] = []byte{}
		board[rowIdx][cellIdx] = insert
		return true
	}
	// fmt.Println("ignore fail...")
	return false
}


func solveSudoku2(board[][]byte, vboard [][][]byte)bool{
	// fmt.Printf("solveSudoku2\n%v\n\n", getBoardStr(board))
	minI, minJ, minLen := -1, -1, 10
	finish := true
	for vRowIdx, vRow := range vboard{
		for vCellIdx, cell := range vRow{
			if board[vRowIdx][vCellIdx] == '.' && len(cell) == 0{
				return false
			}
			if len(cell) >= 1{
				finish = false
				if minLen > len(cell){
					minI, minJ = vRowIdx, vCellIdx
				}
			}
		}
	}
	if finish{
		// fmt.Printf("finish..... \n%v\n%v", getBoardStr(board), getVBoardStr(vboard))
		return true
	}

	for _, n := range vboard[minI][minJ]{
		cpBoard := copyBoard(board)
		cpVBoard := copyVBoard(vboard)
		if fill(cpBoard, cpVBoard, minI, minJ, n){
			if solveSudoku2(cpBoard, cpVBoard){
				for i, row := range cpBoard{
					for j, cell := range row{
						board[i][j] = cell
					}
				}
				return true
			}
		}
	}
	// fmt.Printf("fail..... \n%v\n%v", getBoardStr(board), getVBoardStr(vboard))
	return false
}


func solveSudoku(board [][]byte){
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
	solveSudoku2(board, vboard)
	// fmt.Println(getVBoardStr(vboard))
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
			[]byte{'.','.','9','7','4','8','.','.','.'},
			[]byte{'7','.','.','.','.','.','.','.','.'},
			[]byte{'.','2','.','1','.','9','.','.','.'},
			[]byte{'.','.','7','.','.','.','2','4','.'},
			[]byte{'.','6','4','.','1','.','5','9','.'},
			[]byte{'.','9','8','.','.','.','3','.','.'},
			[]byte{'.','.','.','8','.','3','.','2','.'},
			[]byte{'.','.','.','.','.','.','.','.','6'},
			[]byte{'.','.','.','2','7','5','9','.','.'},
		},
		

	}
	for i:=0; i < len(to_test); i++{
		p1 := to_test[i]
		fmt.Println(getBoardStr(p1))
		solveSudoku(p1)
		fmt.Println(getBoardStr(p1))
	}

}
