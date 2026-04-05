package pkg

import (
	"fmt"
	"io"
	"os"
	"slices"
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

var (
	Separator = []byte("\n")
	Red       = []byte("\033[31m")
	Green     = []byte("\033[32m")
	BoldRed   = []byte("\033[1;31m")
	BoldGreen = []byte("\033[1;32m")
	Reset     = []byte("\033[0m")
)

func readOrExit(path string) ([]byte, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf(string(Red)+"error while reading the file: %+v\n", err)
	}

	return data, nil
}

func Compare(fileMode bool, filePathOne, filePathTwo *string, diffWidth int) error {
	if fileMode {
		fileOne, err := readOrExit(*filePathOne)
		if err != nil {
			return err
		}

		fileTwo, err := readOrExit(*filePathTwo)
		if err != nil {
			return err
		}

		actions := Diff(fileOne, fileTwo)

		if slices.Compare(actions, slices.Repeat([]Action{ActionMatch}, len(actions))) == 0 {
			fmt.Println("\nthe texts are the same")
			return nil
		}

		RenderDiff(fileOne, fileTwo, actions, diffWidth)
		return nil
	}

	fmt.Println("Enter the old string, Press Ctrl+D (or Ctrl+Z then Enter on Windows) to move on to the new string:")
	inputOne, err := io.ReadAll(os.Stdin)

	if err != nil {
		return err
	}

	inputTwo, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	actions := Diff(inputOne, inputTwo)

	if slices.Compare(actions, slices.Repeat([]Action{ActionMatch}, len(actions))) == 0 {
		fmt.Println("\nthe texts are the same")
		return nil
	}

	RenderDiff(inputOne, inputTwo, actions, diffWidth)
	return nil
}
