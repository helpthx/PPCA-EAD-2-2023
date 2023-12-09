package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var contagem int
	fmt.Scanln(&contagem)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	name := scanner.Text()
	numeros_proibidos := make([]int, contagem)

	numeros_proibidos_str := strings.Split(name, " ")
	for i, s := range numeros_proibidos_str {
		numeros_proibidos[i], _ = strconv.Atoi(s)
	}

	numeros := []int{}

	for scanner.Scan() {
		name := scanner.Text()
		var valores int
		valores, _ = strconv.Atoi(name)
		numeros = append(numeros, valores)

	}

	for _, numero := range numeros {
		if contains(numeros_proibidos, numero) {
			fmt.Println("sim")
		} else {
			fmt.Println("nao")
		}

	}

}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
