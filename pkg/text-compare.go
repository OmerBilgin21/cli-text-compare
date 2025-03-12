package pkg

import (
	// "bufio"
	"cli-text-compare/internal"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	// "github.com/eiannone/keyboard"
)

var Red = "\033[31m"
var Green = "\033[32m"
var Reset = "\033[0m"

func TextCompare() {
	fmt.Printf("%s red part %s regular part %s green part\n", Red, Reset, Green)
	fmt.Printf(Reset)
	fmt.Println("Enter input one, Press Ctrl+D (or Ctrl+Z then Enter on Windows) to move on to input two:")
	inputOne, err := io.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal(internal.FatalError)
	}

	strInOne := string(inputOne)

	inputTwo, err := io.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal(internal.FatalError)
	}

	strInTwo := string(inputTwo)

	oneArr := strings.Split(strInOne, "\n")
	twoArr := strings.Split(strInTwo, "\n")
	if len(twoArr) > len(oneArr) {
		DiffStr(&twoArr, &oneArr)

	} else {
		DiffStr(&oneArr, &twoArr)
	}

	diffedOne := strings.Join(oneArr, "\n")
	diffedTwo := strings.Join(twoArr, "\n")

	fmt.Println(fmt.Sprintf("Diff: \n\n%s\n\n%s", diffedOne, diffedTwo))
}
