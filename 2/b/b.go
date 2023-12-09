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

func quickselect(slice []int, k int) int {
    if k < 0 || k >= len(slice) {
        panic("k must be between 0 and the length of the slice")
    }

    left := 0
    right := len(slice) - 1
    pivot := partitionSelect(slice, left, right)

    for pivot != k {
        if pivot < k {
            left = pivot + 1
        } else {
            right = pivot - 1
        }

        pivot = partitionSelect(slice, left, right)
    }

	return slice[pivot]
}

func partitionSelect(slice []int, left int, right int) int {
    pivot := slice[right]

    i := left - 1
    for j := left; j < right; j++ {
        if slice[j] <= pivot {
            i++
            temp := slice[i]
            slice[i] = slice[j]
            slice[j] = temp
        }
    }

    temp := slice[i+1]
    slice[i+1] = slice[right]
    slice[right] = temp

    return i + 1
}

func main() {
	// Read the int input.
	var n, p, x int
	fmt.Scanln(&n, &p, &x)

	var v []int
	for i := 0; i < n; i++ {
		var element int
		fmt.Scanln(&element)
		v = append(v, element)
	}

	_ = quickselect(v, min(n-1, p*x))
	_ = quickselect(v, min(n-1, (p+1)*(x-1)))
	v = quickSort(v, min(n-1, p*x), min(n-1, (p+1)*(x-1)))

	for i := p*x; i < min(n, (p+1)*(x)); i++ {
		fmt.Println(v[i])
	}
}


