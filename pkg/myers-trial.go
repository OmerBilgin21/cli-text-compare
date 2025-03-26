package pkg

import (
	"cli-text-compare/internal"
	"strings"
)

// file is currently as brute force as it gets but I think it works?
// will try to slowly convert it to myers diff alg. after I understand it better
// to also appreciate how fast it will be compared to this bs

func prepareData(longer []string, shorter []string) ([][]string, [][]string) {
	longerRows := len(longer)
	longerGrid := make([][]string, longerRows)
	for i := 0; i < longerRows; i++ {
		longerRowCols := strings.Split(longer[i], "")
		longerGrid[i] = make([]string, len(longerRowCols))
		longerGrid[i] = longerRowCols
	}

	shorterRows := len(shorter)
	shorterGrid := make([][]string, shorterRows)
	for i := 0; i < shorterRows; i++ {
		shorterRowCols := strings.Split(shorter[i], "")
		shorterGrid[i] = make([]string, len(shorterRowCols))
		shorterGrid[i] = shorterRowCols
	}

	return longerGrid, shorterGrid
}

func MyersTrial(longerPtr *[]string, shorterPtr *[]string) (string, string) {
	longer, shorter := *longerPtr, *shorterPtr
	longerGrid, shorterGrid := prepareData(longer, shorter)
	sLenOrig, lLenOrig := len(shorterGrid), len(longerGrid)

	var i = 0

	for ; i < sLenOrig; i++ {
		row1 := shorterGrid[i]
		row2 := longerGrid[i]
		lRow, sRow := internal.GetBiggerArray(row1, row2)
		lRowOrig, sRowOrig := len(lRow), len(sRow)

		var j = 0
		for ; j < sRowOrig; j++ {
			lChar := lRow[j]
			sChar := sRow[j]
			if lChar != sChar {
				lRow[j] = Green + lChar
				sRow[j] = Red + sChar
			} else {
				lRow[j] = Reset + lChar
				sRow[j] = Reset + sChar
			}
		}

		for k := j; k < lRowOrig; k++ {
			lRow = append(lRow, Green)
		}

		lRow = append(lRow, Reset)
		sRow = append(sRow, Reset)
	}

	for l := i; l < lLenOrig; l++ {
		longerGrid[l] = append(longerGrid[l], Green)
	}

	var longerReturnArr = make([]string, len(longerGrid))
	for a := 0; a < len(longerGrid); a++ {
		longerReturnArr[a] = strings.Join(longerGrid[a], "")
	}

	var shorterReturnArr = make([]string, len(shorterGrid))
	for a := 0; a < len(shorterGrid); a++ {
		shorterReturnArr[a] = strings.Join(shorterGrid[a], "")
	}
	longerStr := strings.Join(longerReturnArr, "\n")
	shorterStr := strings.Join(shorterReturnArr, "\n")

	return longerStr, shorterStr
}
