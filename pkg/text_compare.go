package pkg

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"log"
	"os"
)

func TextCompare() {
	var FatalError = "Something went wrong, please send Omer your input to debug and use a website for comparison for now."

	fmt.Println("Enter the old string, Press Ctrl+D (or Ctrl+Z then Enter on Windows) to move on to the new string:")

	inputOne, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(FatalError)
	}
	oldStr := string(inputOne)

	inputTwo, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(FatalError)
	}
	newStr := string(inputTwo)

	additions, deletions := GreedyDiffAlgo(oldStr, newStr)
	oldColoured, newColoured := ColourTheDiffs(additions, deletions, oldStr, newStr)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Old text:", "New text:"})
	table.SetBorder(true)
	table.SetColumnSeparator(Reset + "|" + Reset)
	table.SetRowSeparator(Reset + "~" + Reset)
	table.Append([]string{oldColoured, newColoured})
	table.Render()
}
