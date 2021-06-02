/**
题目概述:
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

*/
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 迭代遍历
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
   dummyHead:=&ListNode{}
   res:=dummyHead
   for l1 !=nil ||l2 !=nil{
   	  if (l1 !=nil && l2==nil) || (l1 !=nil && l1.Val<l2.Val){
   	  	 res.Next=l1
   	  	 res=l1
   	  	 l1=l1.Next
	  }else {
	  	res.Next=l2
	  	res=l2
	  	l2=l2.Next
	  }

   }
	return dummyHead.Next
}

// 递归
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}
	var dummyHead *ListNode
	if l1.Val>=l2.Val{
		dummyHead=l2
		dummyHead.Next=mergeTwoLists(l1,l2.Next)
	}else{
		dummyHead=l1
		dummyHead.Next=mergeTwoLists(l1.Next,l2)
	}
	return dummyHead.Next
}