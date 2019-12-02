package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var nums []int

func main() {
	for pos := 0; pos+4 <= len(nums); pos += 4 {
		if !execCmd(nums[pos], nums[pos+1], nums[pos+2], nums[pos+3]) {
			break
		}
	}
	log.Println(nums[0])
}

func execCmd(opcode, in1, in2, resPos int) bool {
	switch opcode {
	case 1:
		nums[resPos] = nums[in1] + nums[in2]
		return true
	case 2:
		nums[resPos] = nums[in1] * nums[in2]
		return true
	default:
		return false
	}
}

// could be more efficient by reading the integers in to the result array
// while parsing the file line by line
func init() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	strNums := strings.Split(string(b), ",")

	for _, n := range strNums {
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		nums = append(nums, i)
	}

	// Restore state
	nums[1] = 12
	nums[2] = 2
}
