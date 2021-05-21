/**
题目概述:
给定两个数组，编写一个函数来计算它们的交集。

*/
package main

import "sort"

// 哈希方式
func intersect(nums1 []int, nums2 []int) []int {
	set := map[int]int{}
	for _, num1 := range nums1 {
		set[num1] += 1
	}
	k := 0
	for _, num2 := range nums2 {
		if set[num2] > 0 {
			set[num2] -= 1
			nums2[k] = num2
			k++
		}
	}
	return nums2[:k]

}

// 进阶 双指针
func intersect_2(nums1 []int, nums2 []int) []int {
	i, j, k := 0, 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)  // 排序
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++   // 小的索引++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}
