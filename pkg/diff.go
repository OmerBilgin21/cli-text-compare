package pkg

import (
	"strings"
)

// the startX and startY are kind of useless, they're just for my convenience
type Coordinate struct {
	startX int
	startY int
	endX   int
	endY   int
}

func GreedyDiffAlgo(oldStr string, newStr string) ([]Coordinate, []Coordinate) {
	oldArr, newArr := strings.Split(oldStr, ""), strings.Split(newStr, "")
	var additions []Coordinate
	var deletions []Coordinate

	x, y := 0, 0
	for x < len(newArr) && y < len(oldArr) {

		// just slide
		if newArr[x] == oldArr[y] {
			for x < len(newArr) && y < len(oldArr) && newArr[x] == oldArr[y] {
				x++
				y++
			}
		} else if x+1 < len(newArr) && newArr[x+1] == oldArr[y] {
			additions = append(additions, Coordinate{
				startX: x,
				startY: y,
				endX:   x + 1,
				endY:   y,
			})
			x++
		} else if y+1 < len(oldArr) && oldArr[y+1] == newArr[x] {
			deletions = append(deletions, Coordinate{
				startX: x,
				startY: y,
				endX:   x,
				endY:   y + 1,
			})
			y++
		} else {
			additions = append(additions, Coordinate{
				startX: x,
				startY: y,
				endX:   x + 1,
				endY:   y,
			})

			deletions = append(deletions, Coordinate{
				startX: x,
				startY: y,
				endX:   x,
				endY:   y + 1,
			})
			x++
			y++
		}
	}

	return additions, deletions
}
