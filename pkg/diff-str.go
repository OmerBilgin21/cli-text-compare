package pkg

import (
	"cli-text-compare/internal"
	"fmt"
	"strings"
)

func markRow(longer []string, shorter []string) (string, string) {
	fmt.Printf("we're at markrow\n")
	var i = 0
	var shorterLen = len(shorter)
	var longerLen = len(longer)
	fmt.Printf("len(shorter): %v\n", len(shorter))
	// var notEqualWord = false
	for ; i < shorterLen; i++ {
		lWord := longer[i]
		sWord := shorter[i]

		if lWord != sWord {
			longer = internal.StrInsert(longer, i, Green)
			shorter = internal.StrInsert(shorter, i, Red)
			fmt.Printf("length after insert: %v\n", len(shorter))
			// notEqualWord = true
		} else {
			// if notEqualWord == true {
			// 	longer = internal.StrInsert(longer, i, Reset)
			// 	shorter = internal.StrInsert(shorter, i, Reset)
			// }
			//
			// notEqualWord = false
		}
	}

	fmt.Printf("i after word arrangement %d\n", i)

	for ; i < longerLen; i++ {
		longer = internal.StrInsert(longer, i, Green)
	}

	fmt.Printf("i after marking rest of the longer sentence %d\n", i)

	return strings.Join(longer, " "), strings.Join(shorter, " ")
}

func DiffStr(longerPtr *[]string, shorterPtr *[]string) {
	longer := *longerPtr
	shorter := *shorterPtr

	i := 0

	for i < len(shorter) {
		fmt.Println("yo")
		longerRowWords := strings.Split(longer[i], " ")
		shorterRowWords := strings.Split(shorter[i], " ")

		longer[i], shorter[i] = markRow(longerRowWords, shorterRowWords)
		i++
	}

	fmt.Printf("i after word/sentence arrangement %d\n", i)

	for i < len(longer) {
		fmt.Println(longer[i])
		i++
	}

	fmt.Printf("i after marking rest of the longer parag. lines: %d\n", i)
}
