package main

import (
	"fmt"
	"math/rand"
	"strings"
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
				row = append(row, "")
			} else if elem[i] == 1 {
				/*
					eventually want to print out hashtags between the range of 1 and 3
					but for now just print 2 to make sure it works for debugging case
				*/
				randomNumberOfTimes := rand.Intn(4)
				toPrint := strings.Repeat("#", randomNumberOfTimes)
				row = append(row, toPrint)
			}
		}
		result = append(result, row)
	}
	for _, elem := range result {
		fmt.Println(elem)
	}
}

/*
assumptions:
 1. All matrices will be square matrices
*/
func Next_Board_State(initial_board [][]int) {
	rows := len(initial_board)
	cols := len(initial_board[0])
	result := [][]int{}
	copy(result, initial_board)

	/*
		TODO:
			2. Make sure not to confuse rows with columns. Look closely when one increases and not
			3. Take my time to write the code
			4. Implement tests!
	*/

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// create temporary array to store values for cell state computation
			temp_arr := []int{}
			var a int
			var b int
			var c int
			var d int
			var e int
			var f int
			var g int
			var h int

			// Regular top edge with 5 neighbors
			if i == 0 && (j != 0 && j != rows-1) {
				// top edge with 5 neighbors
				a, b, c, d, e = initial_board[i][j-1], initial_board[i+1][j-1], initial_board[i+1][j], initial_board[i+1][j+1], initial_board[i][j+1]
				// Regular bottom edge with 5 neighbors
			} else if i == rows-1 && (j != 0 && j != rows-1) {
				a, b, c, d, e = initial_board[i][j-1], initial_board[i-1][j-1], initial_board[i-1][j], initial_board[i-1][j+1], initial_board[i][j+1]
				// Regular right edge with 5 neighbors
			} else if j == rows-1 && (i != 0 && i != rows-1) {
				a, b, c, d, e = initial_board[i-1][j], initial_board[i-1][j-1], initial_board[i][j-1], initial_board[i+1][j-1], initial_board[i+1][j]
				// Regular left edge with 5 neighbors
			} else if j == 0 && (i != 0 && i != rows-1) {
				a, b, c, d, e = initial_board[i-1][j], initial_board[i-1][j+1], initial_board[i][j+1], initial_board[i+1][j+1], initial_board[i+1][j]
				// Top Left Corner
			} else if i == 0 && j == 0 {
				a, b, c = initial_board[i][j+1], initial_board[i+1][j+1], initial_board[i+1][j]
				// Top Right Corner
			} else if i == 0 && j == rows-1 {
				a, b, c = initial_board[i][j-1], initial_board[i+1][j-1], initial_board[i+1][j]
				// Bottom Left Corner
			} else if i == rows-1 && j == 0 {
				a, b, c = initial_board[i-1][j], initial_board[i-1][j+1], initial_board[i][j+1]
				// Bottom Right Corner
			} else if i == rows-1 && j == rows-1 {
				a, b, c = initial_board[i-1][j], initial_board[i-1][j-1], initial_board[i][j-1]
				// Trivial case where the cell lies in the middle with 8 neighbors
			} else {
				a, b, c, d, e, f, g, h = initial_board[i-1][j-1], initial_board[i-1][j], initial_board[i-1][j+1], initial_board[i][j+1], initial_board[i+1][j+1], initial_board[i+1][j], initial_board[i+1][j-1], initial_board[i][j-1]
			}
			temp_arr = append(temp_arr, a, b, c, d, e, f, g, h)

			// calculate the rules
			count := 0
			for _, elem := range temp_arr {
				count += elem
			}
			fmt.Printf("The number of live cells is %d\n", count)
		}
	}
}

func main() {
	tt := Random_State(3, 3)
	// Render(tt)
	for _, elem := range tt {
		fmt.Println(elem)
	}
	Next_Board_State(tt)
}
