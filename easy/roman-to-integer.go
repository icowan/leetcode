/**
 * @Time : 2019-09-02 16:08
 * @Author : solacowa@gmail.com
 * @File : roman-to-integer
 * @Software: GoLand
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	var r = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var num int
	var l = len(s) - 1
	for i := 0; i <= l; i++ {
		var left = r[string(s[i])]
		var right int
		if i+1 <= l {
			right = r[string(s[i+1])]
		}

		if left < right {
			num += right - left
			i++
		} else {
			num += left
		}
	}

	return num
}

// 如果左边的小于右边的 则用右边的 - 左边的
