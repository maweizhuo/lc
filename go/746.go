/**
题目概述:

数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。

每当你爬上一个阶梯你都要花费对应的体力值，一旦支付了相应的体力值，你就可以选择向上爬一个阶梯或者爬两个阶梯。

请你找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。

 

示例 1：

输入：cost = [10, 15, 20]
输出：15
解释：最低花费是从 cost[1] 开始，然后走两步即可到阶梯顶，一共花费 15 。


*/
package main

func minCostClimbingStairs(cost []int) int {
   var prev,curr int
   for _,v:=range cost{
      prev,curr=curr,min(prev+v,curr+v)
   }
   return min(prev,curr)
}

func min(a,b int)int  {
	if a>b{
		return b
	}
	return a
}

func minCostClimbingStairs_1(cost []int) int {
	length:=len(cost)
	dp:=make([]int,length)
	dp[0]=cost[0]
	dp[1]=cost[1]
	for i:=2;i<length;i++{
		dp[i]=min(dp[i-1]+cost[i],dp[i-2]+cost[i])
	}
	return min(dp[length-1],dp[length-2])
}