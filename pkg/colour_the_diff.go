package pkg

var (
	Red       = []byte("\033[31m")
	Green     = []byte("\033[32m")
	BoldRed   = []byte("\033[1;31m")
	BoldGreen = []byte("\033[1;32m")
	Reset     = []byte("\033[0m")
)

func returnByteSum(args ...[]byte) []byte {
	var sum []byte
	for _, arg := range args {
		sum = append(sum, arg...)
	}
	return sum
}

func ColourTheDiffs(alt, neu []byte, actions []Action) []byte {
	i, j := 0, 0
	var out []byte
	newLine := []byte("-newline-")

	for _, action := range actions {
		switch action {
		case ActionMatch:
			out = append(out, alt[i])
			i++
			j++
		case ActionDelete:
			if alt[i] == '\n' {
				out = append(out, returnByteSum(Red, newLine, []byte{alt[i]}, Reset)...)
			} else {
				out = append(out, returnByteSum(Red, []byte{alt[i]}, Reset)...)
			}
			i++
		case ActionInsert:
			if neu[j] == '\n' {
				out = append(out, returnByteSum(Green, newLine, []byte{neu[j]}, Reset)...)
			} else {
				out = append(out, returnByteSum(Green, []byte{neu[j]}, Reset)...)
			}
			j++
		case ActionSubstitute:
			// well I used BoldRed and BoldGreen for showing
			// substitute old and substitute new changes
			// but it's not the best to say the least, open for ideas

			if alt[i] == '\n' || neu[j] == '\n' {
				if alt[i] == '\n' {
					out = append(out, returnByteSum(BoldRed, newLine, []byte{alt[i]}, Reset, BoldGreen, []byte{neu[j]}, Reset)...)
				}
				if neu[j] == '\n' {
					out = append(out, returnByteSum(BoldRed, []byte{alt[i]}, Reset, BoldGreen, newLine, []byte{neu[j]}, Reset)...)
				}
			} else {
				out = append(out, returnByteSum(BoldRed, []byte{alt[i]}, Reset, BoldGreen, []byte{neu[j]}, Reset)...)
			}

			i++
			j++
		}
	}

	return out
}
