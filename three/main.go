package main

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

var grid map[uint32]*Point

// Point holds data about a point on a grid.
type Point struct {
	X        int
	Y        int
	HasWire1 bool
	HasWire2 bool
}

func main() {
	initGrid()

	// get the intersections
	var intersections []*Point

	for _, p := range grid {
		if p.HasWire1 && p.HasWire2 {
			intersections = append(intersections, p)
		}
	}

	// find the most smol (more efficient to do this above but whatevs)
	var smol int

	for idx, i := range intersections {
		d := manhattanDistance(i.X, i.Y)

		if idx == 0 || d < smol {
			smol = d
		}
	}

	log.Print(smol)
}

func initGrid() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	grid = make(map[uint32]*Point)

	lines := strings.Split(string(b), "\n")

	plotLines(strings.Split(lines[0], ","), true)
	plotLines(strings.Split(lines[1], ","), false)
}

func manhattanDistance(x, y int) int {
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func plotLines(directions []string, isOne bool) {
	// set current X and Y starting points, these will be used as a cursor in the below
	// loop.
	cX, cY := 0, 0

	for _, d := range directions {
		direction := d[0:1]
		distance, err := strconv.Atoi(d[1:])
		if err != nil {
			panic(err)
		}
		// log.Printf("%s:%d", direction, distance)
		cX, cY = plotLine(cX, cY, direction, distance, isOne)
	}
}

// returns new cursor position
func plotLine(curX, curY int, direction string, distance int, isOne bool) (int, int) {
	for i := 0; i < distance; i++ {
		switch direction {
		case "R":
			curX++
		case "L":
			curX--
		case "U":
			curY++
		case "D":
			curY--
		default:
			panic("FUCK!")
		}
		addPoint(curX, curY, isOne)
	}

	return curX, curY
}

func addPoint(x, y int, isOne bool) {
	h := getXYHash(x, y)

	// log.Printf("%d %d,%d added", h, x, y)

	// Add this point to the grid if it is not already present.
	if _, ok := grid[h]; !ok {
		grid[h] = &Point{X: x, Y: y}
	}

	if isOne {
		grid[h].HasWire1 = true
		return
	}

	grid[h].HasWire2 = true
	return
}

// using https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function for speed since
// we don't need any actual crypto functionality
func getXYHash(x, y int) uint32 {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%d,%d", x, y)))
	return h.Sum32()
}
