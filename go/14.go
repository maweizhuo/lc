/**
题目概述:
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

 

示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"


*/
package main

import "strings"

func longestCommonPrefix(strs []string) string {
   if len(strs) ==0{
   	return ""
   }
   prefix:=strs[0]
   for _,str:=range strs{
   	  for strings.Index(str,prefix) !=0{
   	  	if len(prefix)==0{
   	  		return ""
		}
   	  	prefix=prefix[:len(prefix)-1] // 移除最后一个字符
	  }
   }
   return prefix
}