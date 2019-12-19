/**
 * @Time : 19/12/2019 1:53 PM
 * @Author : solacowa@gmail.com
 * @File : main
 * @Software: GoLand
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	//程序 07：输出特殊图案
	//题目：输出特殊图案，请在 c 环境中运行，看一看，Very Beautiful!
	//1.程序分析：字符共有 256 个。不同字符，图形不一样。
	//2.程序源代码：
	var a, b = 176, 219
	fmt.Printf("%c%c%c%c%c \n", b, a, a, a, b)
	fmt.Printf("%c%c%c%c%c \n", a, b, a, b, a)
	fmt.Printf("%c%c%c%c%c \n", a, a, b, a, a)
	fmt.Printf("%c%c%c%c%c \n", a, b, a, b, a)
	fmt.Printf("%c%c%c%c%c \n", b, a, a, a, b)

	//程序 08：9*9 口诀
	//题目：输出 9*9 口诀。
	//1.程序分析：分行与列考虑，共 9 行 9 列，i 控制行，j 控制列。
	//2.程序源代码：
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%-3d", j, i, i*j) /*-3d 表示左对齐，占 3 位*/
		}
		fmt.Printf("\n")
	}

	//程序 09：象棋棋盘
	//题目：要求输出国际象棋棋盘。
	//1.程序分析：用 i 控制行， j 来控制列，根据i+j的和的变化来控制输出黑方格，还是白方格。
	//2.程序源代码：
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 {
				fmt.Printf("■")
			} else {
				fmt.Printf("□")
			}
		}
		fmt.Printf("\n")
	}

	//程序 10：打印楼梯
	//题目：打印楼梯，同时在楼梯上方打印笑脸。
	//1.程序分析：用 i 控制行，j 来控制列，j 根据 i 的变化来控制输出黑方格的个数。
	//2.程序源代码：
	for i := 0; i < 8; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("■■")
		}
		fmt.Printf("☺\n")
	}

	// 程序 11：兔子问题
	// 题目：古典问题：有一对兔子，从出生后第3个月起每个月都生一对兔子，小兔子长到第三个月后每个月又生一对兔子，假如兔子都不死，问每个月的兔子总数为多少？
	// 1.程序分析： 兔子的规律为数列 1,1,2,3,5,8,13,21....
	// 2.程序源代码：
	var m1, m2 int = 1, 1
	for i := 1; i < 12; i++ {
		fmt.Println(m1, m2)
		m1 += m2
		m2 += m1
	}

	// 程序 12：求素数
	//题目：判断 101-200 之间有多少个素数，并输出所有素数。
	//1.程序分析：判断素数的方法：用一个数分别去除2到sqrt(这个数)，如果能被整除，则表明此数不是素数，反之是素数。
	//2.程序源代码：

	var i, j, k, count int = 0, 0, 0, 0
	for i = 101; i <= 200; i++ {
		k = int(math.Sqrt(float64(i)))
		for j = 2; j <= k; j++ {
			if i%j == 0 {
				break
			}
		}
		if j == k+1 {
			fmt.Println(i)
			count++
		}
	}
	fmt.Println("total", count)

	//程序 13：水仙花数
	//题目：打印出所有的“水仙花数”，所谓“水仙花数”是指一个三位数，其各位数字立方和等于该数本身。例如：153 是一个“水仙花数”，因为 153=1 的三次方＋5 的三次方＋3 的三次方。
	//1.程序分析：利用for循环控制100-999个数，每个数分解出个位，十位，百位。
	//2.程序源代码：
	for num := 100; num < 1000; num++ {
		i := num / 100
		j := num / 10 % 10
		k := num % 10
		if i*i*i+j*j*j+k*k*k == num {
			fmt.Printf("%d^3+%d^3+%d^3=%d\n", i, j, k, num)
		}
	}

}
