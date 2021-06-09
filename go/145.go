/**
题目概述:
给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
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
func postorderTraversal(root *TreeNode) []int {
   var res []int
   var inner func(node *TreeNode)
   inner= func(node *TreeNode) {
	   if node ==nil{
		   return
	   }
	   inner(node.Left)
	   inner(node.Right)
	   res=append(res,node.Val)
   }
   inner(root)
   return res
}

// 非递归（栈,迭代）
func postorderTraversal(root *TreeNode) []int {
    var res []int
    var prev *TreeNode
    stack :=[]*TreeNode{}
    for root !=nil || len(stack)>0{
       for root !=nil{
       	 stack=append(stack,root)
       	 root=root.Left
	   }
       top:=len(stack)-1
       root=stack[top]
       stack=stack[:top]
       if root.Right ==nil || root.Right==prev{
       	  // 赋值
       	  res=append(res,root.Val)
       	  prev=root
       	  root=nil
	   }else{
	   	  stack=append(stack,root)
	   	  root=root.Right
	   }
	}
    return res
}