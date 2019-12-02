package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			c := NewIntcomputer(noun, verb)
			res := c.Execute()

			if res == 19690720 {
				fmt.Printf("%d: %d\n", 100*noun+verb, res)
				return
			}
		}
	}
}

// Intcomputer represents an Intcomputer machine.
type Intcomputer struct {
	Nums []int
}

// NewIntcomputer initializes the sacred machine.
func NewIntcomputer(noun, verb int) *Intcomputer {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	strNums := strings.Split(string(b), ",")
	var nums []int

	for _, n := range strNums {
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		nums = append(nums, i)
	}

	// Set state
	nums[1] = noun
	nums[2] = verb

	return &Intcomputer{
		Nums: nums,
	}
}

// Execute the machine against the current numbers set.
func (i *Intcomputer) Execute() int {
	for pos := 0; pos+4 <= len(i.Nums); pos += 4 {
		if !i.execCmd(i.Nums[pos], i.Nums[pos+1], i.Nums[pos+2], i.Nums[pos+3]) {
			break
		}
	}
	return i.Nums[0]
}

func (i *Intcomputer) execCmd(opcode, in1, in2, resPos int) bool {
	switch opcode {
	case 1:
		i.Nums[resPos] = i.Nums[in1] + i.Nums[in2]
		return true
	case 2:
		i.Nums[resPos] = i.Nums[in1] * i.Nums[in2]
		return true
	default:
		return false
	}
}
