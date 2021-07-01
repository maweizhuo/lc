/**
题目概述:
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。
进阶：

你能否实现线性时间复杂度、仅使用额外常数空间的算法解决此问题?
 
示例 1：
输入：nums = [3,0,1]
输出：2
解释：n = 3，因为有 3 个数字，所以所有的数字都在范围 [0,3] 内。2 是丢失的数字，因为它没有出现在 nums 中。

示例 2：
输入：nums = [0,1]
输出：2
解释：n = 2，因为有 2 个数字，所以所有的数字都在范围 [0,2] 内。2 是丢失的数字，因为它没有出现在 nums 中。


*/
package main

// 位运算 ，异或
func missingNumber(nums []int) int {
  res:=0
  for i,v:= range nums{
  	res ^= i^v
  }
  return res ^ len(nums)
}

// 数学，高斯
func missingNumber_math(nums []int) int {
	n:=len(nums)
	sum:=n+(n+1)/2
	sum1:=0
	for _,v:=range nums{
		sum1+=v
	}
	return sum-sum1
}