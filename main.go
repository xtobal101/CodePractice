package main

import (
	"fmt"
)

/* Remove the third repeated character from a given string
   aaabbbccc -> aabbcc
*/

func removeThird(S string) string {
	result := ""

	rep := 0
	for i, c := range S {

		//fmt.Println(i, c)
		if i == 0 {
			result = result + string(c)
		} else if i > 0 {
			if S[i-1] == S[i] {
				rep++
			} else {
				rep = 0
			}
			if rep < 2 {
				result = result + string(c)
			}
		}

	}

	return result
}

func main() {

	palabra := "aaabbbccc"
	fmt.Println(removeThird(palabra))

	palabra2 := "aaabbbcccdddd111111111222233334444///////3#####"
	fmt.Println(removeThird(palabra2))

	palabra3 := ""
	fmt.Println(removeThird(palabra3))

	palabra4 := "aaaaaaspdghñair´85948325p98yegohdfñgh34uu45u934u598488484848484888888888afasdfas"
	fmt.Println(removeThird(palabra4))

}
