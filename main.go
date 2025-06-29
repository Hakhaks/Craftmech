package main

import (
	"fmt"
)

// A CRUDE AWAKENING

const (
	gridW = 30
	gridH = 30
)

// ================
const (
	DirectionNone = iota
	DirectionLeft
	DirectionUp
	DirectionRight
	DirectionDown
	DirectionLast
)

var DirectionName = []string{"-", "<", "^", ">", "v"}
var DirectionDXY = [][]int{{0, 0}, {-1, 0}, {0, 1}, {1, 0}, {0, -1}}

// ================
type IGrid interface {
	W() int
	H() int
	Get(int, int) IComponent
	AddC(IComponent, []int, float64)
	GetC(IComponent, []int) IComponent
	AddXY(int, int, float64)
}

// ================
type Grid struct {
	c [gridW][gridH]IComponent
}

// ================
type IComponent interface {
	Step() bool
	ToString() string
	Add(float64)
	Direction() int
}

// ================
type Connection struct {
	gridRef    IGrid
	direction  []int // TODO: handle multiple
	energy     float64
	nextEnergy float64
}

func (c *Connection) Step() bool {
	c.energy = c.nextEnergy
	c.nextEnergy = 0

	if c.energy == 0 {
		return false
	}
	if c.direction != DirectionNone {
		c.gridRef.AddC(c, DirectionDXY[c.direction], c.energy)
		return true
	}
	// Look at neighbors
	// Prioritise those without direction
	valid := []int{}
	for i := DirectionNone + 1; i < DirectionLast; i++ {
		nb := c.gridRef.GetC(c, DirectionDXY[i])
		if nb != nil && nb.Direction() != DirectionNone {
			valid = append(valid, i)
		}
	}
	if len(valid) == 0 {
		for i := DirectionNone + 1; i < DirectionLast; i++ {
			nb := c.gridRef.GetC(c, DirectionDXY[i])
			if nb != nil {
				valid = append(valid, i)
			}
		}
	}
	if len(valid) > 0 {
	}
	return false
}

func (c *Connection) Add(energy float64) {
	c.nextEnergy += energy
}
func (c *Connection) Direction() int {
	return c.direction
}

func (c *Connection) ToString() string {
	if c.energy == 0 {
		return DirectionName[c.direction]
	}
	return fmt.Sprintf("%.1f", c.energy)
}

// ================
type Generator struct {
}

func (g *Generator) Step() bool {
	return false
}

func (g *Generator) ToString() string {
	return "1"
}

// ================
func main() {
	// Init grid
	grid := make([][]IComponent, gridW)
	for x := range grid {
		grid[x] = make([]IComponent, gridH)
		for y := range grid[x] {
			grid[x][y] = &Connection{}
		}
	}
	// Print grid
	printGrid(grid)
	// TODO:
	// Flood fill
	// Simulate grid
}

func printGrid(grid [][]IComponent) {
	for x := range grid {
		for y := range grid[x] {
			fmt.Printf(" %v ", grid[x][y].ToString())
		}
		fmt.Println("")
	}
}
