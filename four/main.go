package main

import (
	"log"
	"strconv"
)

func main() {
	num := findValidPasswords(359282, 820401)
	log.Println(num)
}

func findValidPasswords(min, max int) (count int) {
	for i := min; i < max; i++ {
		if obeysRules(i) {
			count++
		}
	}
	return
}

func obeysRules(num int) bool {
	neverDecreases := true

	l := numToList(num)

	groups := 0

	for i := 0; i < len(l); i++ {
		// see how many we have in a row
		cnt := 0
		cur := l[i]

		for j := i + 1; j < len(l); j++ {
			if l[j] != cur {
				break
			}
			cnt++
		}

		if cnt == 1 {
			groups++
		}

		if cnt > 1 {
			i += cnt
		}
	}

	for i := 1; i < len(l); i++ {
		if l[i] < l[i-1] {
			neverDecreases = false
		}
	}

	return groups > 0 && neverDecreases
}

func numToList(num int) (listNum []int) {
	strNum := strconv.Itoa(num)

	for i := 0; i < len(strNum); i++ {
		n, _ := strconv.Atoi(string(strNum[i]))
		listNum = append(listNum, n)
	}
	return
}
