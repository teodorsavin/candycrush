package utils

import (
	"example/teodorsavin/candycrush/table"
	"fmt"
	"sort"
)

func PrintTable(ccTable table.Table) {
	var i, j int
	coordinates := orderCoordinates()

	fmt.Print(table.ColorReset, "    ")
	for _, y := range coordinates {
		fmt.Print(y, " ")
	}
	fmt.Println()
	fmt.Println("-----------------------")
	for i = 0; i < ccTable.Length; i++ {
		fmt.Print(table.ColorReset, i, " | ")
		for j = 0; j < ccTable.Width; j++ {
			fmt.Print(ccTable.Candies[i][j].ColorHex, ccTable.Candies[i][j].Color[0:1] + " ")
		}
		fmt.Println()
	}
	fmt.Println(table.ColorReset)
}

func orderCoordinates() []string {
	keys := make([]string, 0)
	for _, y := range table.Coordinates {
		keys = append(keys, y)
	}
	sort.Strings(keys)
	return keys
}
