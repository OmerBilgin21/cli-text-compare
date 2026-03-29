package pkg

import (
	"fmt"
	"io"
	"os"
)

func TextCompare() {
	fmt.Println("Enter the old string, Press Ctrl+D (or Ctrl+Z then Enter on Windows) to move on to the new string:")

	inputOne, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	oldStr := string(inputOne)

	inputTwo, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	newStr := string(inputTwo)

	actions := Diff(oldStr, newStr)
	colored := ColourTheDiffs(oldStr, newStr, actions)
	fmt.Println("\nresult:")
	fmt.Printf(colored)
}
