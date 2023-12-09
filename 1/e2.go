package main

import (
	"fmt"
	"os"
)

type Tower struct {
	Height int
	Width  int
	Blocks [][]rune
}

func NewTower(height, width int) *Tower {
	return &Tower{
		Height: height,
		Width:  width,
		Blocks: make([][]rune, height),
	}
}

func (t *Tower) ReadInput() {
	for i := 0; i < t.Height; i++ {
		t.Blocks[i] = []rune(os.Args[1+i])
	}
}

func (t *Tower) CanSee(x, y int) bool {
	// Check if the block is empty.
	if t.Blocks[y][x] != '#' {
		return true
	}

	// Check if the block is on the edge of the tower.
	if x == 0 || x == t.Width-1 || y == 0 || y == t.Height-1 {
		return true
	}

	// Check if the block is visible from above.
	for i := y - 1; i >= 0; i-- {
		if t.Blocks[i][x] != '#' {
			return true
		}
	}

	// Check if the block is visible from below.
	for i := y + 1; i < t.Height; i++ {
		if t.Blocks[i][x] != '#' {
			return true
		}
	}

	// Check if the block is visible from the left.
	for i := x - 1; i >= 0; i-- {
		if t.Blocks[y][i] != '#' {
			return true
		}
	}

	// Check if the block is visible from the right.
	for i := x + 1; i < t.Width; i++ {
		if t.Blocks[y][i] != '#' {
			return true
		}
	}

	return false
}

func (t *Tower) CountVisibleBlocks() int {
	count := 0
	for y := 0; y < t.Height; y++ {
		for x := 0; x < t.Width; x++ {
			if t.CanSee(x, y) {
				count++
			}
		}
	}

	return count
}

func main52() {
	// Read the input.
	height := 5
	width := 10
	tower := NewTower(height, width)
	tower.ReadInput()

	// Calculate the number of visible blocks.
	visibleBlocks := tower.CountVisibleBlocks()

	// Print the output.
	fmt.Println(visibleBlocks)
}
