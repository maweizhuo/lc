/**
题目概述:
给定一个整数数组，判断是否存在重复元素。

如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。

*/
package main

func containsDuplicate(nums []int) bool {
   set:=map[int]struct{}{}
   for _,v:=range nums{
   	if _,ok:=set[v];ok{
   		return true
	}
   	set[v]=struct{}{}
   }
   return false

}
