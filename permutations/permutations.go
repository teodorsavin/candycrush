package permutations

import (
	"errors"
	"example/teodorsavin/candycrush/userInput"
	"example/teodorsavin/candycrush/utils"
	"math/rand"
)
import "example/teodorsavin/candycrush/table"

func RemoveThreeOrMoreHorizontal(ccTable *table.Table) bool {
	var i int
	maxI := len(ccTable.Candies[0]) - 1
	var duplications = false

	for i = maxI; i >= 0; i-- {
		candyDuplicationOnLine := remove3ConsecutiveDuplicatesHorizontal(&ccTable.Candies, i)
		if candyDuplicationOnLine {
			duplications = true
			utils.PrintTable(*ccTable)
		}
	}

	return duplications
}

func remove3ConsecutiveDuplicatesHorizontal(ccTable *[][]table.Candy, row int) bool {
	candyTable := *ccTable
	candyLine := candyTable[row]
	var i, startDuplication, endDuplication int
	var threeDuplicates = false
	startDuplication = -1
	endDuplication = -1
	maxI := len(candyLine) - 1

	for i = 0; i < maxI-1; i++ {
		if candyLine[i].Color == candyLine[i+1].Color && candyLine[i].Color == candyLine[i+2].Color {
			if threeDuplicates == false && startDuplication == -1 {
				startDuplication = i
				endDuplication = i+2
			} else if threeDuplicates == true {
				endDuplication = i+2
			}
			threeDuplicates = true
		} else {
			threeDuplicates = false
		}
	}

	if startDuplication != -1 && endDuplication != -1 && endDuplication - startDuplication > 1 {
		moveCandiesDownHorizontal(startDuplication, endDuplication, ccTable, row)
		return true
	}

	return false
}

func moveCandiesDownHorizontal(startDuplication int, endDuplication int, ccTable *[][]table.Candy, row int) {
	var j int
	candyTable := *ccTable

	for j = startDuplication; j <= endDuplication; j++ {
		if row > 0 {
			candyTable[row][j] = candyTable[row-1][j]
		} else {
			candyTable[row][j] = table.CandiesConst[rand.Intn(7)]
		}
	}

	if row > 0 {
		moveCandiesDownHorizontal(startDuplication, endDuplication, ccTable, row-1)
	}
}

func RemoveThreeOrMoreVertical(cctable *table.Table) bool {
	var i int
	maxI := len(cctable.Candies) - 1
	var duplications = false

	for i = maxI; i >= 0; i-- {
		candyDuplicationOnLine := remove3ConsecutiveDuplicatesVertical(&cctable.Candies, i)
		if candyDuplicationOnLine {
			duplications = true
			utils.PrintTable(*cctable)
		}
	}

	return duplications
}

func remove3ConsecutiveDuplicatesVertical(ccTable *[][]table.Candy, col int) bool {
	candyTable := *ccTable
	var i, startDuplication, endDuplication int
	var threeDuplicates = false
	startDuplication = -1
	endDuplication = -1
	maxI := len(candyTable) - 1

	for i = 0; i < maxI-1; i++ {
		if candyTable[i][col].Color == candyTable[i+1][col].Color && candyTable[i][col].Color == candyTable[i+2][col].Color {
			if threeDuplicates == false && startDuplication == -1 {
				startDuplication = i
				endDuplication = i+2
			} else if threeDuplicates == true {
				endDuplication = i+2
			}
			threeDuplicates = true
		} else {
			threeDuplicates = false
		}
	}

	if startDuplication != -1 && endDuplication != -1 && endDuplication - startDuplication > 1 {
		moveCandiesDownVertical(startDuplication, endDuplication, ccTable, col)
		return true
	}

	return false
}

func moveCandiesDownVertical(startDuplication int, endDuplication int, ccTable *[][]table.Candy, col int) {
	var i,j int
	candyTable := *ccTable

	i = 0
	for j = endDuplication; j >= 0; j-- {
		i++
		if startDuplication-i >= 0 {
			candyTable[j][col] = candyTable[startDuplication-i][col]
		} else {
			candyTable[j][col] = table.CandiesConst[rand.Intn(7)]
		}
	}
}

func makeMove(ccTable *table.Table, oldX int, oldY int, newX int, newY int) {
	candyTable := *ccTable
	tempCandy := candyTable.Candies[oldX][oldY]
	candyTable.Candies[oldX][oldY] = candyTable.Candies[newX][newY]
	candyTable.Candies[newX][newY] = tempCandy
}

func Play(ccTable *table.Table) error {
	for {
		for {
			duplication := RemoveThreeOrMoreHorizontal(ccTable)

			if duplication == false {
				duplication = RemoveThreeOrMoreVertical(ccTable)
			}

			if duplication == false {
				break
			}
		}

		// wait for user permutation
		oldX, oldY, newX, newY, err := userInput.WaitForUserInput()
		if err != nil {
			return errors.New("something went wrong")
		}
		makeMove(ccTable, oldX, oldY, newX, newY)
		utils.PrintTable(*ccTable)

	}
	return nil
}