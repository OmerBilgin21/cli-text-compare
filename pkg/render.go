package pkg

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func returnByteSum(args ...[]byte) []byte {
	var sum []byte
	for _, arg := range args {
		sum = append(sum, arg...)
	}
	return sum
}

func colourTheDiffs(alt, neu []byte, actions []Action) ([]byte, []byte) {
	i, j := 0, 0
	var oldOut []byte
	var newOut []byte
	newLine := []byte("-newline-")

	for _, action := range actions {
		switch action {
		case ActionMatch:

			oldOut = append(oldOut, alt[i])
			newOut = append(newOut, neu[j])
			i++
			j++
		case ActionDelete:
			if alt[i] == '\n' {
				oldOut = append(oldOut, returnByteSum(Red, newLine, []byte{alt[i]}, Reset)...)
			} else {
				oldOut = append(oldOut, returnByteSum(Red, []byte{alt[i]}, Reset)...)
			}
			i++
		case ActionInsert:
			if neu[j] == '\n' {
				newOut = append(newOut, returnByteSum(Green, newLine, []byte{neu[j]}, Reset)...)
			} else {
				newOut = append(newOut, returnByteSum(Green, []byte{neu[j]}, Reset)...)
			}
			j++
		case ActionSubstitute:
			// well I used BoldRed and BoldGreen for showing
			// substitute old and substitute new changes
			// but it's not the best to say the least, open for ideas

			if alt[i] == '\n' || neu[j] == '\n' {
				if alt[i] == '\n' {
					oldOut = append(oldOut, returnByteSum(BoldRed, newLine, []byte{alt[i]}, Reset)...)
					newOut = append(newOut, returnByteSum(BoldGreen, []byte{neu[j]}, Reset)...)
				}
				if neu[j] == '\n' {
					oldOut = append(oldOut, returnByteSum(BoldRed, []byte{alt[i]}, Reset)...)
					newOut = append(newOut, returnByteSum(BoldGreen, newLine, []byte{neu[j]}, Reset)...)
				}
			} else {
				oldOut = append(oldOut, returnByteSum(BoldRed, []byte{alt[i]}, Reset)...)
				newOut = append(newOut, returnByteSum(BoldGreen, []byte{neu[j]}, Reset)...)
			}

			i++
			j++
		}
	}

	return oldOut, newOut
}

func RenderDiff(oldThingy, newThingy []byte, actions []Action, width int) {
	oldOut, newOut := colourTheDiffs(oldThingy, newThingy, actions)

	top := lipgloss.NewStyle().
		Width((width*2)+1).
		Border(lipgloss.NormalBorder(), true, true, true, true).
		Render("Result: ")

	left := lipgloss.NewStyle().
		Width(width).
		BorderRight(true).
		BorderStyle(lipgloss.NormalBorder()).
		Render(string(oldOut))

	right := lipgloss.NewStyle().
		Width(width).Render(string(newOut))

	joined := lipgloss.JoinHorizontal(lipgloss.Center, left, right)

	outer := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Render(joined)

	final := lipgloss.JoinVertical(lipgloss.Center, top, outer)
	fmt.Println(final)
}
