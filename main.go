package main

import (
	"fmt"
	"math/rand"
)

func Random_State(width, height int) [][]int {
	result := [][]int{}
	for i := 0; i < height; i++ {
		temp := make([]int, width)
		for idx := range temp {
			temp[idx] = rand.Intn(2)
		}
		result = append(result, temp)
	}
	return result
}

func Render(random_state [][]int) {
	result := [][]string{}
	for _, elem := range random_state {
		row := []string{}
		for i := 0; i < len(elem); i++ {
			// if it's dead literally do nothing
			if elem[i] == 0 {
				row = append(row, "💀")
			} else if elem[i] == 1 {
				/*
					eventually want to print out hashtags between the range of 1 and 3
					but for now just print 2 to make sure it works for debugging case
				*/
				row = append(row, "#")
			}
		}
		result = append(result, row)
	}
	for _, elem := range result {
		fmt.Println(elem)
	}
}

func main() {
	tt := Random_State(5, 5)
	Render(tt)
	for _, elem := range tt {
		fmt.Println(elem)
	}
}
