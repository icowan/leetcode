/**
 * @Time : 2019-09-27 09:48
 * @Author : solacowa@gmail.com
 * @File : quick-sort
 * @Software: GoLand
 */

package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func quickSort(arr []int64) {

}

func main() {

	f, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = f.Close()
	}()

	w := csv.NewWriter(f)
	var data [][]string
	for i := 90000000001; i <= 90000100000; i++ {
		data = append(data, []string{
			strconv.Itoa(i),
			"syswin123",
		})
	}

	_ = w.WriteAll(data)
	w.Flush()
}
