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
	var game_field = [9]int{
		1, 2, 3,
		4, 3, 5,
		5, 3, 4}

	for horizontalCheck == true && verticalCheck == true {

		fmt.Println("This is the current game-board:")
		var i = 0

		for i < 3 {
			fmt.Println(game_field[i], game_field[i+1], game_field[i+2])
			horizontalCheck = checkThreeInHorizontal(game_field[i], game_field[i+1], game_field[i+2])
			verticalCheck = checkThreeInVertical(game_field[i], game_field[i+3], game_field[i+6])
			i++
		}

		fmt.Println("For selecting the first item enter a number from 0 to 8")

		var first_selection = userInputReader()

		var first_item = game_field[first_selection]

		var second_legal = inputIsLegal(first_selection)

		fmt.Println("For selecting the second item enter a number from \n", second_legal)

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

func checkThreeInHorizontal(left_gem int, middle_gem int, right_gem int) bool {
	// check for same row
	if left_gem/middle_gem == 1 && left_gem/right_gem == 1 && middle_gem/right_gem == 1 {
		fmt.Println("Horizontal Row of Diamonds")
		return false

	} else {
		return true
	}
}

func checkThreeInVertical(upper_gem int, middle_gem int, lower_gem int) bool {
	if upper_gem/middle_gem == 1 && upper_gem/lower_gem == 1 && middle_gem/lower_gem == 1 {
		fmt.Println("Vertical Row of Diamonds")
		return false

	} else {
		return true
	}
}

func inputIsLegal(item int) []int {

	if item == 0 {
		allowed := []int{1, 3}
		return allowed
	}
	if item == 1 {
		allowed := []int{0, 2, 4}
		return allowed
	}
	if item == 2 {
		allowed := []int{1, 5}
		return allowed
	}
	if item == 3 {
		allowed := []int{0, 5, 6}
		return allowed
	}
	if item == 4 {
		allowed := []int{1, 3, 5, 7}
		return allowed
	}
	if item == 5 {
		allowed := []int{2, 4, 8}
		return allowed
	}
	if item == 6 {
		allowed := []int{3, 7}
		return allowed
	}
	if item == 7 {
		allowed := []int{4, 6, 8}
		return allowed
	}
	if item == 8 {
		allowed := []int{5, 7}
		return allowed
	} else {
		return []int{0}
	}

}
