package pkg

import (
	"math"
	"slices"
	"strings"
)

type Matrix [][]int

type Action string

const (
	ActionDelete     Action = "delete"
	ActionInsert     Action = "insert"
	ActionSubstitute Action = "substitute"
	ActionMatch      Action = "match"
	ActionAbove             = ActionDelete
	ActionLeft              = ActionInsert
	ActionTopLeft           = ActionSubstitute
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

func calculateCost(i, j int, a, b []string, matrix Matrix) int {
	if a[i-1] == b[j-1] {
		return matrix[i-1][j-1]
	}

	above := matrix[i-1][j] + 1
	left := matrix[i][j-1] + 1
	topLeft := matrix[i-1][j-1] + 1
	return min(above, left, topLeft)
}

// func handleDelimiterReconciliation(i int, lineCount float64, actions *[]Action, actionToAdd Action) {
// 	if i+1 != int(lineCount) {
// 		*actions = append(*actions, actionToAdd)
// 	}
// }

// func Diff(strOne string, strTwo string) []Action {
// 	linesOne, linesTwo := strings.Split(strOne, "\n"), strings.Split(strTwo, "\n")
// 	arrCtr := math.Max(float64(len(linesOne)), float64(len(linesTwo)))
//
// 	var actions []Action
//
// 	for i := range int(arrCtr) {
// 		if len(linesOne)-1 >= i && len(linesTwo)-1 >= i {
// 			if linesOne[i] == linesTwo[i] {
// 				// here, lines are the same, so we should just add the
// 				// character count amount of "ActionMatch"s into the actions array
// 				actions = append(actions, slices.Repeat([]Action{ActionMatch}, len(strings.Split(linesOne[i], "")))...)
// 			} else {
// 				actions = append(actions, levenshtein(linesOne[i], linesTwo[i])...)
// 			}
//
// 			// we should add a ActionMatch to the lenght of the string because
// 			// in go (and every programming language I know for that matter) when we split strings from a character or a set of characters
// 			// the delimiter is excluded from the substrings (logically so), so we have effectively lost the new line from the end of every line
// 			// and I would prefer to keep the texts 100% original solely due to a fear of side-effects later on
// 			// this is true so long as we're not in the last line, last line just ends, does not require an addition
// 			if linesOne[i] != "" {
// 				handleDelimiterReconciliation(i, arrCtr, &actions, ActionMatch)
// 			}
// 		}
//
// 		if len(linesOne)-1 < i {
// 			deletions := slices.Repeat([]Action{ActionDelete}, len(strings.Split(linesTwo[i], "")))
// 			actions = append(actions, deletions...)
//
// 			if linesTwo[i] != "" {
// 				// same delimiter reason, but here it's a deletion as the first line is empty
// 				handleDelimiterReconciliation(i, arrCtr, &actions, ActionDelete)
// 			}
// 		}
//
// 		if len(linesTwo)-1 < i {
// 			additions := slices.Repeat([]Action{ActionDelete}, len(strings.Split(linesOne[i], "")))
// 			actions = append(actions, additions...)
//
// 			if linesOne[i] != "" {
// 				// same delimiter reason, but here it's an addition as the second line is empty
// 				handleDelimiterReconciliation(i, arrCtr, &actions, ActionInsert)
// 			}
// 		}
//
// 	}
//
// 	return actions
// }

// implemented by looking at: https://gist.github.com/jasonm23/449e7c572b46942361bc808357019dda
// though jasonm23 refers to this as the "myers DIFF Algorithm" it seemed more like
// The levenshtein Algorithm to me
func Diff(strOne string, strTwo string) []Action {
	a, b := strings.Split(strOne, ""), strings.Split(strTwo, "")
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
