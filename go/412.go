/**
题目概述:
写一个程序，输出从 1 到 n 数字的字符串表示。

1. 如果 n 是3的倍数，输出“Fizz”；

2. 如果 n 是5的倍数，输出“Buzz”；

3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。



*/
package main

import "strconv"

func fizzBuzz(n int) []string {
	var arr []string
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			arr = append(arr, "FizzBuzz")
		} else if i%3 == 0 {
			arr = append(arr, "Fizz")
		} else if i%5 == 0 {
			arr = append(arr, "Buzz")
		} else {
			arr = append(arr, strconv.Itoa(i))
		}
	}
	return arr
}