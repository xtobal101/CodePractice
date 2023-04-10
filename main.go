package main

import (
	"fmt"
)

/* Calculate the minimal rotation of arrows to make the string all equal
   >>><<v^ -> 4
*/

func minRotArrow(S string) int {
	result := 0
	max := 0
	totArrows := 0

	uA := 0
	dA := 0
	lA := 0
	rA := 0

	for _, ch := range S {
		if ch == '^' {
			totArrows++
			uA++
		} else if ch == 'v' {
			totArrows++
			dA++
		} else if ch == '<' {
			totArrows++
			lA++
		} else if ch == '>' {
			totArrows++
			rA++
		}
	}

	sums := [4]int{uA, dA, lA, rA}

	for _, v := range sums {
		if v > max {
			max = v
		}
	}

	result = totArrows - max

	return result
}

func main() {

	word0 := ""
	fmt.Println(word0, minRotArrow(word0))

	word1 := ">^<v"
	fmt.Println(word1, minRotArrow(word1))

	word2 := ">>>><<<^^v"
	fmt.Println(word2, minRotArrow(word2))

	word3 := ">^<v>^<v>^<v"
	fmt.Println(word2, minRotArrow(word3))

	word4 := "^"
	fmt.Println(word4, minRotArrow(word4))

	word5 := "1234^QWER"
	fmt.Println(word5, minRotArrow(word5))

	word6 := ">^"
	fmt.Println(word6, minRotArrow(word6))

	word7 := ">^<v>^alskdj34564756/$(&%$&$&%$haksjdfh<v>^<vv"
	fmt.Println(word7, minRotArrow(word7))
}
