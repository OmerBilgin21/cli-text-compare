package pkg

import (
	"cli-text-compare/internal"
	"fmt"
	"strings"
)

func markRow(longer []string, shorter []string) (string, string) {

	nonEqualWordIndex := -1
	for i := 0; i < len(longer); i++ {
		lWord := longer[i]

		if len(shorter) < i {
			longer = internal.StrInsert(longer, i, Green)
			continue
		}

		sWord := shorter[i]

		if lWord != sWord {
			if nonEqualWordIndex == -1 {
				nonEqualWordIndex = i
				fmt.Println("i: ", i)
				longer = internal.StrInsert(longer, i, Green)
				if len(shorter) >= i {
					shorter = internal.StrInsert(shorter, i, Red)
				}
			}
		} else {
			if nonEqualWordIndex != -1 {
				longer = internal.StrInsert(longer, nonEqualWordIndex, Reset)
				shorter = internal.StrInsert(shorter, nonEqualWordIndex, Reset)
			}
		}
	}

	fmt.Println("nonEqualWordIndex: ", nonEqualWordIndex)
	if nonEqualWordIndex != -1 {
		longer = internal.StrInsert(longer, nonEqualWordIndex, Reset)
		shorter = internal.StrInsert(shorter, nonEqualWordIndex, Reset)
	}

	return strings.Join(longer, " "), strings.Join(shorter, " ")

}

func DiffStr(longerPtr *[]string, shorterPtr *[]string) {
	longer := *longerPtr
	shorter := *shorterPtr

	for i := 0; i < len(longer); i++ {
		longerRowWords := strings.Split(longer[i], " ")
		shorterRowWords := strings.Split(shorter[i], " ")
		nLonger, nShorter := markRow(longerRowWords, shorterRowWords)
		longer[i] = nLonger
		shorter[i] = nShorter
	}

}
