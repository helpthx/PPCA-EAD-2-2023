package main2

import (
	"bufio"
	"fmt"
	"os"
)

func main2() {
	scanner := bufio.NewScanner(os.Stdin)

	var lineCount int

	for scanner.Scan() {
		lineCount++
	}
	fmt.Println(lineCount)
}
