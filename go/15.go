/**
题目概述:
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：
输入：nums = []
输出：[]
示例 3：
输入：nums = [0]
输出：[]


*/
package main

import "sort"

// 排序+双指针
func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for l < r && n2 == nums[l] {
					l++
				}
				for l < r && n3 == nums[r] {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
