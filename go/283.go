/**
题目概述:
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。


*/
package main

// 双指针 i遍历数组，j指向非0的坑位
func moveZeroes(nums []int)  {
	j:=0
	for i:=0;i<len(nums);i++{
		if nums[i] !=0{
			nums[i],nums[j]=nums[j],nums[i]
			j++
		}
	}
}

// 双指针 i遍历数组，后赋值（j以后的为0）
func moveZeroes_2(nums []int)  {
	j:=0
	for i:=0;i<len(nums);i++{
		if nums[i] !=0{
			nums[j]=nums[i]
			j++
		}
	}
	for i:=j;i<len(nums);i++{
		nums[i]=0
	}
}
