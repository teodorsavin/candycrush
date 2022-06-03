package table

import (
	"math/rand"
	"time"
)

const ColorReset = "\033[0m"
const ColorRed = "\033[31m"
const ColorGreen = "\033[32m"
const ColorYellow = "\033[33m"
const ColorBlue = "\033[34m"
const ColorPurple = "\033[35m"
const ColorCyan = "\033[36m"
const ColorWhite = "\033[37m"

var Coordinates = map[int]string {
	0: "A",
	1: "B",
	2: "C",
	3: "D",
	4: "E",
	5: "F",
	6: "G",
	7: "H",
	8: "I",
	9: "J",
}

type Candy struct {
	Id int
	Color string
	ColorHex string
}

type Table struct {
	Length int
	Width int
	Candies [][]Candy
}

var CandiesConst = map[int]Candy {
	0: {Id: 1, Color: "Red", ColorHex: ColorRed},
	1: {Id: 2, Color: "Green", ColorHex: ColorGreen},
	2: {Id: 3, Color: "Blue", ColorHex: ColorBlue},
	3: {Id: 4, Color: "Yellow", ColorHex: ColorYellow},
	4: {Id: 5, Color: "Purple", ColorHex: ColorPurple},
	5: {Id: 6, Color: "Cyan", ColorHex: ColorCyan},
	6: {Id: 7, Color: "White", ColorHex: ColorWhite},
}

func GenerateInitialTable(length int, width int) Table {
	candies := make([][]Candy, length)
	for k := range candies {
		candies[k] = make([]Candy, width)
	}
	table := Table{Length: length, Width: width, Candies: candies}
	var i, j int

	rand.Seed(time.Now().UnixNano())

	for i = 0; i < length; i++ {
		for j = 0; j < width; j++ {
			table.Candies[i][j] = CandiesConst[rand.Intn(7)]
		}
	}

	return table
}