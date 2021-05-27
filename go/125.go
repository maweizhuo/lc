/**
题目概述:
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

*/
package main

import "strings"

// 双指针。回文串的意思前面读和后面读是一致的
func isPalindrome(s string) bool {
	s=strings.ToLower(s)
    l:=0
    r:=len(s)-1
    for l<r{
    	if !isValid(s[l]){
    		l++
    		continue
		}
    	if !isValid(s[r]){
    		r--
			continue
		}
    	if s[l]!=s[r]{
    		return false
		}
    	l++
    	r--
	}
    return true

}

func isValid(c byte)bool{
	return (c>='a' && c<='z') || (c>='A' && c<='Z') || (c>='0' && c<='9')
}
