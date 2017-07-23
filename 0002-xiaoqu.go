package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i int64
	xiaoqu := os.Args[1]
	j, _ := strconv.ParseInt(os.Args[2], 0, 32)
	for i = 1; i <= j; i++ {
		fmt.Printf("https://bj.lianjia.com/xiaoqu/%s/pg%d/\n", xiaoqu, i)
	}
}
