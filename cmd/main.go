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
	fileModUsage := "usage: clidiff --file ./relative/path.json /absolute/path.json"
	stdinUsage := "usage: clidiff or clidiff --stdin"
	diffWidthUsage := "usage: clidiff --diff-width 75"
	usage := "Run the tool either with --stdin or with --file arguments to get the diff of two things.\nIf you pass no flags, blank comparison mode (--stdin) will run\nTo change the width of the rendered diffs, pass --diff-width X\n"

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
