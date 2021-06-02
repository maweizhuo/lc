/**
题目概述:
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false

*/
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 转成数组，双指针
func isPalindrome(head *ListNode) bool {
    vals:=[]int{}
    for head !=nil{
    	vals=append(vals,head.Val)
    	head=head.Next
	}
    // 双指针
    //slow,fast:=0,len(vals)-1
    //for slow <fast{
    //	if vals[slow]!=vals[fast]{
    //		return  false
	//	}
    //	slow++
    //	fast--
	//}
	fast:=len(vals)
	for i:=0;i<fast/2;i++{
		if vals[i] != vals[fast-i-1]{
           return false
		}
	}
    return true
}

// 快慢指针，翻转慢指针。比较
func isPalindrome(head *ListNode) bool {
   if head ==nil{
   	return true
   }
   slow,fast:=head,head
   for fast != nil && fast.Next !=nil {  // 找中点
       slow=slow.Next
       fast=fast.Next.Next // 多走一个
   }

   if fast !=nil{     //奇数链表取下一点
   	slow=slow.Next
   }

   var pre *ListNode
   for slow !=nil{            // 后半段反转链表
   	 slow,slow.Next,pre=slow.Next,pre,slow
   }

   for pre != nil {
   	 if pre.Val !=head.Val{
   	 	return false
	  }
   	  pre=pre.Next
   	  head=head.Next
   }
   return true //总共O(2n）/ O(2)

}