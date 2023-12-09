package main

import (
	"fmt"
	"sort"
)

func main() {
	// Read the inputs from the terminal.
	var size int
	fmt.Scanln(&size)

	var inputs []int
	for i := 0; i < size; i++ {
		var element int
		fmt.Scanln(&element)
		inputs = append(inputs, element)
		//elemenetList = append(elemenetList, element)
	}

	// Remove duplicate numbers from the inputs.
	uniqueInputs := make(map[int]struct{})
	for _, input := range inputs {
		uniqueInputs[input] = struct{}{}
	}

	elements := make([]int, 0, len(uniqueInputs))
	for k := range uniqueInputs {
		elements = append(elements, k)
	}
	// Sort the unique inputs.

	// Add 1000 to the end of the unique inputs if the amount is not even.
	if len(elements)%2 != 0 {
		elements = append(elements, 1000000000)
	}
	sort.Ints(elements)

	// Sum the next numbers in the unique inputs.
	var newInputs []int
	for i := 0; i < len(uniqueInputs); i += 2 {
		newInputs = append(newInputs, elements[i]+elements[i+1])
	}
	//fmt.Println(elements)
	//fmt.Println(newInputs)
	unionSlice := union(elements, newInputs)
	// Print the new inputs.
	// for _, input := range newInputs {
	// 	uniqueInputs[input] = struct{}{}
	// }

	// var n = len(uniqueInputs)
	// final := make([]int, 0, n)
	// for k := range uniqueInputs {
	// 	final = append(elements, k)
	// }

	// // Sort the unique inputs.
	// sort.Ints(final)

	// n = len(elements)

	// fmt.Println(elements)
	//fmt.Println(unionSlice)
	n := len(unionSlice)

	for i := 0; i < n; i = i + 4 {
		fmt.Println(unionSlice[i])
	}

	fmt.Println("Elementos:", n)
}

func union(slice1 []int, slice2 []int) []int {
	// Sort the two slices.
	//sort.Ints(slice1)
	//sort.Ints(slice2)

	// Create a new slice to store the union of the two slices.
	unionSlice := make([]int, 0)

	// Iterate over the two slices and add each unique element to the union slice.
	i := 0
	j := 0
	for i < len(slice1) && j < len(slice2) {
		if slice1[i] < slice2[j] {
			unionSlice = append(unionSlice, slice1[i])
			i++
		} else if slice1[i] > slice2[j] {
			unionSlice = append(unionSlice, slice2[j])
			j++
		} else {
			// The elements are equal, so only add one to the union slice.
			unionSlice = append(unionSlice, slice1[i])
			i++
			j++
		}
	}

	// Add any remaining elements from the first slice to the union slice.
	for i < len(slice1) {
		unionSlice = append(unionSlice, slice1[i])
		i++
	}

	// Add any remaining elements from the second slice to the union slice.
	for j < len(slice2) {
		unionSlice = append(unionSlice, slice2[j])
		j++
	}

	return unionSlice
}
