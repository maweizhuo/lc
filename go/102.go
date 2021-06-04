/**
题目概述:

给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。


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
//递归
var res [][]int
func levelOrder(root *TreeNode) [][]int {
       if root ==nil{
       	   return nil
	   }
	  res=make([][]int,0)
	  dfs(root,0)
	  return res

}

func dfs(root *TreeNode, level int)  {
	if root==nil{
		return
	}
	if level==len(res){
		res=append(res,[]int{})
	}
	res[level]=append(res[level],root.Val)
    dfs(root.Left,level+1)
    dfs(root.Right,level+1)
}

func levelOrder(root *TreeNode) [][]int {
    result:=[][]int{}
    if root ==nil{
    	return result
	}
    queue:=[]*TreeNode{root}
    for level:=0;0<len(queue);level++{ // 层级遍历
      result = append(result, []int{})
      next:=[]*TreeNode{}
      for j:=0;j<len(queue);j++{ // 层级元素遍历
      	node:=queue[j]
      	result[level]=append(result[level],node.Val)
      	if node.Left !=nil{
      		next=append(next,node.Left)
		}
      	if node.Right !=nil{
      		next=append(next,node.Right)
		}
	  }
      queue=next
	}
  return result
}

