/**
题目概述:
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

位运算 -- 异或运算
1.任何数和自己做异或运算，结果为 00，即 a⊕a=0a⊕a=0 。
2.任何数和 00 做异或运算，结果还是自己，即 a⊕0=⊕a⊕0=⊕。
3.异或运算中，满足交换律和结合律，也就是 a⊕b⊕a=b⊕a⊕a=b⊕(a⊕a)=b⊕0=ba⊕b⊕a=b⊕a⊕a=b⊕(a⊕a)=b⊕0=b。
*/
package main


func singleNumber(nums []int) int {
     single :=0
     for _,num:=range nums{
     	single ^=num
	 }
     return single
}

