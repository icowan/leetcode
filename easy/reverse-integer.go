/**
 * @Time : 2019-09-02 15:27
 * @Author : solacowa@gmail.com
 * @File : reverse-integer
 * @Software: GoLand
 */

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(reverse(-2147483412))
}

func reverse(x int) int {
	if x < 10 && x >= 0 {
		return x
	}
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	var neg bool
	var s string
	if x < 0 {
		neg = true
		s = strconv.Itoa(x - x - x)
	} else {
		s = strconv.Itoa(x)
	}

	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}

	if neg {
		res = "-" + res
	}
	i, _ := strconv.Atoi(res)

	if i > math.MaxInt32 || i < math.MinInt32 {
		return 0
	}

	return i
}

func reverse2(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		if !(-(1<<31) <= y && y <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return y
}
