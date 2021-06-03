/**
题目概述:
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。


*/
package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
   return validBST(root,math.MinInt64,math.MaxInt64)
}

func validBST(root *TreeNode,min,max int)bool{
  if root ==nil{
  	return true
  }
  if root.Val<=min || root.Val>=max{
  	return false
  }
  return validBST(root.Left,min,root.Val) && validBST(root.Right,root.Val,max)
}