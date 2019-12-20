/**
 * @Time : 20/12/2019 3:13 PM
 * @Author : solacowa@gmail.com
 * @File : first-unique-character-in-a-string
 * @Software: GoLand
 */

package main

import (
	"fmt"
	"strings"
)

//387
//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
//案例:
//s = "leetcode"
//返回 0.
//s = "loveleetcode",
//返回 2.
//注意事项：您可以假定该字符串只包含小写字母。
func firstUniqChar(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return -1
	}
	if len(s) == 1 {
		return 0
	}

	if len(s) == 2 {
		if s[0:1] == s[1:2] {
			return -1
		}
		return 0
	}
	if strings.Contains(s, s[0:1]) {

	}

	for k, v := range s {
		var isExists bool
		for _, vv := range s[k+1:] {
			if v == vv {
				isExists = true
				break
			}
		}
		if !isExists {
			return k
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("aadadaad"))
}
