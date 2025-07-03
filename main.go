package main

import (
	"fmt"
)

// A CRUDE AWAKENING

const (
	gridW   = 30
	gridH   = 30
	maxIter = 10
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
	Step() bool
	Get(int, int) IComponent
	AddC(IComponent, []int, float64)
	AddXY(int, int, float64)
}

// ================
type Grid struct {
	w, h int
	c    [][]IComponent
}

func MakeGrid(w, h int) *Grid {
	g := &Grid{w: w, h: h}
	g.c = make([][]IComponent, gridW)
	for x := range gridW {
		g.c[x] = make([]IComponent, gridH)
	}
	return g
}

func (g *Grid) Step() bool {
	anyChange := false
	for x := range g.w {
		for y := range g.w {
			c := g.c[x][y]
			if c != nil {
				anyChange = c.Step() || anyChange
			}
		}
	}
	return anyChange
}

func (g *Grid) AddC(source IComponent, directions []int, energy float64) {

	// Add energy from source to directions
	for di := range directions {
		d := DirectionDXY[directions[di]]
		g.AddXY(source.X()+d[0], source.Y()+d[1], energy/float64(len(directions)))
	}
}

func (g *Grid) AddXY(x int, y int, energy float64) {
	g.Get(x, y).Add(energy)
}

func (g *Grid) Get(x int, y int) IComponent {
	if !g.InBounds(x, y) {
		return nil
	}
	return g.c[x][y]
}

func (g *Grid) InBounds(x int, y int) bool {
	return x >= 0 && y >= 0 && x < g.w && y < g.h
}

func (g *Grid) H() int { return g.h }
func (g *Grid) W() int { return g.w }

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
	grid := MakeGrid(gridW, gridH)
	// TODO: add components
	// TODO: add initial energy
	for iter := 0; iter < maxIter; iter++ {
		printGrid(grid, iter)
		if !grid.Step() {
			fmt.Println("No further grid changes; Terminating simulation")
			break
		}
	}
}

func printGrid(grid IGrid, iter int) {
	fmt.Printf("Grid, iter %v\n", iter)
	for x := range gridW {
		for y := range gridH {
			c := grid.Get(x, y)
			str := "|   "
			if c != nil {
				str = fmt.Sprintf(" %v ", c.ToString())
			}
			fmt.Print(str)
		}
		fmt.Println("")
	}
}

func Assert(b bool) {
	if !b {
		panic("Assertion failed")
	}
}
