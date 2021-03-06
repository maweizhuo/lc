/**
题目概述:
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。


 */
package main

func removeDuplicates(nums []int)int{
	// 双指针解法
	slow:=0
	length:=len(nums)
	for fast:=1;fast<length;fast++{
		if nums[fast] !=nums[slow]{
			slow++
			nums[slow]=nums[fast]
		}
	}
	return slow+1
}