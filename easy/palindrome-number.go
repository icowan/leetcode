/**
 * @Time : 2019-09-02 16:00
 * @Author : solacowa@gmail.com
 * @File : palindrome-number
 * @Software: GoLand
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(isPalindrome(12321))
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	var l = len(s) - 1
	for i := 0; i < l; i++ {
		if s[i] != s[l-i] {
			return false
		}
	}
	return true
}

//func isPalindrome(x int) bool {
//	if x < 0 {
//		return false
//	}
//	var r int = 0
//	t := x
//	for t != 0  {
//		r = r * 10 + t % 10
//		t /= 10
//	}
//	if r == x {
//		return true
//	} else {
//		return false
//	}
//}
