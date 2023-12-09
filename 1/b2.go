package main

import (
	"fmt"
)

func main4() {
	var n int
	var sum int

	sum = 0

	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		var num int
		fmt.Scanln(&num)
		sum += num
	}
	fmt.Println(sum)
}
