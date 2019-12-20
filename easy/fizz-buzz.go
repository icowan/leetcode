/**
 * @Time : 20/12/2019 3:11 PM
 * @Author : solacowa@gmail.com
 * @File : fizz-buzz
 * @Software: GoLand
 */

package main

import "strconv"

//412
//写一个程序，输出从 1 到 n 数字的字符串表示。
//1. 如果 n 是3的倍数，输出“Fizz”；
//2. 如果 n 是5的倍数，输出“Buzz”；
//3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。
func fizzBuzz(n int) []string {
	var rs []string
	for i := 1; i <= n; i++ {
		var s string
		if i%3 == 0 {
			s = "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		if i%3 != 0 && i%5 != 0 {
			s = strconv.Itoa(i)
		}
		rs = append(rs, s)
	}
	return rs
}
