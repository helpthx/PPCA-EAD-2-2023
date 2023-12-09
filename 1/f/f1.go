package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func main() {
	var contagem int
	fmt.Scanln(&contagem)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	name := scanner.Text()
	numeros_proibidos := []int{}

	numeros_proibidos_str := strings.Split(name, " ")

	for _, s := range numeros_proibidos_str {
		valor, _ := strconv.Atoi(s)
		numeros_proibidos = append(numeros_proibidos, valor)
	}

	numeros := []int{}

	for scanner.Scan() {
		name := scanner.Text()
		var valores int
		valores, _ = strconv.Atoi(name)
		numeros = append(numeros, valores)

	}

	numeros_proibidos = quickSortStart(numeros_proibidos)

	for i := 0; i < len(numeros); i++ {
		if binarySearch(numeros[i], numeros_proibidos){
			fmt.Println("sim")
		} else {
			fmt.Println("nao")
		}
	}
}

func binarySearch(needle int, haystack []int) bool {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}
