/**
题目概述:
给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

 */
package main

func maxProfit(prices []int) int {
	// 贪心思想
    profit:=0
    for i:=0;i<len(prices)-1;i++ {
    	if prices[i]<prices[i+1]{
    		profit+=prices[i+1]-prices[i]
		}
	}
    return profit
}