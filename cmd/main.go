package main

import (
	"cli-text-compare/pkg"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	args := os.Args[1:]
	fileModUsage := "usage: clidiff --file ./relative/path.json /absolute/path.json, to have clidiff read your inputs from existing files on your system."
	stdinUsage := "usage: clidiff or clidiff --stdin, to have clidiff read your inputs from standard input."
	diffWidthUsage := "usage: clidiff --diff-width 75, to change the output diff view's width."
	diffOnlyUsage := "usage: clidiff --diff-only, to have clidiff only display the differing lines on the output."
	omitNewlineDelimiter := "usage: clidiff --omit-newline-delimiter, to have clidiff not produce -newline- delimiter when new lines are a part of the diff."
	usage := `Run the tool either with --stdin or with --file arguments to get the diff of two things.
Default mode is blank comparison mode (--stdin).
Available flags:
 * --file to be able to diff two files on your system. (e.g. clidiff --file ./some/filepath.go /home/user/projects/yourproject/file.go)
 * --diff-width 75 to be able to change the width of the output table, default is 50.
 * --stdin opens a standard input reader with instructions on what to do, default mode of clidiff.
 * --diff-only display the output with changed lines only instead of whole file/string outputs.
 * --omit-newline-delimiter do not display the -newline- delimiter when new line characters a part of the diff.
 * --help print either this help message or specific help messages regarding each flag.
 `

	fileMode := flag.Bool("file", false, fileModUsage)
	diffOnly := flag.Bool("diff-only", false, diffOnlyUsage)
	noNewlineDelimiter := flag.Bool("omit-newline-delimiter", false, omitNewlineDelimiter)
	diffWidth := flag.Int("diff-width", 50, diffWidthUsage)
	flag.Bool("help", false, usage)
	flag.Bool("stdin", false, stdinUsage)

	flagHelpWasAsked := false
	if slices.Contains(args, "--help") {
		for _, arg := range args {
			if strings.Contains(arg, "stdin") {
				flagHelpWasAsked = true
				fmt.Println(stdinUsage)
				os.Exit(0)
			}

			if strings.Contains(arg, "file") {
				flagHelpWasAsked = true
				fmt.Println(fileModUsage)
				os.Exit(0)
			}

			if strings.Contains(arg, "diff-width") {
				flagHelpWasAsked = true
				fmt.Println(diffWidthUsage)
				os.Exit(0)
			}

			if strings.Contains(arg, "diff-only") {
				flagHelpWasAsked = true
				fmt.Println(diffOnlyUsage)
				os.Exit(0)
			}
		}

		if !flagHelpWasAsked {
			fmt.Println(usage)
			os.Exit(0)
		}
	}

	flag.Parse()

	if *fileMode {
		if len(flag.Args()) < 2 {
			fmt.Println(fileModUsage)
			fmt.Printf(string(pkg.Red))
			os.Exit(1)
		}

		err := pkg.Compare(*fileMode, &flag.Args()[0], &flag.Args()[1], *diffWidth, *diffOnly, *noNewlineDelimiter)

		if err != nil {
			fmt.Printf("error while diffing: %+v", err)
			os.Exit(1)
		}

		os.Exit(0)
	}

	err := pkg.Compare(*fileMode, nil, nil, *diffWidth, *diffOnly, *noNewlineDelimiter)

	if err != nil {
		fmt.Printf("error while diffing: %+v", err)
		os.Exit(1)
	}
}
