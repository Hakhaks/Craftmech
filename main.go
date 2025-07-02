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
	c [][]IComponent
}

func MakeGrid(w, h int) *Grid {
	g := &Grid{}
	g.c = make([][]IComponent, gridW)
	for x := range gridW {
		g.c[x] = make([]IComponent, gridH)
	}
	return g
}

// AddC implements IGrid.
func (g *Grid) AddC(IComponent, []int, float64) {
	// TODO:
	panic("unimplemented")
}

// AddXY implements IGrid.
func (g *Grid) AddXY(int, int, float64) {
	// TODO:
	panic("unimplemented")
}

// Get implements IGrid.
func (g *Grid) Get(int, int) IComponent {
	// TODO:
	panic("unimplemented")
}

// GetM implements IGrid.
func (g *Grid) GetM(IComponent, []int) []IComponent {
	// TODO:
	panic("unimplemented")
}

// H implements IGrid.
func (g *Grid) H() int {
	// TODO:
	panic("unimplemented")
}

// W implements IGrid.
func (g *Grid) W() int {
	// TODO:
	panic("unimplemented")
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
	directions []int
	energy     float64
	nextEnergy float64
}

func (c *Connection) Step() bool {
	c.energy = c.nextEnergy
	c.nextEnergy = 0

	if c.energy == 0 {
		return false
	}
	if len(c.directions) == 0 {
		// Look for neighbors to set direction
		// Prioritise those without direction
		for i := DirectionNone + 1; i < DirectionLast; i++ {
			nb := c.gridRef.Get(c.X()+DirectionDXY[i][0], c.Y()+DirectionDXY[i][1])
			if nb != nil && len(nb.Directions()) > 0 {
				c.directions = append(c.directions, i)
			}
		}
		// Look for any neighbor, even those with direction
		if len(c.directions) == 0 {
			for i := DirectionNone + 1; i < DirectionLast; i++ {
				nb := c.gridRef.Get(c.X()+DirectionDXY[i][0], c.Y()+DirectionDXY[i][1])
				if nb != nil {
					c.directions = append(c.directions, i)
				}
			}
		}
	}
	// With direction set, neighbors for energy transfer is determined
	if len(c.directions) > 0 {
		c.gridRef.AddC(c, c.directions, c.energy)
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
	grid := MakeGrid(gridW, gridH)
	// Print grid
	printGrid(grid)
	// TODO:
	// Flood fill
	// Simulate grid
}

func printGrid(grid IGrid) {
	for x := range gridW {
		for y := range gridH {
			fmt.Printf(" %v ", grid.Get(x, y).ToString())
		}
		fmt.Println("")
	}
}

func Assert(b bool) {
	if !b {
		panic("Assertion failed")
	}
}
