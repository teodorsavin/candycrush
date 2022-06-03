package userInput

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example/teodorsavin/candycrush/table"
)

func WaitForUserInput() (int, int, int, int, error) {
	fmt.Print("Coordinates of the element that you want to move (i.e. C1): ")
	inputCoordinates := bufio.NewScanner(os.Stdin)
	inputCoordinates.Scan()
	oldCoordinates := inputCoordinates.Text()
	if len(oldCoordinates) != 2 {
		return -1, -1, -1, -1, errors.New("invalid input")
	}

	oldCol, oldRow := extractCoordinates(oldCoordinates)
	fmt.Print("Where do you want to move it (i.e. D1): ")

	inputNewCoordinates := bufio.NewScanner(os.Stdin)
	inputNewCoordinates.Scan()
	newCoordinates := inputNewCoordinates.Text()
	if len(newCoordinates) != 2 {
		fmt.Println("I said coordinates, not what you ate last night.")
		return -1, -1, -1, -1, errors.New("invalid input")
	}

	newCol, newRow := extractCoordinates(newCoordinates)
	if newCol != oldCol && newRow != oldRow {
		fmt.Println("invalid move.", "you can move it to an adjacent place vertically or horizontally")
		return -1, -1, -1, -1, errors.New("invalid move")
	}

	return oldCol, oldRow, newCol, newRow, nil
}

func extractCoordinates(input string) (int, int) {
	strY := strings.ToUpper(input[0:1])
	strX := input[1:2]

	x, _ := strconv.Atoi(strX)
	for y, v := range table.Coordinates {
		if v == strY {
			return x, y
		}
	}
	return -1, -1
}