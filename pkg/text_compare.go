package pkg

import (
	"fmt"
	"io"
	"os"
	"slices"
)

func readOrExit(path string) []byte {
	data, err := os.ReadFile(path)

	if err != nil {
		fmt.Printf(string(Red)+"error while reading the file: %+v\n", err)
		os.Exit(1)
	}

	return data
}

func TextCompare(fileMode bool, filePathOne, filePathTwo *string) {
	if fileMode {
		fileOne := readOrExit(*filePathOne)
		fileTwo := readOrExit(*filePathTwo)

		actions := Diff(fileOne, fileTwo)

		if slices.Compare(actions, slices.Repeat([]Action{ActionMatch}, len(actions))) == 0 {
			fmt.Println("\nthe texts are the same")
			return
		}

		colored := ColourTheDiffs(fileOne, fileTwo, actions)
		fmt.Println("\nresult:")
		fmt.Printf(string(colored))
		return
	}

	fmt.Println("Enter the old string, Press Ctrl+D (or Ctrl+Z then Enter on Windows) to move on to the new string:")
	inputOne, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	inputTwo, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	actions := Diff(inputOne, inputTwo)

	if slices.Compare(actions, slices.Repeat([]Action{ActionMatch}, len(actions))) == 0 {
		fmt.Println("\nthe texts are the same")
		return
	}

	colored := ColourTheDiffs(inputOne, inputTwo, actions)

	fmt.Println("\nresult:")
	fmt.Printf(string(colored))
}
