package main4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main4() {
	sl_dias := [][]int{}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		name := scanner.Text()
		ints := make([]int, len(strings.Split(name, " ")))

		for i, s := range strings.Split(name, " ") {
			ints[i], _ = strconv.Atoi(s)
		}
		sl_dias = append(sl_dias, ints)

	}

	for i, dia := range sl_dias {

		// pegar nota maxima
		nota_max := 0
		for i, number := range dia {
			if i == 0 {
				_ = number

			} else if i%2 == 0 {
				if number > nota_max {
					nota_max = number
				}
			}
		}

		// pegar restaurantes com nota maxima
		restaurant_id := []int{}
		for i, number := range dia {
			if i%2 == 0 && i != 0 {
				if number == nota_max {
					restaurant_id = append(restaurant_id, dia[i-1])
				}
			}
		}

		fmt.Println("Dia", i+1)
		fmt.Println(findMin(restaurant_id))
		fmt.Println("")

	}

}

func findMin(a []int) (min int) {
	min = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
	}
	return min
}
