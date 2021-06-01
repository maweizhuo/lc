/**
题目概述:
给定一个正整数 n ，输出外观数列的第 n 项。

「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。

你可以将其视作是由递归公式定义的数字字符串序列：

countAndSay(1) = "1"
countAndSay(n) 是对 countAndSay(n-1) 的描述，然后转换成另一个数字字符串。
前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
第一项是数字 1
描述前一项，这个数是 1 即 “ 一 个 1 ”，记作 "11"
描述前一项，这个数是 11 即 “ 二 个 1 ” ，记作 "21"
描述前一项，这个数是 21 即 “ 一 个 2 + 一 个 1 ” ，记作 "1211"
描述前一项，这个数是 1211 即 “ 一 个 1 + 一 个 2 + 二 个 1 ” ，记作 "111221"


*/
package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
   result:=[]byte{'1'}
   for i:=1;i<n;i++{
     temp:=[]byte{} // 用于搬运空字符串
     for j:=0;j<len(result);j++{
        count:=1  // 统计字符个数
		 // 避免越界,所以j+1要小于len(res)
        for j+1<len(result) && result[j]==result[j+1]{
        	count++
        	j++
		}
		 // 将个数count和下标j所指字符加入temp
        temp=append(temp, byte(count+'0'),result[j])
	 }
     result=temp // 更新result
   }
   return string(result)
}

// 递归方式
func countAndSay_2(n int) string {
	var r, s string
	count := 1
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}
	s = countAndSay_2(n - 1)
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			// 如果s[i]!=s[i-1]，则把前面字符打印
			r += fmt.Sprintf("%s%s", strconv.Itoa(count), string(s[i-1]))
			count = 1
		} else {
			// 如果相等则计数器+1
			count++
		}
		// 如果到末尾了，打印当前字符
		if i+1 == len(s) {
			r += fmt.Sprintf("%s%s", strconv.Itoa(count), string(s[i]))
			s = r
			return s
		}
	}
	return s

}