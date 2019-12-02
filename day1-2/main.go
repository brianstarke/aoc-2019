package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	nums := getNumNums()
	fuelSum := 0

	for _, mass := range nums {
		fuel := calcFuel(mass)
		fuelSum += fuel
	}

	log.Println(fuelSum)
}

func calcFuel(mass int) int {
	initFuel := int(mass/3) - 2
	_, addlFuel := calcAdditionalFuel(initFuel, 0)
	return initFuel + addlFuel
}

func calcAdditionalFuel(fuel, fuelSum int) (int, int) {
	additionalFuel := int(fuel/3) - 2
	if additionalFuel < 1 {
		return 0, fuelSum
	}
	return calcAdditionalFuel(additionalFuel, fuelSum+additionalFuel)
}

// could be more efficient by reading the integers in to the result array
// while parsing the file line by line
func getNumNums() []int {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	strNums := strings.Split(string(b), "\n")
	var numNums []int

	for _, n := range strNums {
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		numNums = append(numNums, i)
	}
	return numNums
}
