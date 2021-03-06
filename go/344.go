/**
题目概述:
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。


*/
package main

func reverseString(s []byte)  {
	l:=0
	r:=len(s)-1
   for l<r{
   	    s[l],s[r]=s[r],s[l]
   	    l++
   	    r--
	}

}

func reverseString_2(s []byte)  {
	length:=len(s)
	mid:=length/2
	for i:=0;i<mid;i++{
		s[i],s[length-i-1]=s[length-i-1],s[i]
	}

}