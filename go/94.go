/**
题目概述:
给定一个二叉树的根节点 root ，返回它的 中序 遍历。

示例 1：

输入：root = [1,null,2,3]
输出：[1,3,2]

进阶: 递归算法很简单，你可以通过迭代算法完成吗？


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
// 递归
func inorderTraversal(root *TreeNode) []int {
   var res []int
   var inner func(node *TreeNode)
   inner=func(node *TreeNode){
   	  if node ==nil{
		  return
	  }
   	  inner(node.Left)
   	  res=append(res,node.Val)
   	  inner(node.Right)
   }
   inner(root)
   return res
}

// 非递归（栈,迭代）
func inorderTraversal(root *TreeNode) []int {
     var res []int
     stack:=[]*TreeNode{}
     for root !=nil || len(stack)>0{
     	for root !=nil{
     		stack=append(stack,root)
     		root=root.Left
		}
     	top:=len(stack)-1
     	root=stack[top]
		stack=stack[:top]
     	res=append(res,root.Val)
     	root=root.Right
	 }
     return res
}