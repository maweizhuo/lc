/**
题目概述:
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

go 中单引号字符输出这个字符的ascii码
eg .fmt.Println('b','a') // 98,97
*/
package main


func firstUniqChar(s string) int {
    hash:=make([]int,26)
    for _,v:=range s{ //  遍历s,每个字符出现次数统计
    	hash[v-'a']++
	}
    for k,v:=range s{  // 再次遍历s，找出第一个频次为1的字符的索引
       if hash[v-'a']==1{
       	return k
	   }
    }
    return -1

}

