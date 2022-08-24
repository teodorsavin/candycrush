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
	oldCol, oldRow, err := readCoordinates(os.Stdin)
	if err != nil {
		return -1, -1, -1, -1, errors.New("invalid input for current position")
	}

	fmt.Print("Where do you want to move it (i.e. D1): ")
	newCol, newRow, err := readCoordinates(os.Stdin)
	if err != nil {
		return -1, -1, -1, -1, errors.New("invalid input for new position")
	}

	if newCol != oldCol && newRow != oldRow {
		fmt.Println("invalid move.", "you can move it to an adjacent place vertically or horizontally")
		return -1, -1, -1, -1, errors.New("invalid move")
	}

	if abs(newCol - oldCol) > 1 || abs(newRow - oldRow) > 1 {
		fmt.Println("invalid move.", "you can move it to an adjacent place vertically or horizontally")
		return -1, -1, -1, -1, errors.New("invalid move")
	}

	return oldCol, oldRow, newCol, newRow, nil
}

func readCoordinates(in *os.File) (int, int, error) {
	inputCoordinates := bufio.NewScanner(in)
	inputCoordinates.Scan()
	coordinates := inputCoordinates.Text()
	if len(coordinates) != 2 {
		return -1, -1, errors.New("invalid input")
	}

	col, row := extractCoordinates(coordinates)

	return col, row, nil
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

// Abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}