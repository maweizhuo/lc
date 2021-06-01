/**
题目概述:
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？

*/
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 迭代法
func reverseList(head *ListNode) *ListNode {
   var newHead,next *ListNode
	for head !=nil{
		//先保存访问的节点的下一个节点，保存起来
		//留着下一步访问的
		next=head.Next
		//每次访问的原链表节点都会成为新链表的头结点，
		//其实就是把新链表挂到访问的原链表节点的
		//后面就行了
		head.Next=newHead
		//更新新链表
		newHead=head
		//重新赋值，继续访问
		head=next
	}
   return newHead
}

// 迭代法 (go语法糖)
func reverseList_2(head *ListNode) *ListNode {
	var newHead *ListNode
	for head !=nil{
		head,head.Next,newHead=head.Next,newHead,head
	}
	return newHead
}

// 递归
func reverseList_3(head *ListNode) *ListNode {
   if head ==nil|| head.Next ==nil{
   	return head
   }
   newHead:=reverseList_3(head.Next)
   head.Next.Next=head
   head.Next=nil
   return newHead
}
