package main

import (
	"cli-text-compare/pkg"
	"flag"
	"fmt"
	"os"
)

func main() {
	fileModUsage := "usage: clidiff --file ./relative/path.json /absolute/path.json"
	stdinUsage := "usage: clidiff --stdin"
	usage := "Run the tool either with --stdin or with --file arguments to get the diff of two things.\nIf you pass no flags, blank comparison mode (--stdin) will run\n"

	stdinMode := flag.Bool("stdin", false, stdinUsage)
	fileMode := flag.Bool("file", false, fileModUsage)
	helpMode := flag.Bool("help", false, usage)

	flag.Parse()

	if *helpMode {
		if *stdinMode {
			fmt.Println(stdinUsage)
			os.Exit(0)
		}

		if *fileMode {
			fmt.Println(fileModUsage)
			os.Exit(0)
		}

		fmt.Printf(usage)
		os.Exit(0)
	}

	if *fileMode {
		if len(flag.Args()) != 2 {
			fmt.Println(fileModUsage)
			fmt.Printf(string(pkg.Red))
			os.Exit(1)
		}

		pkg.TextCompare(*fileMode, &flag.Args()[0], &flag.Args()[1])
		os.Exit(0)
	}

	pkg.TextCompare(*fileMode, nil, nil)
}
