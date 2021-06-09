/**
题目概述:

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶


*/
package main
/**
第一级台阶走法：1 种，一步走一级

第二级台阶走法：2 种，分开走两步每次走一级台阶，第二种直接一步到位。

第三级台阶走法： 3种，1 级台阶的走法 + 第 2 级台阶的走法，由于第二级台阶走法的最后一步位置不同所以走三级台阶就有 3 种走法

我们从第 4 步开始走，因为 n 个台阶走法 = n - 1 或者 n - 2

 */
func climbStairs(n int) int {
  if n<3{
  	return n
  }
  dp:=make([]int,n+1)
  dp[1]=1
  dp[2]=2
  for i:=3;i<=n;i++{
    dp[i]=dp[i-1]+dp[i-2]
  }
  return dp[n]
}

func climbStairs_1(n int) int {
	prev,curr:=1,1
	for i:=2;i<=n;i++{
		next:=prev+curr
		prev=curr
		curr=next
	}
	return curr
}

func climbStairs_2(n int) int {
	a,b:=1,1
	for i:=0;i<n;i++{
		a,b=b,a+b
	}
	return a
}