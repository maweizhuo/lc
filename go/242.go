/**
题目概述:
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

输入: s = "anagram", t = "nagaram"
输出: true
*/
package main



func isAnagram(s string, t string) bool {
	hash:=map[rune]int{}
	for _,v:=range s{ //  遍历s,每个字符出现次数统计
		hash[v]++
	}
	for _,v:=range t{
       if count,ok:=hash[v]; !ok|| count==0{
           return false
	   }
       hash[v]--
	}
	for _,v:=range hash{
		if v!=0{
			return false
		}
	}
	return true
}

func isAnagram_2(s string, t string) bool {
	if len(s) !=len(t){
		return false
	}
	hash:=make([]int,26)
	for i:=0;i<len(s);i++ { //  遍历s,每个字符出现次数统计
		hash[s[i]-'a']++
		hash[t[i]-'a']--
	}

	for _,v:=range hash{
		if v!=0{
			return false
		}
	}
	return true
}