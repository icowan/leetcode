/**
 * @Time : 2019-09-02 15:08
 * @Author : solacowa@gmail.com
 * @File : twosum
 * @Software: GoLand
 */

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

链接：https://leetcode-cn.com/problems/two-sum
*/

package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

// 两数之合
func twoSum(nums []int, target int) []int {
	var sum []int

F:
	for k, v := range nums {
		if v+nums[k+1] == target {
			sum = append(sum, k, k+1)
			break
		}
		for key, val := range nums[k+1:] {
			if v+val == target {
				sum = append(sum, k, k+1+key)
				break F
			}
		}
	}

	return sum
}

//func twoSum(nums []int, target int) []int {
//	tmpMap := make(map[int]int, 0)
//	for key, value := range nums {
//		tmp := target - value
//		if i, ok := tmpMap[tmp]; ok {
//			return []int{i, key}
//		}
//		tmpMap[value] = key
//	}
//	return []int{}
//}
