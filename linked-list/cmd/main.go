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
	"strconv"
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
	{
		for i := 1; i < 10; i++ {
			for j := 1; j <= i; j++ {
				fmt.Printf("%d*%d=%-3d", j, i, i*j) /*-3d 表示左对齐，占 3 位*/
			}
			fmt.Printf("\n")
		}
	}

	//程序 09：象棋棋盘
	//题目：要求输出国际象棋棋盘。
	//1.程序分析：用 i 控制行， j 来控制列，根据i+j的和的变化来控制输出黑方格，还是白方格。
	//2.程序源代码：
	{
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
	}
	//程序 10：打印楼梯
	//题目：打印楼梯，同时在楼梯上方打印笑脸。
	//1.程序分析：用 i 控制行，j 来控制列，j 根据 i 的变化来控制输出黑方格的个数。
	//2.程序源代码：
	{
		for i := 0; i < 8; i++ {
			for j := 0; j <= i; j++ {
				fmt.Printf("■■")
			}
			fmt.Printf("☺\n")
		}
	}

	// 程序 11：兔子问题
	// 题目：古典问题：有一对兔子，从出生后第3个月起每个月都生一对兔子，小兔子长到第三个月后每个月又生一对兔子，假如兔子都不死，问每个月的兔子总数为多少？
	// 1.程序分析： 兔子的规律为数列 1,1,2,3,5,8,13,21....
	// 2.程序源代码：
	{
		var m1, m2 int = 1, 1
		for i := 1; i < 12; i++ {
			fmt.Println(m1, m2)
			m1 += m2
			m2 += m1
		}
	}

	// 程序 12：求素数
	//题目：判断 101-200 之间有多少个素数，并输出所有素数。
	//1.程序分析：判断素数的方法：用一个数分别去除2到sqrt(这个数)，如果能被整除，则表明此数不是素数，反之是素数。
	//2.程序源代码：

	{
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
	}

	//程序 13：水仙花数
	//题目：打印出所有的“水仙花数”，所谓“水仙花数”是指一个三位数，其各位数字立方和等于该数本身。例如：153 是一个“水仙花数”，因为 153=1 的三次方＋5 的三次方＋3 的三次方。
	//1.程序分析：利用for循环控制100-999个数，每个数分解出个位，十位，百位。
	//2.程序源代码：
	{
		for num := 100; num < 1000; num++ {
			i := num / 100
			j := num / 10 % 10
			k := num % 10
			if i*i*i+j*j*j+k*k*k == num {
				fmt.Printf("%d^3+%d^3+%d^3=%d\n", i, j, k, num)
			}
		}
	}

	//程序 14：分解质因数
	//题目：将一个正整数分解质因数。例如：输入 90,打印出 90=2*3*3*5。
	//程序分析：对n进行分解质因数，应先找到一个最小的质数k，然后按下述步骤完成：
	//(1)如果这个质数恰等于n，则说明分解质因数的过程已经结束，打印出即可。
	//(2)如果n<>k，但n能被k整除，则应打印出k的值，并用n除以k的商,作为新的正整数n,重复执行第一步。
	//(3)如果n不能被k整除，则用k+1作为k的值, 重复执行第一步。
	//2.程序源代码：
	{
		var n, i int = 0, 0
		fmt.Printf("please input a number:")
		fmt.Scanf("%d\n", &n) // 90
		fmt.Printf("%d=", n)
		for i = 2; i <= n; i++ {
			for n != i {
				if n%i == 0 {
					fmt.Printf("%d*", i)
					n = n / i
					// i = 2, n = 45
					// i = 3, n = 15
					// i = 3, n = 5
					// i = 5, n = 1
				} else {
					break
				}
			}
		}
		fmt.Printf("%d\n", n)
	}

	//程序 15：三元表达式
	//题目：利用条件运算符的嵌套来完成此题：学习成绩>=90 分的同学用 A 表示，60-89 分之
	//间的用 B 表示，60分以下的用C表示。
	//1.程序分析：(a>b)?a:b 这是条件运算符的基本例子。
	//2.程序源代码：
	{
		var score int = 0
		fmt.Scanf("%d", &score)
		fmt.Println(B(score >= 90).C("A", B(score >= 60).C("B", "C")))
	}

	//程序 16：最大公约数和最小公倍数
	//题目：输入两个正整数 m 和 n，求其最大公约数和最小公倍数。
	//1.程序分析：利用辗除法。
	//2.程序源代码：
	{
		var m, n, r, x int
		fmt.Scanf("%d%d", &m, &n)
		x = m * n
		for n != 0 {
			r = m % n
			m = n
			n = r
		}
		fmt.Printf("%d %d\n", m, x/m)
	}
}

type B bool

//Go语言没有三元表达式,自写函数模仿。
func (b B) C(t, f interface{}) interface{} {
	if bool(b) == true {
		return t
	}
	return f
}

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
