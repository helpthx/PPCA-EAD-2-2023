package main

import (
	"fmt"
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
	// Create a concurrent hash table.

	// Read the int input.
	var n int
	fmt.Scanln(&n)
	//type void struct{}
	//var member void
	//hashTable := make(map[int]void)
	var elemenetList []int

	for i := 0; i < n; i++ {
		var element int
		fmt.Scanln(&element)
		//hashTable[element] = member
		elemenetList = append(elemenetList, element)
	}

	elemenetList = quickSortStart(elemenetList)

	elemenetList, size := removeDuplicates(elemenetList)

	// Extract list from map
	if size%2 != 0 {
		elemenetList = append(elemenetList, 1000000000)
	}

	// sizeCrazy := 0
	// for i := 0; i < size-1; i += 2 {
	// 	sizeCrazy = sizeCrazy + 1
	// 	elemenetList[size+sizeCrazy] = elemenetList[i] + elemenetList[i+1]
	// }

	elemenetList, size = ArraySum(elemenetList, size)

	//elemenetList, size = removeDuplicates(elemenetList)

	for i := 0; i < size; i = i + 4 {
		fmt.Println(elemenetList[i])
	}

	fmt.Println("Elementos:", size)
}

func ArraySum(arr []int, n int) ([]int, int) {
	var output []int
	for i := 0; i < n; i = i + 2 {
		sumVar := arr[i] + arr[i+1]
		if !binarySearch(sumVar, arr) {
			output = append(output, arr[i]+arr[i+1])
		}
	}
	if n == 1 {
		output = append(output, arr[0]+arr[1])
	}
	final := append(arr, output...)
	return final, len(final)
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func removeDuplicates(slice []int) ([]int, int) {
	// Create a new slice to store the unique values.
	var uniqueSlice []int

	// Iterate over the slice and add each unique value to the new slice.
	previousValue := 0
	for _, value := range slice {
		if value != previousValue {
			uniqueSlice = append(uniqueSlice, value)
			previousValue = value
		}
	}

	return uniqueSlice, len(uniqueSlice)
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
