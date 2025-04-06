package pkg

import (
	"strings"
	_ "strings"
)

var Red = "\033[31m"
var Green = "\033[32m"
var Reset = "\033[0m"

func ColourTheDiffs(
	additions []Coordinate,
	deletions []Coordinate,
	oldStr string,
	newStr string,
) (string, string) {
	// for deletions I have to look at the changing Y
	// for additions I have to look at the changing X
	oldArr, newArr := strings.Split(oldStr, ""), strings.Split(newStr, "")

	for _, addition := range additions {
		xPos := addition.startX
		if xPos < len(newArr) {
			newArr[xPos] = Green + newArr[xPos] + Reset
		}
	}

	for _, deletion := range deletions {
		yPos := deletion.startY
		if yPos < len(oldArr) {
			oldArr[yPos] = Red + oldArr[yPos] + Reset
		}
	}

	return strings.Join(oldArr, ""), strings.Join(newArr, "")
}
