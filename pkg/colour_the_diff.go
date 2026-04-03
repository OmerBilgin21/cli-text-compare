package pkg

import (
	"strings"
)

const (
	Red       = "\033[31m"
	Green     = "\033[32m"
	BoldRed   = "\033[1;31m"
	BoldGreen = "\033[1;32m"
	Reset     = "\033[0m"
)

func ColourTheDiffs(oldStr, newStr string, actions []Action) string {
	alt := strings.Split(oldStr, "")
	neu := strings.Split(newStr, "")
	i, j := 0, 0
	var out []string

	for _, action := range actions {
		switch action {
		case ActionMatch:
			out = append(out, alt[i])
			i++
			j++
		case ActionDelete:
			if alt[i] == "\n" {
				out = append(out, Red+"-newline-"+alt[i]+Reset)
			} else {
				out = append(out, Red+alt[i]+Reset)
			}
			i++
		case ActionInsert:
			if neu[j] == "\n" {
				out = append(out, Green+"-newline-"+neu[j]+Reset)
			} else {
				out = append(out, Green+neu[j]+Reset)
			}
			j++
		case ActionSubstitute:
			// well I used BoldRed and BoldGreen for showing
			// substitute old and substitute new changes
			// but it's not the best to say the least, open for ideas

			if alt[i] == "\n" || neu[j] == "\n" {
				if alt[i] == "\n" {
					out = append(out, BoldRed+"-newline-"+alt[i]+Reset, BoldGreen+neu[j]+Reset)
				}
				if neu[j] == "\n" {
					out = append(out, BoldRed+alt[i]+Reset, BoldGreen+"-newline-"+neu[j]+Reset)
				}
			} else {
				out = append(out, BoldRed+alt[i]+Reset, BoldGreen+neu[j]+Reset)
			}

			i++
			j++
		}
	}

	return strings.Join(out, "")
}
