package main

import (
	"example/teodorsavin/candycrush/permutations"
	"example/teodorsavin/candycrush/utils"
)
import "example/teodorsavin/candycrush/table"

const TableLength = 10
const TableWidth = 10

func main() {
	var initialTable table.Table

	initialTable = table.GenerateInitialTable(TableLength, TableWidth)
	utils.PrintTable(initialTable)

	permutations.Play(&initialTable)
}