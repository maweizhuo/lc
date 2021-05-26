/**
题目概述:
请你判断一个 9x9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
数独部分空格内已填入了数字，空白格用 '.' 表示。

注意：

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。


*/
package main

// 哈希方式
func isValidSudoku(board [][]byte) bool {
	var heng [9]map[byte]int
	var zong [9]map[byte]int
	var fang [9]map[byte]int
	for i := 0; i < 9; i++ {
		if heng[i] == nil {
			heng[i] = make(map[byte]int)
		}
		for j := 0; j < 9; j++ {
			if zong[j] == nil {
				zong[j] = make(map[byte]int)
			}
			// 横判断
			if _, ok := heng[i][board[i][j]]; ok {
				return false
			} else {
				if board[i][j] != '.' {
					heng[i][board[i][j]] = 1
				}

			}

			// 纵判断
			if _, ok := zong[j][board[i][j]]; ok {
				return false
			} else {
				if board[i][j] != '.' {
					zong[j][board[i][j]] = 1
				}

			}

			// 方块判断
			index := (i/3)*3 + j/3
			if fang[index] == nil {
				fang[index] = make(map[byte]int)
			}

			if _, ok := fang[index][board[i][j]]; ok {
				return false
			} else {
				if board[i][j] != '.' {
					fang[index][board[i][j]] = 1
				}
			}

		}
	}
	return true
}

