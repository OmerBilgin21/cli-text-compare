package pkg

import (
	"bytes"
	"math"
	"slices"
)

func min(x, y, z int) int {
	return int(math.Min(math.Min(float64(x), float64(y)), float64(z)))
}

func findSmallestResistance(i, j int, matrix Matrix) Action {
	above := matrix[i-1][j]
	left := matrix[i][j-1]
	topLeft := matrix[i-1][j-1]

	if topLeft <= above && topLeft <= left {
		return ActionTopLeft
	}
	if above <= left {
		return ActionAbove
	}
	return ActionLeft
}

func calculateCost(i, j int, a, b []byte, matrix Matrix) int {
	if a[i-1] == b[j-1] {
		return matrix[i-1][j-1]
	}

	above := matrix[i-1][j] + 1
	left := matrix[i][j-1] + 1
	topLeft := matrix[i-1][j-1] + 1
	return min(above, left, topLeft)
}

func Diff(oldInput []byte, newInput []byte) []Action {
	oldLines, newLines := bytes.SplitAfter(oldInput, Separator), bytes.SplitAfter(newInput, Separator)
	arrCtr := math.Max(float64(len(oldLines)), float64(len(newLines)))

	var actions []Action

	for i := range int(arrCtr) {
		// this loop one and below are basically us running out of one of the lines arrays
		// so depending on which one, it's just a straight up addition/removal until longer ends
		if len(oldLines)-1 < i {
			for j := i; j < len(newLines); j++ {
				for range len(newLines[j]) {
					actions = append(actions, ActionInsert)
				}
			}

			break
		}

		if len(newLines)-1 < i {
			for j := i; j < len(oldLines); j++ {
				for range len(oldLines[j]) {
					actions = append(actions, ActionDelete)
				}
			}
			break
		}

		if slices.Equal(oldLines[i], newLines[i]) {
			// here, lines are the same, so we should just add the
			// character count amount of "ActionMatch"s into the actions array
			actions = append(actions, slices.Repeat([]Action{ActionMatch}, len(oldLines[i]))...)
		} else {
			actions = append(actions, levenshtein(oldLines[i], newLines[i])...)
		}

	}

	return actions
}

// implemented by looking at: https://gist.github.com/jasonm23/449e7c572b46942361bc808357019dda
// though jasonm23 refers to this as the "myers DIFF Algorithm" it seemed more like
// The levenshtein Algorithm to me
func levenshtein(a []byte, b []byte) []Action {
	// a, b := strings.Split(strOne, ""), strings.Split(strTwo, "")
	m, n := len(a), len(b)

	matrix := make(Matrix, m+1)
	for i := range matrix {
		matrix[i] = make([]int, n+1)
	}

	for i := range n + 1 {
		matrix[0][i] = i
	}
	for j := range m + 1 {
		matrix[j][0] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			matrix[i][j] = calculateCost(i, j, a, b, matrix)
		}
	}

	var actions []Action
	i, j := m, n
	for i > 0 || j > 0 {
		if i > 0 && j > 0 && a[i-1] == b[j-1] {
			actions = append(actions, ActionMatch)
			i--
			j--
		} else if i == 0 {
			actions = append(actions, ActionInsert)
			j--
		} else if j == 0 {
			actions = append(actions, ActionDelete)
			i--
		} else {
			action := findSmallestResistance(i, j, matrix)
			actions = append(actions, action)
			switch action {
			case ActionAbove:
				i--
			case ActionLeft:
				j--
			case ActionTopLeft:
				i--
				j--
			}
		}
	}

	slices.Reverse(actions)
	return actions
}
