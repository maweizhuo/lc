/**
题目概述:
统计所有小于非负整数 n 的质数的数量。

示例 1：

输入：n = 10
输出：4
解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
示例 2：

输入：n = 0
输出：0
示例 3：

输入：n = 1
输出：0


*/
package main

// 埃氏筛
//枚举没有考虑到数与数的关联性，因此难以再继续优化时间复杂度。接下来我们介绍一个常见的算法，该算法由希腊数学家厄拉多塞（\rm EratosthenesEratosthenes）提出，称为厄拉多塞筛法，简称埃氏筛。
func countPrimes(n int) int {
	cnt:=0
   isPrime :=make([]bool,n)
   for i:=range isPrime{
   	isPrime[i]=true
   }
   for i:=2;i<n;i++{
   	  if isPrime[i]{
   	  	cnt++
   	  	for j:=2*i;j<n;j+=i{
   	  		isPrime[j]=false
		}
	  }
   }
   return cnt

}

// 线性筛
// 相较于埃氏筛，我们多维护一个 \textit{primes}primes 数组表示当前得到的质数集合。我们从小到大遍历，如果当前的数 xx 是质数，就将其加入 \textit{primes}primes 数组。
func countPrimes_1(n int) int {
	isPrime :=make([]bool,n)
	primes:=[]int{}
	for i:=range isPrime{
		isPrime[i]=true
	}
	for i:=2;i<n;i++{
		if isPrime[i] {
			primes=append(primes,i)
		}
		for _,p:=range primes{
			if i*p>=n{
				break
			}
			isPrime[i*p]=false
			if i%p==0{
				break
			}
		}
	}
	return len(primes)

}