package pkg

import (
	"cli-text-compare/internal"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
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
	longer, shorter := internal.GetBiggerArray(twoArr, oneArr)
	longerStr, shorterStr := MyersTrial(&longer, &shorter)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Text One", "Text Two"})
	table.Append([]string{longerStr, shorterStr})
	table.Render()
}
