/**
题目概述:

给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

返回它的最大深度 3 。

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
func isSymmetric(root *TreeNode) bool {
  return check(root,root)
}

func check(p,q *TreeNode)bool{
  if p==nil && q==nil{
	  return true
  }
  if p ==nil || q==nil{
  	return false
  }
  return p.Val==q.Val && check(p.Left,q.Right) && check(p.Right,q.Left)
}

//迭代
//比较过程需要 逐层比较，可以使用 广度优先搜索（BFS） 解决，广度优先遍历可以借助 队列 来实现。
//此处 两两入队，两两处队。
func isSymmetric(root *TreeNode) bool {
  if root ==nil{
  	return true
  }
  return isMirror(root.Left,root.Right)
}

func isMirror(left,right *TreeNode) bool {
  q:=[]*TreeNode{}
  q=append(q,left,right)
  for len(q)>0{
  	n1,n2:=q[0],q[1]
  	q=q[2:]
  	if n1==nil && n2==nil{
  		continue
	}
  	if n1==nil || n2==nil{
  		return false
	}
  	if n1.Val !=n2.Val{
  		return false
	}
  	q=append(q,n1.Left,n2.Right)
  	q=append(q,n1.Right,n2.Left)
  }
  return true
}