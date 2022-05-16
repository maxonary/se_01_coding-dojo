package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var horizontalCheck = true
	var verticalCheck = true

	// gem colors reach from 1 to 5
	var game_field = [16]int{
		1, 2, 3, 4,
		4, 3, 5, 3,
		5, 3, 4, 2,
		3, 1, 2, 5}

	for horizontalCheck == true && verticalCheck == true {

		fmt.Println("This is the current game-board:")
		var i = 0

		for i < 4 {
			fmt.Println(game_field[i], game_field[i+1], game_field[i+2], game_field[i+3])
			horizontalCheck = checkThreeInHorizontal(game_field[i], game_field[i+1], game_field[i+2], game_field[i+3])
			verticalCheck = checkThreeInVertical(game_field[i], game_field[i+4], game_field[i+8], game_field[i+12])
			i++
		}

		fmt.Println("For selecting the first item enter a number from 0 to 15")

		var first_selection = userInputReader()

		var first_item = game_field[first_selection]

		fmt.Println("For selecting the second item enter a number from 0 to 15")

		var second_selection = userInputReader()

		var second_item = game_field[second_selection]

		// switching values
		game_field[first_selection] = second_item
		game_field[second_selection] = first_item
	}
}

func userInputReader() int {
	reader := bufio.NewReader(os.Stdin)
	item_string, err := reader.ReadString('\n')

	// remove the \n
	item_string = strings.TrimSuffix(item_string, "\n")

	// convert string to int
	item_int, err := strconv.Atoi(item_string)
	_ = err
	return item_int
}

func checkThreeInHorizontal(left_gem int, middle_left_gem int, middle_right_gem int, right_gem int) bool {
	// check for same row
	if left_gem/middle_left_gem == 1 && left_gem/middle_right_gem == 1 && middle_left_gem/right_gem == 1 ||
		right_gem/middle_left_gem == 1 && right_gem/middle_right_gem == 1 && middle_left_gem/right_gem == 1 {
		fmt.Println("Horizontal Row of Diamonds")
		return false

	} else {
		return true
	}
}

func checkThreeInVertical(upper_gem int, middle_upper_gem int, middle_lower_gem int, lower_gem int) bool {
	if upper_gem/middle_upper_gem == 1 && upper_gem/middle_lower_gem == 1 && middle_upper_gem/middle_lower_gem == 1 ||
		lower_gem/middle_upper_gem == 1 && lower_gem/middle_lower_gem == 1 && middle_upper_gem/middle_lower_gem == 1 {
		fmt.Println("Vertical Row of Diamonds")
		return false

	} else {
		return true
	}
}
