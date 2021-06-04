/**
题目概述:

给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。

高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。



*/
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 中序遍历 左->根->右
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums)==0{
		return nil
	}
	mid:=len(nums)/2
	left:=nums[:mid]
	right:=nums[mid+1:]
	return &TreeNode{nums[mid], sortedArrayToBST(left),sortedArrayToBST(right)}
}