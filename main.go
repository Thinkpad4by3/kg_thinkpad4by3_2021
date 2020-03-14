package main

import (
	"fmt"
	"os"
)

func main() {
	defTable := make(map[rune]byte) //create character map from str1 to str2

	args := os.Args[1:]
	str1 := args[0]
	str2 := args[1]

	if len(str1) != len(str2) { //check if str1 and str2 are the same length
		fmt.Println("false")
		return
	}

	for index, char := range str1 { //loop through all characters in str1
		if defTable[char] != 0 { //if definition exists, check if it is the same
			if str2[index] != defTable[char] { //if not, return false
				fmt.Println("false")
				return
			}
		} else { //if definition doesn't exist, add it to the map
			defTable[char] = str2[index]
		}
	}

	fmt.Println("true")
}
