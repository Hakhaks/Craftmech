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
	GetM(IComponent, []int) []IComponent
	AddXY(int, int, float64)
}

// ================
type Grid struct {
	c [gridW][gridH]IComponent
}

// ================
type IComponent interface {
	X() int
	Y() int
	Step() bool
	ToString() string
	Add(float64)
	Directions() []int
}

// ================
type Connection struct {
	gridRef    IGrid
	x, y       int
	directions []int // TODO: handle multiple
	energy     float64
	nextEnergy float64
}

func (c *Connection) Step() bool {
	c.energy = c.nextEnergy
	c.nextEnergy = 0

	if c.energy == 0 {
		return false
	}
	if len(c.directions) > 0 {
		c.gridRef.AddC(c, c.directions, c.energy)
		return true
	}
	// TODO: Sinks
	// Look for neighbors to transfer energy
	// Prioritise those without direction
	nbs := []IComponent{}
	for i := DirectionNone + 1; i < DirectionLast; i++ {
		nb := c.gridRef.Get(c.X()+DirectionDXY[i][0], c.Y()+DirectionDXY[i][1])
		if nb != nil && len(nb.Directions()) > 0 {
			nbs = append(nbs, nb)
		}
	}
	// Look for any neighbor, even those with direction
	if len(nbs) == 0 {
		for i := DirectionNone + 1; i < DirectionLast; i++ {
			nb := c.gridRef.Get(c.X()+DirectionDXY[i][0], c.Y()+DirectionDXY[i][1])
			if nb != nil {
				nbs = append(nbs, nb)
			}
		}
	}
	// Add to neighbors
	if len(nbs) > 0 {
		for _, nb := range nbs {
			nb.Add(c.energy / float64(len(nbs)))
		}
		return true
	}
	return false
}

func (c *Connection) Add(energy float64) {
	c.nextEnergy += energy
}

func (c *Connection) Directions() []int {
	return c.directions
}

func (c *Connection) X() int {
	return c.x
}

func (c *Connection) Y() int {
	return c.y
}

func (c *Connection) ToString() string {
	if c.energy == 0 {
		ret := ""
		for i := range c.directions {
			ret += DirectionName[c.directions[i]]
		}
		return ret
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
