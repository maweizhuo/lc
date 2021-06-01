/**
题目概述:
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

进阶：你能尝试使用一趟扫描实现吗？
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

*/
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 快慢指针方式
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 采用全世界都会的快慢指针
	dummy := &ListNode{Next: head} //  虚拟节点

	low, fast := dummy, dummy
	// fast先走n+1步，然后一起前进，当fast为nil时，low就来到了倒数第n+1个节点，(因为需要是倒数第n+1个节点，才方便删除倒数第n个节点)
	for i := 0; i < n+1; i++ {
		fast = fast.Next //因为n始终合法，所以不需要考虑边界
	}

	for fast != nil {
		fast = fast.Next
		low = low.Next
	}
	// 删除倒数第n个节点
	low.Next = low.Next.Next
	return dummy.Next

}

func removeNthFromEnd_2(head *ListNode, n int) *ListNode {
	// 采用全世界都会的快慢指针
	dummy := &ListNode{Next: head} //  虚拟节点

	low, fast := dummy, dummy
	i:=0
	for fast != nil {
		fast = fast.Next
		i++
		if i>n+1{
			low = low.Next
		}
	}
	// 删除倒数第n个节点
	low.Next = low.Next.Next
	return dummy.Next

}

