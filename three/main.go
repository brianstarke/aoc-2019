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
	wiremap  map[wiretype]*wire
	wiretype int
	wire     struct {
		Length int
	}
	point struct {
		X        int
		Y        int
		Wires    wiremap
		Distance int // manhattan distance from origin
	}
)

// constants
const (
	WIRE1 wiretype = 1
	WIRE2 wiretype = 2
)

func main() {
	initGrid()

	// get the intersections
	var intersections []*point

	for _, p := range grid {
		if len(p.Wires) > 1 && hasWiretype(p.Wires, WIRE1) && hasWiretype(p.Wires, WIRE2) {
			intersections = append(intersections, p)
		}
	}

	var closest int
	var shortest int

	for idx, i := range intersections {
		l := sumWirelengths(i.Wires)

		if idx == 0 {
			closest = i.Distance
			shortest = l
		}

		if i.Distance < closest {
			closest = i.Distance
		}

		if l < shortest {
			shortest = l
		}
	}

	log.Printf("Part 1 (closest intersection): %d", closest)
	log.Printf("Part 2 (shortest intersection): %d", shortest)
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

func plotLines(directions []string, w wiretype) {
	// set current X and Y starting points, these will be used as a cursor in the below
	// loop.
	cX, cY, length := 0, 0, 0

	for _, d := range directions {
		direction := d[0:1]
		distance, err := strconv.Atoi(d[1:])
		if err != nil {
			panic(err)
		}
		cX, cY, length = plotLine(cX, cY, length, direction, distance, w)
	}
}

// returns new cursor position
func plotLine(curX, curY, length int, direction string, distance int, w wiretype) (int, int, int) {
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
		length++
		addPoint(curX, curY, length, w)
	}

	return curX, curY, length
}

func addPoint(x, y, length int, wt wiretype) {
	h := getXYHash(x, y)

	// Add this point to the grid if it is not already present.
	if _, ok := grid[h]; !ok {
		grid[h] = &point{X: x, Y: y, Distance: manhattanDistance(x, y), Wires: make(wiremap)}
	}

	grid[h].Wires[wt] = &wire{Length: length}

	return
}

// Helper function to check map for key existence.
func hasWiretype(wires wiremap, wt wiretype) bool {
	if _, ok := wires[wt]; !ok {
		return false
	}
	return true
}

// Helper function to sum the lengths of ze wires.
func sumWirelengths(wires wiremap) (sum int) {
	for _, v := range wires {
		sum += v.Length
	}
	return
}

// using https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function for speed since
// we don't need any actual crypto functionality
func getXYHash(x, y int) uint32 {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%d,%d", x, y)))
	return h.Sum32()
}
