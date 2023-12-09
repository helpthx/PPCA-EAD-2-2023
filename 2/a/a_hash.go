package main

import (
	"fmt"
)

func main() {
	// Create a concurrent hash table.

	// Read the int input.
	var size int
	fmt.Scanln(&size)
	hashTable := make(map[int]struct{})
	//var elemenetList []int

	for i := 0; i < size; i++ {
		var element int
		fmt.Scanln(&element)
		hashTable[element] = struct{}{}
		//elemenetList = append(elemenetList, element)
	}

	// Extract list from map
	if size%2 != 0 {
		hashTable[1000000000] = struct{}{}
		//elemenetList = append(elemenetList, 1000000000)
	}
	elements := make([]int, 0, len(hashTable))
	for k := range hashTable {
		elements = append(elements, k)
	}

	elements = mergeSort(elements)

	var size_output = len(elements)

	var hashTableF = ArraySum(elements, size_output, hashTable)

	var sizeF = len(hashTableF)

	elements_final := make([]int, 0, len(hashTableF))
	for k := range hashTableF {
		elements_final = append(elements_final, k)
	}

	elements_final = mergeSort(elements_final)

	for i := 0; i < sizeF; i = i + 4 {
		fmt.Println(elements_final[i])
	}

	fmt.Println("Elementos:", sizeF)
}

func ArraySum(arr []int, n int, hashTable map[int]struct{}) map[int]struct{} {
	for i := 0; i+1 < n; i = i + 2 {
		hashTable[arr[i]+arr[i+1]] = struct{}{}
	}
	if n == 1 {
		hashTable[arr[0]+arr[1]] = struct{}{}
	}

	return hashTable
}

func mergeSort(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
