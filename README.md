# CLI Compare Tool

A diffing tool that you can be %100 sure that it doesn't save anything to anywhere.  
You could technically diff even binaries with it.  
Though I wouldn't recommend that as this tool does not employ the infamous binary diffing algorithms such as bsdiff or xdelta.  

<img width="2016" height="749" alt="image" src="https://github.com/user-attachments/assets/ebaf8154-4ea7-4ebc-862d-58f3dd2120e4" />

Run clidiff either with --stdin or with --file arguments to get the diff of two things.
Default mode is blank comparison mode (--stdin).
Available flags:
 * --file to be able to diff two files on your system. (e.g. clidiff --file ./some/filepath.go /home/user/projects/yourproject/file.go)
 * --diff-width 75 to be able to change the width of the output table, default is 50.
 * --stdin opens a standard input reader with instructions on what to do, default mode of clidiff.
 * --diff-only display the output with changed lines only instead of whole file/string outputs.
 * --omit-newline-delimiter do not display the -newline- delimiter when new line characters a part of the diff.
 * --help print either this help message or specific help messages regarding each flag.


Grab the binaries from releases tab or build it yourself using `go build -o clidiff ./cmd/main.go` or `make build`
