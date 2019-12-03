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

var grid map[uint32]*point

type (
	wire  int
	point struct {
		X        int
		Y        int
		Wires    []wire
		Distance int // manhattan distance from origin
	}
)

// constants
const (
	WIRE1 wire = 1
	WIRE2 wire = 2
)

func main() {
	initGrid()

	// get the intersections
	var intersections []*point

	for _, p := range grid {
		if len(p.Wires) > 1 {
			intersections = append(intersections, p)
		}
	}

	// find the most smol (more efficient to do this above but whatevs)
	var smol int

	for idx, i := range intersections {
		if idx == 0 || i.Distance < smol {
			smol = i.Distance
		}
	}

	log.Print(smol)
}

func initGrid() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	grid = make(map[uint32]*point)

	lines := strings.Split(string(b), "\n")

	plotLines(strings.Split(lines[0], ","), WIRE1)
	plotLines(strings.Split(lines[1], ","), WIRE2)
}

func manhattanDistance(x, y int) int {
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func plotLines(directions []string, w wire) {
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
		cX, cY = plotLine(cX, cY, direction, distance, w)
	}
}

// returns new cursor position
func plotLine(curX, curY int, direction string, distance int, w wire) (int, int) {
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
		addPoint(curX, curY, w)
	}

	return curX, curY
}

func addPoint(x, y int, w wire) {
	h := getXYHash(x, y)

	// Add this point to the grid if it is not already present.
	if _, ok := grid[h]; !ok {
		grid[h] = &point{X: x, Y: y, Distance: manhattanDistance(x, y)}
	}

	grid[h].Wires = append(grid[h].Wires, w)

	return
}

// using https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function for speed since
// we don't need any actual crypto functionality
func getXYHash(x, y int) uint32 {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%d,%d", x, y)))
	return h.Sum32()
}
