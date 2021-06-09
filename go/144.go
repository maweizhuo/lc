/**
题目概述:
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
输入：root = [1,null,2,3]
输出：[1,2,3]
进阶：递归算法很简单，你可以通过迭代算法完成吗？

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
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var inner func(node *TreeNode)
	inner= func(node *TreeNode) {
		if node ==nil{
			return
		}
		res = append(res, node.Val)
		inner(node.Left)
		inner(node.Right)
	}
	inner(root)
	return res
}

// 非递归（栈,迭代）
func preorderTraversal(root *TreeNode) []int {
	var res []int
	stack:=[]*TreeNode{}
	node:=root
	for node !=nil || len(stack)>0{
		for node !=nil{
			res=append(res,node.Val)
			stack=append(stack,node)
			node=node.Left
		}
		node=stack[len(stack)-1].Right // 出栈
		stack=stack[:len(stack)-1]
	}
    return res
}