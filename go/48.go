/**
题目概述:
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[[7,4,1],[8,5,2],[9,6,3]]


*/
package main

// 旋转。四个角的值替换
func rotate(matrix [][]int)  {

	length:=len(matrix)
	for i:=0;i<length/2;i++{ // 循环一半即可
		for j:=i;j<length-i-1;j++{
			temp:=matrix[i][j]
			m:=length-j-1
			n:=length-i-1
			matrix[i][j]=matrix[m][i]
			matrix[m][i]=matrix[n][m]
			matrix[n][m]=matrix[j][n]
			matrix[j][n]=temp
		}
	}
}
