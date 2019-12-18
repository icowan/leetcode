/**
 * @Time : 2019-09-02 16:50
 * @Author : solacowa@gmail.com
 * @File : shiftingLetters
 * @Software: GoLand
 */

package main

import (
	"fmt"
)

func main() {

	// "mkgfzkkuxownxvfvxasy"
	// [505870226,437526072,266740649,224336793,532917782,311122363,567754492,595798950,81520022,684110326,137742843,275267355,856903962,148291585,919054234,467541837,622939912,116899933,983296461,536563513]
	//fmt.Println(shiftingLetters("bad", []int{10, 20, 30}))
	fmt.Println(shiftingLetters("mkgfzkkuxownxvfvxasy", []int{505870226, 437526072, 266740649, 224336793, 532917782, 311122363, 567754492, 595798950, 81520022, 684110326, 137742843, 275267355, 856903962, 148291585, 919054234, 467541837, 622939912, 116899933, 983296461, 536563513}))
}

func shiftingLetters(S string, shifts []int) string {
	var l = len(S) - 1
	var str string
	for i := 0; i <= l; i++ {
		var s = S[i]
		for n := i; n < len(shifts); n++ {
			s += uint8(shifts[n])
			s = isz(s)
		}
		str += string(s)
	}

	return str
}

func isz(s uint8) uint8 {
	if s > 122 {
		s = 97 + (s - 123)
		fmt.Println(s)
		return isz(s)
	}
	return s
}

func numSubarrayProductLessThanK(nums []int, k int) int {

	return 0
}
