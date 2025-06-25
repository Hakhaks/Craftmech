package main

import (
	"fmt"
)

const (
	gridW = 30
	gridH = 30
)

func main() {
	// A CRUDE AWAKENING
	// TODO:
	// Init grid
	grid := make([][]int, gridW)
	for x := range grid {
		grid[x] = make([]int, gridH)
	}
	// Print grid
	printGrid(grid)
	// Add components
	//    Generator
	//    Some connections to generator
	// Implement flood fill
	// Simulate inventory
}

func printGrid(grid [][]int) {
	for x := range grid {
		for y := range grid[x] {
			fmt.Printf(" %v ", grid[x][y])
		}
		fmt.Println("")
	}
}
