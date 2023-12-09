package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main3() {
	// Read the input from the command line.
	scanner := bufio.NewScanner(os.Stdin)

	// Read the first line of input.
	scanner.Scan()
	dimensions := scanner.Text()

	// Split the dimensions into two integers.
	_, height := parseDimensions(dimensions)

	// Read the map from the command line.
	mapData := make([][]rune, height)
	for i := 0; i < height; i++ {
		scanner.Scan()
		mapData[i] = []rune(scanner.Text())
	}

	// Read the number of trees from the command line.
	scanner.Scan()
	numTrees, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Count the number of visible trees.
	visibleTrees := countVisibleTrees(mapData, numTrees)

	// Print the output.
	fmt.Println(visibleTrees)
	printMap(mapData)
}

func parseDimensions(dimensions string) (int, int) {
	width, err := strconv.Atoi(dimensions[:strings.Index(dimensions, " ")])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	height, err := strconv.Atoi(dimensions[strings.Index(dimensions, " ")+1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return width, height
}

func countVisibleTrees(mapData [][]rune, numTrees int) int {
	visibleTrees := 0

	// Iterate over the map and count the number of visible trees.
	for i := 0; i < len(mapData); i++ {
		for j := 0; j < len(mapData[i]); j++ {
			if mapData[i][j] == 't' && isVisibleTree(mapData, i, j, numTrees) {
				visibleTrees++
			}
		}
	}

	return visibleTrees
}

func isVisibleTree(mapData [][]rune, i, j, numTrees int) bool {
	// Check if the tree is visible from the top.
	for k := i - 1; k >= 0; k-- {
		if mapData[k][j] != '#' {
			break
		}

		if k == 0 {
			return true
		}
	}

	// Check if the tree is visible from the left.
	for k := j - 1; k >= 0; k-- {
		if mapData[i][k] != '#' {
			break
		}

		if k == 0 {
			return true
		}
	}

	// Check if the tree is visible from the bottom.
	for k := i + 1; k < len(mapData); k++ {
		if mapData[k][j] != '#' {
			break
		}

		if k == len(mapData)-1 {
			return true
		}
	}

	// Check if the tree is visible from the right.
	for k := j + 1; k < len(mapData[i]); k++ {
		if mapData[i][k] != '#' {
			break
		}

		if k == len(mapData[i])-1 {
			return true
		}
	}

	return false
}

func printMap(mapData [][]rune) {
	for _, row := range mapData {
		for _, char := range row {
			fmt.Print(char)
		}

		fmt.Println()
	}
}
